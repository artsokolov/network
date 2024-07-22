package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"social_network/model"
	"time"
)

type Notifications struct {
	collection *mongo.Collection
}

func NewNotifications(collection *mongo.Collection) *Notifications {
	return &Notifications{collection: collection}
}

func (notifications *Notifications) Create(ctx context.Context, notification *model.Notification) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err := notifications.collection.InsertOne(ctx, notification)
	if err != nil {
		return err
	}

	return nil
}

func (notifications *Notifications) ByUser(ctx context.Context, notificationIds []primitive.ObjectID) ([]model.Notification, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	notificationsResult := make([]model.Notification, 0)

	cur, err := notifications.collection.Find(ctx, bson.D{{"_id", bson.D{{"$in", notificationIds}}}})
	if err != nil {
		return notificationsResult, err
	}

	for cur.Next(context.TODO()) {
		var post model.Notification
		if err := cur.Decode(&post); err != nil {
			return notificationsResult, err
		}
		notificationsResult = append(notificationsResult, post)
	}

	if err := cur.Err(); err != nil {
		return notificationsResult, err
	}

	return notificationsResult, nil
}
