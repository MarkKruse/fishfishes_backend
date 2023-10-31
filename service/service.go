package service

import (
	"net/http"

	common "fishfishes_backend/common"
	"github.com/gin-gonic/gin"
)

type Repo interface {
	GetAllSpots(string) common.Fish_spots
	CheckLogin(user common.User) bool
}

type Service struct {
	Repo Repo
}

func NewService(repo Repo) Service {
	return Service{
		Repo: repo,
	}
}

func (s Service) GetAllSpots(c *gin.Context) {
	id := c.Query("userId")
	if len(id) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Found no userID"})
		return
	}
	spots := s.Repo.GetAllSpots(id)
	c.IndentedJSON(http.StatusOK, spots)
}
