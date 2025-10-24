package routes

import (
	controller "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/controllers"
	middleware "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())

	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
}
