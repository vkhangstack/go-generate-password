package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "Jennifer", Price: 8.95},
	{ID: "2", Title: "Blue Train", Artist: "Jennifer", Price: 8.95},
	{ID: "3", Title: "Blue Train", Artist: "Jennifer", Price: 8.95},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// func deleteAlbum(c *gin.Context) {
// 	var newAlbum album

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	for i, album := range albums {
// 		if album.ID == newAlbum.ID {
// 			albums = append(albums[:i], albums[i+1:]...)
// 			c.IndentedJSON(http.StatusOK, newAlbum)
// 			return
// 		}
// 		i--
// 	}
// 	c.IndentedJSON(http.StatusNotFound, newAlbum)
// }

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
}

func main() {
	r := gin.Default()
	r.GET("/", getAlbums)
	r.GET("/:id", getAlbumsById)
	r.POST("/", addAlbum)
	r.DELETE("/", addAlbum)
	r.Run(":3000")
}
