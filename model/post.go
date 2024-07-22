package model

import "go.mongodb.org/mongo-driver/bson/primitive"

var (
	postLike    = 1
	postDislike = -1
)

type Post struct {
	ID         primitive.ObjectID `bson:"_id"`
	Content    string             `bson:"content"`
	Author     primitive.ObjectID `bson:"author"`
	LikesCount int                `bson:"likesCount"`
}

func newPost(content string, authorId primitive.ObjectID) *Post {
	return &Post{
		ID:         primitive.NewObjectID(),
		Content:    content,
		Author:     authorId,
		LikesCount: 0,
	}
}
