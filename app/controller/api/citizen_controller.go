package controller

import (
	"net/http"
	controller "quest/controller/request"
	mapper "quest/dto/mapper"
	"quest/encryption"
	"quest/model"
	"quest/service"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CitizenController struct {
	citizenService *service.CitizenService
}

func NewCitizenController(citizenService *service.CitizenService) *CitizenController {
	return &CitizenController{
		citizenService: citizenService,
	}
}

func (cc *CitizenController) RegisterCitizen(ctx *gin.Context) {
	var request controller.RegisterCitizenRequest

	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	if request.CitizenID < 10000000 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The ID must be at least 8 digits long"})
		return
	}

	nameRegex := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !nameRegex.MatchString(strings.TrimSpace(request.Name)) || len(strings.TrimSpace(request.Name)) < 3 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Name length must be larger than 3"})
		return
	}

	if len(strings.TrimSpace(request.Address)) < 12 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Address must be larger than 12 characters"})
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	if !emailRegex.MatchString(request.Email) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid email format"})
		return
	}

	passwordRegex := regexp.MustCompile(`^(?=.*\d).{7,}$`)
	if !passwordRegex.MatchString(request.Password) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid password format (+7 characters, 1 number)"})
		return
	}

	citizen := model.Citizen{
		ID:         uint(request.CitizenID),
		Name:       request.Name,
		Address:    request.Address,
		Email:      request.Email,
		OperatorID: uint(request.OperatorID),
		Documents:  []model.Document{},
	}

	err = cc.citizenService.RegisterCitizen(&citizen)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Citizen successfully registered"})
}

func (cc *CitizenController) UpdateCitizen(ctx *gin.Context) {
	var request controller.UpdateCitizenRequest

	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	citizen := model.Citizen{
		ID:         uint(request.CitizenID),
		Name:       request.Name,
		Address:    request.Address,
		Email:      request.Email,
		OperatorID: uint(request.OperatorID),
		Documents:  []model.Document{},
	}

	cc.citizenService.UpdateCitizen(&citizen)
}

func (cc *CitizenController) ValidateCitizen(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	if id < 10000000 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The Citizen ID must be at least 8 digits long"})
		return
	}

	citizen, err := cc.citizenService.GetCitizen(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToCitizenDTO(citizen))
}

func (cc *CitizenController) TransferCitizen(ctx *gin.Context) {
	var request controller.TransferCitizenRequest

	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	if request.CitizenID < 10000000 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The Citizen ID must be at least 8 digits long"})
		return
	}

	err = cc.citizenService.TransferCitizen(request.CitizenID, request.CurrentOperatorID, request.NewOperatorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Citizen succesfully transfered"})
}

func (cc *CitizenController) GetCitizenDocuments(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	if id < 10000000 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The Citizen ID must be at least 8 digits long"})
		return
	}

	documents, err := cc.citizenService.GetCitizenDocuments(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": mapper.ToDocumentDTOArray(documents)})
}

func (cc *CitizenController) DeleteCitizen(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = cc.citizenService.DeleteCitizen(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Citizen succesfully deleted"})
}

func (cc *CitizenController) LoginCitizen(ctx *gin.Context) {
	auth := controller.LoginRequest{}
	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err, authorize := cc.citizenService.LoginCitizen(auth.CitizenID, auth.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	if !authorize {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	token, err := encryption.SignedLoginToken(auth.CitizenID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.SetCookie("auth", token, 90, "", "", false, false)
	ctx.JSON(http.StatusAccepted, nil)
}

func (cc *CitizenController) CheckAuth(ctx *gin.Context) {
	token, err := ctx.Cookie("auth")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, nil)
		ctx.Abort()
		return
	}

	valid, err := encryption.CheckSignedToken(token)
	if err != nil || !valid {
		ctx.JSON(http.StatusUnauthorized, nil)
		ctx.Abort()
		return
	}

	ctx.Next()
}
