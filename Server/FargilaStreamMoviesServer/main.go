package main

import (
	"fmt"

	controller "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	//fmt.Println("Hello, FargilaStreamMoviesServer!")

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, FargilaStreamMovies!")
	})

	router.GET("/movies", controller.GetMovies())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
