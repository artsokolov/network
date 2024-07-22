package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_network/middleware"
	"social_network/model"
)

func (handler *RouteHandler) Notifications(c *gin.Context) {
	profile, ok := c.Get(middleware.AuthUserKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})

		return
	}

	notifications, err := handler.notificationService.List(c.Request.Context(), profile.(*model.Profile))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, notifications)
}
