package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(300, gin.H{"message": "List of movies"})
	}
}
