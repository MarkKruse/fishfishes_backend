package service

import (
	"context"
	"net/http"
	"strings"

	common "fishfishes_backend/common"
	"github.com/gin-gonic/gin"
)

type Repo interface {
	GetAllSpots(ctx context.Context, id string) common.Fish_spots
	CheckLogin(ctx context.Context, user common.User) (bool, string)
	CreateAccount(ctx context.Context, user common.User) error
	SaveSpot(ctx context.Context, userId string, spot common.Fish_spot) error
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
	spots := s.Repo.GetAllSpots(c, id)
	c.IndentedJSON(http.StatusOK, spots)
}

func (s Service) GetSpotByID(c *gin.Context) {
	userId := c.Query("userId")
	spotId := c.Query("spotId")
	if len(userId) == 0 || len(spotId) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Found no userID or spotId"})
		return
	}
	spots := s.Repo.GetAllSpots(c, userId)

	for _, spot := range spots.Fish_spots {
		if strings.Compare(spot.ID, spotId) == 0 {
			c.IndentedJSON(http.StatusOK, spot)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No spot found"})
}

func (s Service) SaveSpot(c *gin.Context) {
	id := c.Query("userId")
	if len(id) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Found no userID"})
		return
	}

	var spot common.Fish_spot
	if err := c.BindJSON(&spot); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JSON Parse Error"})
		return
	}

	err := s.Repo.SaveSpot(c, id, spot)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return response with updated user data
	c.JSON(http.StatusOK, gin.H{"status": "saved"})
}
