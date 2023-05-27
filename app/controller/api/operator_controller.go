package controller

import (
	"net/http"
	controller "quest/controller/request"
	"quest/model"
	"quest/service"

	"github.com/gin-gonic/gin"
)

type OperatorController struct {
	operatorService *service.OperatorService
}

func NewOperatorController(operatorService *service.OperatorService) *OperatorController {
	return &OperatorController{
		operatorService: operatorService,
	}
}

func (oc *OperatorController) RegisterOperator(ctx *gin.Context) {
	var request controller.RegisterOperatorRequest

	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	// Aquí debes agregar la verificación de los campos de la solicitud

	operator := model.Operator{
		Name:     request.Name,
		URL:      request.URL,
		Citizens: []model.Citizen{},
	}

	err = oc.operatorService.RegisterOperator(&operator)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Operator successfully registered"})
}
