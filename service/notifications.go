package service

import (
	"context"
	"social_network/model"
	"social_network/repository"
)

type Notifications struct {
	notifications *repository.Notifications
}

func NewNotificationsService(notifications *repository.Notifications) *Notifications {
	return &Notifications{notifications: notifications}
}

func (service *Notifications) List(ctx context.Context, profile *model.Profile) ([]model.Notification, error) {
	return service.notifications.ByUser(ctx, profile.Notifications)
}
