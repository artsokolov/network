package request

import (
	"social_network/auth"
	"social_network/model"
)

type CreateProfileRequest struct {
	Name            string `json:"name" binding:"required,max=12"`
	Avatar          string `json:"avatar" binding:"required,url"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}

func (request *CreateProfileRequest) NewProfile() (*model.Profile, error) {
	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	return model.NewProfile(request.Name, request.Avatar, request.Email, hashedPassword), nil
}
