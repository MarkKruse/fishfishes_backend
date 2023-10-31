package main

import (
	repo "fishfishes_backend/repository"
	security "fishfishes_backend/security"
	service "fishfishes_backend/service"
	"github.com/gin-gonic/gin"
)

// ------------Security------------
// security.BasicAuthPermission() |
// security.ValidateAPIKey()      |
// --------------------------------

func main() {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	//Create MongoDB Client

	repo := repo.NewRepo()
	service := service.NewService(repo)
	sec := security.NewSecurity() // Add e.g. MongoDB client

	router.GET("/login", sec.BasicAuthPermission())
	//Example GET
	router.GET("/allSpots", service.GetAllSpots)

	//Example POST
	router.POST("/test", sec.ValidateAPIKey(), service.GetAlbums)

	router.Run("localhost:8080")

}
