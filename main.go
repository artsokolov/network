package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"social_network/config"
	"social_network/db"
	"social_network/handlers"
	"social_network/repository"
	"social_network/service"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := db.Connection(conf)
	defer func(conn *db.Conn) {
		if err := conn.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}(conn)

	profiles := repository.NewProfiles(conn.DB.Collection("profiles"))
	posts := repository.NewPosts(conn.DB.Collection("posts"))
	notifications := repository.NewNotifications(conn.DB.Collection("notifications"))

	postsService := service.NewPostService(
		conn.DB.Client(),
		posts,
		profiles,
		notifications)

	notificationService := service.NewNotificationsService(notifications)

	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	handlers.RegisterRoutes(r, profiles, postsService, notificationService)

	r.Run(":8080")
}
