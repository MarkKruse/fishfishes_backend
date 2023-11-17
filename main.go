package main

import (
	"context"
	"fishfishes_backend/common/mongo"
	"fishfishes_backend/configuration"
	repo "fishfishes_backend/repository"
	"fishfishes_backend/security"
	"fishfishes_backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"net/http"
	"os"
)

// ------------Security------------
// security.BasicAuthPermission() |
// security.ValidateAPIKey()      |
// --------------------------------

// Init is called right on top of main
func init() {
	// Loads the .env file using godotenv.
	// Throws an error is the file cannot be found.
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	ctx := context.Background()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	config := configuration.NewServiceConfiguration(os.Getenv("MONGOURI"), os.Getenv("MONGODATABASE"), os.Getenv("BACKENDAPIKEY"))

	//Create MongoDB Client
	dbClient, err := mongo.NewMongoDatabase(&config.DB, sugar)
	if err != nil {
		os.Exit(1)
		return
	}

	err = dbClient.Connect(ctx)
	if err != nil {
		os.Exit(1)
		return
	}

	repository := repo.NewRepo(dbClient)
	err = repository.InstallIndexes()
	if err != nil {
		logger.Error(fmt.Sprintf("error installing indexes error:%s", err.Error()))
		os.Exit(1)
		return
	}

	service := service.NewService(repository)
	sec := security.NewSecurity(config.BackendAPIKey) // Add e.g. MongoDB client

	router := gin.Default()
	router.UseH2C = true
	// Add routes
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Did you catch one?")
	})
	router.GET("/version", sec.ValidateAPIKey(), service.Version)
	//Example GET
	router.GET("/getAllSpots", sec.ValidateAPIKey(), service.GetAllSpots)
	router.GET("/getSpotByID", sec.ValidateAPIKey(), service.GetSpotByID)
	router.GET("/getMarkers", sec.ValidateAPIKey(), service.GetAllSpotCoordinates)
	router.GET("/getFishlistSalt", sec.ValidateAPIKey(), service.GetFishListSalt)
	router.GET("/getFishlistFresh", sec.ValidateAPIKey(), service.GetFishListFresh)
	router.PUT("/saveSpot", sec.ValidateAPIKey(), service.SaveSpot)

	//Example POST
	router.POST("/login", sec.ValidateAPIKey(), service.CheckLogin)

	router.PUT("/regist", sec.ValidateAPIKey(), service.CreateAccount)

	router.Run(":8080")
}
