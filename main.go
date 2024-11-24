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
	// mailer.SendMail(
	// 	"<h1>Hello from Go!</h1><p>This is a test email.</p>",
	// 	"hezronkimutai600@gmail.com",
	// 	"hezronchelimo.hc@gmail.com",
	// 	"Test Email from Go",
	// )
	// Run the server
	router.Run("localhost:8080")
}
