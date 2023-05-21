package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(ctx *gin.Engine) {
	apis := ctx.Group("/apis")

	apis.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola desde /apis/hello"})
	})

	apis.GET("/bye", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Adiós desde /apis/bye"})
	})
}
