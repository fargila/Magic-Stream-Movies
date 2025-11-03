package routes

import (
	controller "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/controllers"
	middleware "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/middleware"

	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate(client))
	protected.GET("/recommendedmovies", controller.GetRecommendedMovies(client))
	protected.GET("/movie/:imdb_id", controller.GetMovie(client))
	protected.POST("/addmovie", controller.AddMovie(client))
}
