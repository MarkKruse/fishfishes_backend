package service

import (
	"fishfishes_backend/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s Service) CheckLogin(c *gin.Context) {
	var user common.User
	if err := c.BindJSON(&user); err != nil {
		// DO SOMETHING WITH THE ERROR
	}

	if !s.Repo.CheckLogin(user) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong input"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Okay"})
	return
}
