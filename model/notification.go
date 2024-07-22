package model

import "go.mongodb.org/mongo-driver/bson/primitive"

var notificationTypeLike = "like"

type Notification struct {
	Id      primitive.ObjectID `bson:"_id"`
	Type    string             `bson:"type"`
	PostId  primitive.ObjectID `bson:"postId"`
	LikedBy primitive.ObjectID `bson:"likedBy"`
}

func newNotification(postId primitive.ObjectID, profileId primitive.ObjectID) *Notification {
	return &Notification{
		primitive.NewObjectID(),
		notificationTypeLike,
		postId,
		profileId,
	}
}
