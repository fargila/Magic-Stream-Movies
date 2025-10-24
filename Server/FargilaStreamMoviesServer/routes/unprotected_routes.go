package routes

import (
	controller "github.com/fargila/Magic-Stream-Movies/Server/FargilaStreamMoviesServer/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUnprotectedRoutes(router *gin.Engine) {
	router.GET("/movies", controller.GetMovies())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
}
