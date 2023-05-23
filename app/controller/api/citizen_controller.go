package controller

import (
	"net/http"
	controller "quest/controller/request"
	mapper "quest/dto/mapper"
	"quest/model"
	"quest/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CitizenController struct {
	citizenService *service.CitizenService
}

func NewCitizenController(citizenService *service.CitizenService) *CitizenController {
	return &CitizenController{citizenService: citizenService}
}

func (cc *CitizenController) RegisterCitizen(ctx *gin.Context) {
	var request controller.RegisterCitizenRequest

	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	// Aquí debes agregar la verificación de los campos de la solicitud

	citizen := model.Citizen{
		ID:           uint(request.ID),
		Name:         request.Name,
		Address:      request.Address,
		Email:        request.Email,
		OperatorID:   request.OperatorID,
		OperatorName: request.OperatorName,
		Documents:    []model.Document{},
	}

	err = cc.citizenService.RegisterCitizen(&citizen)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "User successfully registered"})
}

func (cc *CitizenController) ValidateCitizen(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	// Aquí debes agregar la verificación de los campos de la solicitud

	citizen, err := cc.citizenService.GetCitizen(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToCitizenDTO(citizen))
}
