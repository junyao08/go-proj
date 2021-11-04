package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controllers.getAlbums)
	router.POST("/albums", controllers.postAlbums)
	router.GET("/albums/:id", controllers.getAlbumByID)
	router.GET("/albums/delete/:id", controllers.deleteAlbumById)

	router.Run("localhost:8090")
}
