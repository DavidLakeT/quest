package controller

import (
	"net/http"
	controller "quest/controller/request"
	"quest/model"
	"quest/service"
	"regexp"
	"strings"

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

	if request.CitizenID < 10000000 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The citizen id must be at least 8 digits long"})
		return
	}

	urlRegex := regexp.MustCompile(`^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$`)
	if !urlRegex.MatchString(request.DocumentUrl) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid url format"})
		return
	}

	titleRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+.[a-zA-Z0-9]+$`)
	if !titleRegex.MatchString(strings.TrimSpace(request.DocumentTitle)) || len(strings.TrimSpace(request.DocumentTitle)) < 5 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Document title length must be larger than 5"})
		return
	}

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

	if request.CitizenID < 10000000 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The citizen id must be at least 8 digits long"})
		return
	}

	urlRegex := regexp.MustCompile(`^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$`)
	if !urlRegex.MatchString(request.DocumentUrl) {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid url format"})
		return
	}

	titleRegex := regexp.MustCompile(`^[^~)('!*<>:;,?"*|/]+$`)
	if !titleRegex.MatchString(strings.TrimSpace(request.DocumentTitle)) || len(strings.TrimSpace(request.DocumentTitle)) < 5 {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Document title length must be larger than 5"})
		return
	}

	err = dc.documentService.AuthenticateDocument(request.CitizenID, request.DocumentTitle)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Document succesfully authenticated"})
}

func (dc *DocumentController) DeleteDocument(ctx *gin.Context) {
	var request controller.DeleteDocumentRequest

	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = dc.documentService.DeleteDocument(request.CitizenID, request.Title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Document succesfully deleted"})
}

func (dc *DocumentController) UpdateDocument(ctx *gin.Context) {
	var request controller.UpdateDocumentRequest

	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	document := &model.Document{
		URL:       request.DocumentUrl,
		Title:     request.DocumentTitle,
		Validated: false,
		CitizenID: request.CitizenID,
	}

	err = dc.documentService.UpdateDocument(document)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Document succesfully updated"})
}
