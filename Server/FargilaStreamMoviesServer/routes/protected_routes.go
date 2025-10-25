package routes

import (
	controller "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/controllers"
	middleware "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/movie/:imdb_id", controller.GetMovie())
	protected.POST("/addmovie", controller.AddMovie())
}
