package repository

import (
	"context"
	"errors"
	"github.com/artsokolov/network/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Posts struct {
	collection *mongo.Collection
}

func NewPosts(collection *mongo.Collection) *Posts {
	return &Posts{collection: collection}
}

func (posts *Posts) Create(ctx context.Context, post *model.Post) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err := posts.collection.InsertOne(ctx, post)
	if err != nil {
		return err
	}

	return nil
}

func (posts *Posts) findManyBy(filter bson.D) ([]model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	postsResult := make([]model.Post, 0)

	cur, err := posts.collection.Find(ctx, filter)
	if err != nil {
		return postsResult, err
	}

	for cur.Next(context.TODO()) {
		var post model.Post
		if err := cur.Decode(&post); err != nil {
			return postsResult, nil
		}
		postsResult = append(postsResult, post)
	}

	if err := cur.Err(); err != nil {
		return postsResult, nil
	}

	return postsResult, nil
}

func (posts *Posts) ByUser(profileId primitive.ObjectID) ([]model.Post, error) {
	return posts.findManyBy(bson.D{{"author", profileId}})
}

func (posts *Posts) IncreaseLikes(ctx context.Context, postId primitive.ObjectID, val int) error {
	res, err := posts.collection.UpdateByID(ctx, postId, bson.D{{"$inc", bson.D{{"likesCount", val}}}})
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 || res.ModifiedCount != 1 {
		return errors.New("more than one post to update")
	}

	return nil
}

func (posts *Posts) LikedPosts(postIds []primitive.ObjectID) ([]model.Post, error) {
	return posts.findManyBy(bson.D{{"_id", bson.D{{"$in", postIds}}}})
}
