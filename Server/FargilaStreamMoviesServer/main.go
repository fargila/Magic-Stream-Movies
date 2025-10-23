package main

import (
	"fmt"
	"log"

	controller "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/controllers"
	database "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/database"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {

	if database.Client == nil {
		log.Fatal("MongoDB client was not initialized")
	}
	router := gin.Default()
	if err := router.SetTrustedProxies([]string{"192.168.0.0/16"}); err != nil {
		log.Fatal("Error while configurating trustfull proxies:", err)
	}

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, FargilaStreamMovies!")
	})

	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
