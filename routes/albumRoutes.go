package routes

import (
	"database/sql"

	"github.com/hezronkimutai/goLU/controllers"

	"github.com/gin-gonic/gin"
)

// AlbumRoutes sets up the album routes
func AlbumRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/albums", controllers.GetAlbums(db))
	router.POST("/albums", controllers.AddAlbum(db))
}
