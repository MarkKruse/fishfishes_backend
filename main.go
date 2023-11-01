package main

import (
	"context"
	"fishfishes_backend/common/mongo"
	"fishfishes_backend/configuration"
	repo "fishfishes_backend/repository"
	"fishfishes_backend/security"
	"fishfishes_backend/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

// ------------Security------------
// security.BasicAuthPermission() |
// security.ValidateAPIKey()      |
// --------------------------------

func main() {
	ctx := context.Background()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	router := gin.Default()

	config := configuration.NewServiceConfiguration()
	config.ServiceFlags()
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
	service := service.NewService(repository)
	sec := security.NewSecurity() // Add e.g. MongoDB client

	//router.GET("/login", sec.BasicAuthPermission())
	//Example GET
	router.GET("/getAllSpots", service.GetAllSpots, sec.ValidateAPIKey())
	router.GET("/getSpotByID", service.GetSpotByID, sec.ValidateAPIKey())
	router.PUT("/saveSpot", service.SaveSpot, sec.ValidateAPIKey())

	//Example POST
	router.POST("/login", service.CheckLogin, sec.ValidateAPIKey())

	router.PUT("/regist", service.CreateAccount, sec.ValidateAPIKey())

	router.Run("localhost:8086")
}
