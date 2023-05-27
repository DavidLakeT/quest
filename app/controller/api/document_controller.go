package controller

import (
	"net/http"
	controller "quest/controller/request"
	"quest/model"
	"quest/service"

	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	documentService *service.DocumentService
}

func NewDocumentController(documentService *service.DocumentService) *DocumentController {
	return &DocumentController{
		documentService: documentService,
	}
}

func (dc *DocumentController) UploadDocument(ctx *gin.Context) {
	var request controller.UploadDocumentRequest

	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	// Aquí debes agregar la verificación de los campos de la solicitud

	document := &model.Document{
		URL:       request.DocumentUrl,
		Title:     request.DocumentTitle,
		Validated: false,
		CitizenID: request.CitizenID,
	}

	err = dc.documentService.RegisterDocument(document)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Document successfully created"})
}

func (dc *DocumentController) AuthenticateDocument(ctx *gin.Context) {
	var request controller.AuthenticateDocumentRequest

	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	// Aquí debes agregar la verificación de los campos de la solicitud

	err = dc.documentService.AuthenticateDocument(request.CitizenID, request.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Document succesfully authenticated"})
}
