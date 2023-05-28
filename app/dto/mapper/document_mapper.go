package dto

import (
	"quest/dto"
	"quest/model"
)

func ToDocumentDTO(document *model.Document) *dto.DocumentDTO {
	documentDTO := &dto.DocumentDTO{
		URL:       document.URL,
		Title:     document.Title,
		Validated: document.Validated,
		CitizenID: document.CitizenID,
	}

	return documentDTO
}

func ToDocumentDTOArray(documents []*model.Document) []*dto.DocumentDTO {
	var result []*dto.DocumentDTO

	for _, document := range documents {
		result = append(result, ToDocumentDTO(document))
	}

	return result
}
