package controllers

import (
	"net/http"
	"strconv"

	"go-proj/practice-lab/go-projects/web-service-gin/extensions"
	"go-proj/practice-lab/go-projects/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album id not found"})
}

func DeleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	intId, _ := strconv.Atoi(id)

	remainingAlbums := extensions.Remove(albums, intId)
	c.IndentedJSON(http.StatusOK, remainingAlbums)
}
