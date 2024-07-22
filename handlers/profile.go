package handlers

import (
	"github.com/artsokolov/network/middleware"
	"github.com/artsokolov/network/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (handler *RouteHandler) CreateProfile(c *gin.Context) {
	var req request.CreateProfileRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	exists, _ := handler.profiles.WithEmail(req.Email)
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User with this email already exists",
		})
		return
	}

	profile, err := req.NewProfile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = handler.profiles.Create(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		log.Fatalf("Error occured during the profile creation: %s", err)
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (handler *RouteHandler) Profile(c *gin.Context) {
	profile, ok := c.Get(middleware.AuthUserKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, profile)
}
