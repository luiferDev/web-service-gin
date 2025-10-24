package routes

import (
	"example/web-service-gin/db"
	"example/web-service-gin/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	var albums []models.Album

	// Get all albums from database
	if err := db.DB.Find(&albums).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create album in database
	if err := db.DB.Create(&newAlbum).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Album created successfully", "album": newAlbum})
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	idUUID, err := uuid.Parse(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid UUID format"})
		return
	}

	var album models.Album
	if err := db.DB.First(&album, idUUID).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

// patchAlbum updates an album partially with the provided fields
func PatchAlbum(c *gin.Context) {
	id := c.Param("id")
	idUUID, err := uuid.Parse(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid UUID format"})
		return
	}

	var album models.Album
	if err := db.DB.First(&album, idUUID).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	var updates models.Album
	if err := c.BindJSON(&updates); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&album).Updates(updates).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func DeleteAlbum(c *gin.Context) {
    id := c.Param("id")
    idUUID, err := uuid.Parse(id)

    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid UUID format"})
        return
    }

    var album models.Album
    if err := db.DB.First(&album, idUUID).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

    if err := db.DB.Delete(&album).Error; err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}
