package handlers

import (
	"github.com/artsokolov/network/middleware"
	"github.com/artsokolov/network/model"
	"github.com/artsokolov/network/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *RouteHandler) CreatePost(c *gin.Context) {
	var req request.CreatePostRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	profile, ok := c.Get(middleware.AuthUserKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong.",
		})
		return
	}

	err := handler.postService.CreatePost(c.Request.Context(), &req, profile.(*model.Profile))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (handler *RouteHandler) Posts(c *gin.Context) {
	profile, ok := c.Get(middleware.AuthUserKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong.",
		})
		return
	}

	posts, err := handler.postService.List(profile.(*model.Profile))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (handler *RouteHandler) LikedPosts(c *gin.Context) {
	profile, ok := c.Get(middleware.AuthUserKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong.",
		})
		return
	}

	posts, err := handler.postService.Liked(profile.(*model.Profile).LikedPostIds())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (handler *RouteHandler) LikePost(c *gin.Context) {
	var req request.Post
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
		})
	}

	profile, ok := c.Get(middleware.AuthUserKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong.",
		})
		return
	}

	err := handler.postService.Like(c.Request.Context(), profile.(*model.Profile), req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "IncreaseLikes isn't ok",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
