package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID            primitive.ObjectID          `bson:"_id"`
	Email         string                      `bson:"email"`
	Name          string                      `bson:"name"`
	Avatar        string                      `bson:"avatar"`
	Password      string                      `bson:"password"`
	Posts         map[primitive.ObjectID]bool `bson:"posts"`
	LikedPosts    map[primitive.ObjectID]bool `bson:"likedPosts"`
	Notifications []primitive.ObjectID        `bson:"notifications"`
}

func NewProfile(name string, avatar string, email string, password string) *Profile {
	return &Profile{
		ID:            primitive.NewObjectID(),
		Email:         email,
		Name:          name,
		Avatar:        avatar,
		Password:      password,
		Posts:         make(map[primitive.ObjectID]bool),
		LikedPosts:    make(map[primitive.ObjectID]bool),
		Notifications: make([]primitive.ObjectID, 0),
	}
}

func (profile *Profile) CreatePost(content string) *Post {
	post := newPost(content, profile.ID)

	profile.Posts[post.ID] = true

	return post
}

func (profile *Profile) Like(postId primitive.ObjectID) (int, *Notification) {
	if _, exists := profile.LikedPosts[postId]; exists {
		delete(profile.LikedPosts, postId)
		return postDislike, nil
	}

	profile.LikedPosts[postId] = true

	return postLike, newNotification(postId, profile.ID)
}

func (profile *Profile) LikedPostIds() []primitive.ObjectID {
	ids := make([]primitive.ObjectID, 0, len(profile.LikedPosts))

	for key, _ := range profile.LikedPosts {
		ids = append(ids, key)
	}

	return ids
}

func (profile *Profile) AddNotification(notification *Notification) {
	profile.Notifications = append(profile.Notifications, notification.Id)
}
