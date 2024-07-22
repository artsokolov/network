package service

import (
	"context"
	"github.com/artsokolov/network/model"
	"github.com/artsokolov/network/repository"
	"github.com/artsokolov/network/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Posts struct {
	client        *mongo.Client
	posts         *repository.Posts
	profiles      *repository.Profiles
	notifications *repository.Notifications
}

func NewPostService(client *mongo.Client, posts *repository.Posts, profiles *repository.Profiles, notifications *repository.Notifications) *Posts {
	return &Posts{client, posts, profiles, notifications}
}

func (service *Posts) transactionHandler(ctx context.Context, handler func(mongo.SessionContext) error) error {
	session, err := service.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	if err = session.StartTransaction(); err != nil {
		return err
	}

	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		err = handler(sessionContext)
		if err != nil {
			return err
		}

		err = session.CommitTransaction(sessionContext)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (service *Posts) CreatePost(ctx context.Context, request *request.CreatePostRequest, profile *model.Profile) error {
	return service.transactionHandler(ctx, func(sessionContext mongo.SessionContext) error {
		newPost := profile.CreatePost(request.Content)

		err := service.posts.Create(sessionContext, newPost)
		if err != nil {
			return err
		}

		err = service.profiles.Update(sessionContext, profile)
		if err != nil {
			return err
		}

		return nil
	})
}

func (service *Posts) List(profile *model.Profile) ([]model.Post, error) {
	res, err := service.posts.ByUser(profile.ID)
	if err != nil {
		return make([]model.Post, 0), err
	}

	return res, nil
}

func (service *Posts) Like(ctx context.Context, profile *model.Profile, postId string) error {
	objId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}

	return service.transactionHandler(ctx, func(sessionContext mongo.SessionContext) error {
		likeValue, notification := profile.Like(objId)
		if notification != nil {
			err = service.notifications.Create(ctx, notification)
			if err != nil {
				return err
			}

			author, err := service.profiles.ByPostId(postId)
			if err != nil {
				return err
			}

			if author.ID != profile.ID {
				author.AddNotification(notification)
				err = service.profiles.Update(sessionContext, author)
				if err != nil {
					return err
				}
			}
		}

		err = service.profiles.Update(sessionContext, profile)
		if err != nil {
			return err
		}

		err = service.posts.IncreaseLikes(sessionContext, objId, likeValue)
		if err != nil {
			return err
		}

		return nil
	})
}

func (service *Posts) Liked(postIds []primitive.ObjectID) ([]model.Post, error) {
	return service.posts.LikedPosts(postIds)
}
