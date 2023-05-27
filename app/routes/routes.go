package routes

import (
	controller "quest/controller/api"

	"github.com/gin-gonic/gin"
)

func RegisterCitizenRoutes(ctx *gin.Engine, controller *controller.CitizenController) {
	apis := ctx.Group("/apis/citizen")

	apis.POST("/registerCitizen", controller.RegisterCitizen)
	apis.GET("/validateCitizen/:id", controller.ValidateCitizen)
	apis.POST("/transferCitizen", controller.TransferCitizen)
	apis.GET("/getCitizenDocuments/:id", controller.GetCitizenDocuments)
}

func RegisterDocumentRoutes(ctx *gin.Engine, controller *controller.DocumentController) {
	apis := ctx.Group("/apis/document")

	apis.POST("/uploadDocument", controller.UploadDocument)
	apis.POST("/authenticateDocument", controller.AuthenticateDocument)
}

func RegisterOperatorRoutes(ctx *gin.Engine, controller *controller.OperatorController) {
	apis := ctx.Group("/apis/operator")

	apis.POST("/registerOperator", controller.RegisterOperator)
}
