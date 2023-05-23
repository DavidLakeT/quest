package routes

import (
	controller "quest/controller/api"

	"github.com/gin-gonic/gin"
)

func RegisterCitizenRoutes(ctx *gin.Engine, controller *controller.CitizenController) {
	apis := ctx.Group("/apis")

	apis.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})

	apis.POST("/registerCitizen", controller.RegisterCitizen)
	apis.GET("/validateCitizen/:id", controller.ValidateCitizen)
}
