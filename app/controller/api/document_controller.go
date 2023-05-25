package controller

import "quest/service"

type DocumentController struct {
	documentService *service.DocumentService
}

func NewDocumentController(documentService *service.DocumentService) *DocumentController {
	return &DocumentController{
		documentService: documentService,
	}
}
