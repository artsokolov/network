package handlers

import (
	"github.com/gin-gonic/gin"
	"social_network/middleware"
	"social_network/repository"
	"social_network/service"
)

type RouteHandler struct {
	profiles            *repository.Profiles
	postService         *service.Posts
	notificationService *service.Notifications
}

func RegisterRoutes(
	router *gin.Engine,
	profiles *repository.Profiles,
	postService *service.Posts,
	notificationService *service.Notifications) {
	routeHandler := &RouteHandler{profiles, postService, notificationService}

	api := router.Group("/api")
	{
		api.POST("/profile", routeHandler.CreateProfile)

		requireAuth := api.Group("")
		requireAuth.Use(middleware.AuthRequired(profiles))
		{
			requireAuth.GET("/profile", routeHandler.Profile)
			requireAuth.GET("/notifications", routeHandler.Notifications)

			posts := requireAuth.Group("/posts")
			{
				posts.GET("/", routeHandler.Posts)
				posts.POST("/", routeHandler.CreatePost)

				posts.GET("/liked", routeHandler.LikedPosts)
				posts.POST("/:id/like", routeHandler.LikePost)
			}
		}
	}
}
