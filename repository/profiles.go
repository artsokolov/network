package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/artsokolov/network/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	ErrInvalidProfileId = errors.New("invalid profile id")
	ErrProfileNotFound  = errors.New("profile was not found")
	ErrUnknown          = errors.New("error while decoding profile")
	ErrUpdateProfile    = errors.New("error while updating profile")
)

type Profiles struct {
	collection *mongo.Collection
}

func NewProfiles(collection *mongo.Collection) *Profiles {
	return &Profiles{collection}
}

func (profiles *Profiles) Create(profile *model.Profile) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := profiles.collection.InsertOne(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}

func (profiles *Profiles) findOneBy(filter bson.D) (*model.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res := profiles.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, ErrProfileNotFound
		}
		return nil, res.Err()
	}

	var profile model.Profile
	err := res.Decode(&profile)
	if err != nil {
		return nil, ErrUnknown
	}

	return &profile, nil
}

func (profiles *Profiles) Find(id string) (*model.Profile, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidProfileId
	}

	return profiles.findOneBy(bson.D{{"_id", objId}})
}

func (profiles *Profiles) ByEmail(email string) (*model.Profile, error) {
	return profiles.findOneBy(bson.D{{"email", email}})
}

func (profiles *Profiles) ByPostId(postId string) (*model.Profile, error) {
	key := fmt.Sprintf("posts.%s", postId)
	return profiles.findOneBy(bson.D{{key, bson.D{{"$exists", "true"}}}})
}

func (profiles *Profiles) WithEmail(email string) (bool, error) {
	count, err := profiles.collection.CountDocuments(context.Background(), bson.D{{"email", email}})
	if err != nil {
		return false, fmt.Errorf("error occured: %s", err)
	}

	return count > 0, nil
}

func (profiles *Profiles) Update(ctx context.Context, profile *model.Profile) error {
	res, err := profiles.collection.UpdateByID(ctx, profile.ID, bson.D{{"$set", profile}})
	if err != nil {
		return ErrUpdateProfile
	}

	if res.MatchedCount != 1 || res.ModifiedCount != 1 {
		return ErrUpdateProfile
	}

	return nil
}
