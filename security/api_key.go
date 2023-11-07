package security

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s Security) ValidateAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.Request.Header.Get("X-API-Key")

		if strings.Compare(apiKey, s.APIKey) != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"status": 404, "message": "Unauthorized"})
			c.Abort()
			return
		}
		return
		//
	}
}
