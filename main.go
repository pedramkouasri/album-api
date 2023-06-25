package main

import (
	"fmt"

	"github.com/pedramkousari/album-api/services"
)

var albums services.Albums = services.Albums{
	Items: []services.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	},
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
}
