package controllers

import (
	"database/sql"
	"net/http"

	"github.com/hezronkimutai/goLU/models"

	"github.com/gin-gonic/gin"
)

// GetAlbums fetches all albums
func GetAlbums(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, artist, price FROM albums")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch albums"})
			return
		}
		defer rows.Close()

		var albums []models.Album
		for rows.Next() {
			var album models.Album
			if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse album"})
				return
			}
			albums = append(albums, album)
		}
		c.JSON(http.StatusOK, albums)
	}
}

// AddAlbum inserts a new album
func AddAlbum(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newAlbum models.Album
		if err := c.ShouldBindJSON(&newAlbum); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		_, err := db.Exec("INSERT INTO albums (id, title, artist, price) VALUES (?, ?, ?, ?)",
			newAlbum.ID, newAlbum.Title, newAlbum.Artist, newAlbum.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add album"})
			return
		}

		c.JSON(http.StatusCreated, newAlbum)
	}
}
