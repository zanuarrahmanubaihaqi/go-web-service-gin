package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type song struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var songs = []song{
	{ID: "1", Title: "Kursi Goyang", Artist: "Four Twenty", Price: 101},
	{ID: "2", Title: "Ijazah", Artist: "Enau", Price: 99},
	{ID: "3", Title: "Saturnus", Artist: "Soegi Bornean", Price: 86},
}

func main() {
	router := gin.Default()
	router.GET("/songs", getSongs)
	router.POST("/songs", postSongs)

	router.Run("localhost:8090")
}

func getSongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, songs)
}

func postSongs(c *gin.Context) {
    var newSong song

    if err := c.BindJSON(&newSong); err != nil {
        return
    }

    // Add the new song to the slice.
    songs = append(songs, newSong)
    c.IndentedJSON(http.StatusCreated, newSong)
}