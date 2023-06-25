package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedramkousari/album-api/services"
)

var albums services.Albums = services.Albums{
	Items: []services.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums.Items)
}

func postAlbums(c *gin.Context) {
	var  album services.Album
	if err := c.BindJSON(&album); err != nil {
		fmt.Println(err)
		return
	}

	albums.Create(album)
	c.IndentedJSON(http.StatusCreated, album)
}

func main() {
	album := services.Album{
		ID:     "4",
		Title:  "AAA",
		Artist: "Pedram KOusari",
		Price:  99.99,
	}
	fmt.Printf("%+v \n", album)

	albums.Create(album)

	// albums = albums[:len(albums)-1] //reove item
	index, err := albums.FindIndex("2")
	if err != nil {
		fmt.Println("not found")
		return
	}

	albums.DeleteByIndex(index)

	fmt.Printf("%+v", albums.Items)

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
