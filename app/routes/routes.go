package routes

import (
	controller "quest/controller/api"

	"github.com/gin-gonic/gin"
)

func RegisterCitizenRoutes(ctx *gin.Engine, controller *controller.CitizenController) {
	apis := ctx.Group("/apis/citizen")

	apis.POST("/registerCitizen", controller.RegisterCitizen)
	apis.PUT("/updateCitizen", controller.CheckAuth, controller.UpdateCitizen)
	apis.DELETE("/deleteCitizen", controller.DeleteCitizen)
	apis.GET("/validateCitizen/:id", controller.ValidateCitizen)
	apis.POST("/transferCitizen", controller.TransferCitizen)
	apis.GET("/getCitizenDocuments/:id", controller.GetCitizenDocuments)
	apis.POST("/login", controller.LoginCitizen)
}

func RegisterDocumentRoutes(ctx *gin.Engine, controller *controller.DocumentController, citizenController *controller.CitizenController) {
	apis := ctx.Group("/apis/document")

	apis.POST("/uploadDocument", citizenController.CheckAuth, controller.UploadDocument)
	apis.PUT("/updateDocument", citizenController.CheckAuth, controller.UpdateDocument)
	apis.PATCH("/authenticateDocument", controller.AuthenticateDocument)
	apis.DELETE("/deleteDocument", citizenController.CheckAuth, controller.DeleteDocument)
}

func RegisterOperatorRoutes(ctx *gin.Engine, controller *controller.OperatorController) {
	apis := ctx.Group("/apis/operator")

	apis.POST("/registerOperator", controller.RegisterOperator)
}
