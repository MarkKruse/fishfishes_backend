package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (s Security) BasicAuthPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Basic Authentication credentials from the request
		username, password, hasAuth := c.Request.BasicAuth()
		if !hasAuth {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		log.Infof("Username: %s", username)
		// Fetch the user from the database

		// If all checks pass, set the user ID in the context for future use
		c.Set("username", username)
		c.Set("password", password)
		c.Next()
	}
}
