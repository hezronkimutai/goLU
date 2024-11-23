package main

import (
	"github.com/hezronkimutai/goLU/config"
	"github.com/hezronkimutai/goLU/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	db := config.ConnectDB()

	// Set up Gin router
	router := gin.Default()

	// Initialize routes with the database
	routes.AlbumRoutes(router, db)

	// Run the server
	router.Run("localhost:8080")
}
