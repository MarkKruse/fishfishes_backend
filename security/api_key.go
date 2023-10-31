package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (s Security) ValidateAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.Request.Header.Get("X-API-Key")

		log.Infof("API-KEY: %s", apiKey)

		//TODO ApiKey logic
		if apiKey != "1234" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": 404, "message": "Unauthorized"})
			c.Abort()
			return
		}
		return
		//
	}
}
