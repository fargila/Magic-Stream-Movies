package routes

import (
	controller "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/controllers"

	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/gin-gonic/gin"
)

func SetupUnprotectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.POST("/register", controller.RegisterUser(client))
	router.POST("/login", controller.LoginUser(client))
	router.GET("/movies", controller.GetMovies(client))
}
