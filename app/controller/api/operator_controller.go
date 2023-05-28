package controller

import (
	"net/http"
	controller "quest/controller/request"
	"quest/model"
	"quest/service"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			validationErrors := make(map[string]string)
			for _, e := range validationErrs {
				var errorMsg string

				switch e.Field() {
				case "name":
					errorMsg = "Error validating operator name field"
				case "operatorUrl":
					errorMsg = "Error validating operator url field"
				default:
					errorMsg = "Validation error"
				}

				validationErrors[e.Field()] = errorMsg
			}

			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": validationErrors})
			return
		}

		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	nameRegex := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !nameRegex.MatchString(strings.TrimSpace(request.Name)) || len(strings.TrimSpace(request.Name)) < 3 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Name length must be larger than 3"})
		return
	}

	urlRegex := regexp.MustCompile(`/^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$/`)
	if !urlRegex.MatchString(request.URL) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid url format"})
		return
	}

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
