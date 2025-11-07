package main

import (
	"fmt"
	"log"

	database "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/database"
	routes "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/routes"

	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {

	router := gin.Default()
	if err := router.SetTrustedProxies([]string{"192.168.0.0/16"}); err != nil {
		log.Fatal("Error while configurating trustfull proxies:", err)
	}

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, FargilaStreamMovies!")
	})

	var client *mongo.Client = database.Connect()

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to reach server: %v", err)
	}

	routes.SetupUnprotectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
