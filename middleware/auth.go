package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_network/auth"
	"social_network/repository"
)

const AuthUserKey = "profile"

func AuthRequired(profiles *repository.Profiles) gin.HandlerFunc {
	return func(c *gin.Context) {
		credentials, err := auth.GetCredentials(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		profile, err := profiles.ByEmail(credentials.Email)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "User was not found",
			})
			return
		}

		err = auth.CheckPasswords(credentials.Password, profile.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Wrong email or password",
			})
			c.Abort()
			return
		}

		c.Set(AuthUserKey, profile)

		c.Next()
	}
}
