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

// global store of albums
var albums = []album{
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	},
	{
		ID:     "3",
		Title:  "Sarag Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func createAlbum(c *gin.Context) {
	var newAlbum album
	// how is invalid input handled?
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	// retrieve request path parameter from the gin context
	id := c.Param("id")

	// find the album with the given id
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// error if the album isn't found
	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{
			"message": "album not found",
		},
	)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", createAlbum)

	router.Run("localhost:8080")
}
