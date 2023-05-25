package controller

import "quest/service"

type OperatorController struct {
	operatorService *service.OperatorService
}

func NewOperatorController(operatorService *service.OperatorService) *OperatorController {
	return &OperatorController{
		operatorService: operatorService,
	}
}
