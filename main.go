package main

import (
	"example/web-service-gin/db"
	"example/web-service-gin/models"

	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {

    db.DBConnection()
    db.DB.AutoMigrate(models.Album{})

    router := gin.Default()
    router.GET("/albums", routes.GetAlbums)
    router.GET("/albums/:id", routes.GetAlbumByID)
    router.POST("/albums", routes.PostAlbums)
    router.PATCH("/albums/:id", routes.PatchAlbum)
    router.DELETE("/albums/:id", routes.DeleteAlbum)

    router.Run("localhost:8080")
}