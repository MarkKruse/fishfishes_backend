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

	found, userId := s.Repo.CheckLogin(c, user)

	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong input"})
		return
	}

	c.IndentedJSON(http.StatusOK, userId)
	return
}

func (s Service) CreateAccount(c *gin.Context) {

	// Get updated user data from request body
	var userData common.User
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Repo.CreateAccount(c, userData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return response with updated user data
	c.JSON(http.StatusOK, gin.H{"user": userData})
}
