package service

import (
	"context"
	"net/http"
	"strings"

	common "fishfishes_backend/common"
	"github.com/gin-gonic/gin"
)

type Repo interface {
	CheckLogin(ctx context.Context, user common.User) (bool, string)
	CreateAccount(ctx context.Context, user common.User) error
	SaveSpot(ctx context.Context, userId string, spot common.Fish_spot) error
	GetAllSpots(ctx context.Context, id string) (*[]common.Fish_spot, error)
}

const VERSION string = "0.0.1"

type Service struct {
	Repo Repo
}

func NewService(repo Repo) Service {
	return Service{
		Repo: repo,
	}
}

func (s Service) Version(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, VERSION)
}

func (s Service) GetAllSpots(c *gin.Context) {
	id := c.Query("userId")
	if len(id) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Found no userID"})
		return
	}
	spots, err := s.Repo.GetAllSpots(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, spots)
}

func (s Service) GetAllSpotCoordinates(c *gin.Context) {
	id := c.Query("userId")
	if len(id) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Found no userID"})
		return
	}
	spots, err := s.Repo.GetAllSpots(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var markers []common.Marker
	for _, spot := range *spots {
		markers = append(markers, spot.Marker)
	}

	c.IndentedJSON(http.StatusOK, markers)
}

func (s Service) GetSpotByID(c *gin.Context) {
	userId := c.Query("userId")
	spotId := c.Query("spotId")
	if len(userId) == 0 || len(spotId) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Found no userID or spotId"})
		return
	}
	spots, err := s.Repo.GetAllSpots(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, spot := range *spots {
		if strings.Compare(spot.Id, spotId) == 0 {
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
