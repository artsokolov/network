package service

import (
	"context"
	"github.com/artsokolov/network/model"
	"github.com/artsokolov/network/repository"
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
