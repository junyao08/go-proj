package main

import (
	"go-proj/practice-lab/go-projects/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.PostAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.GET("/albums/delete/:id", controllers.DeleteAlbumById)

	router.Run("localhost:8090")
}
