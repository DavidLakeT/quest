package service

import (
	"errors"
	"quest/model"
	"quest/repository"

	"gorm.io/gorm"
)

type DocumentService struct {
	documentRepository *repository.DocumentRepository
}

func NewDocumentService(documentRepository *repository.DocumentRepository) *DocumentService {
	return &DocumentService{
		documentRepository: documentRepository,
	}
}

func (ds *DocumentService) RegisterDocument(document *model.Document) error {
	switch ds.documentRepository.CreateDocument(document) {
	case gorm.ErrInvalidField:
		return errors.New("Mensaje de error #1")
	case gorm.ErrDuplicatedKey:
		return errors.New("Mensaje de error #2")
	default:
		return errors.New("Error no identificado")
	}
}

func (ds *DocumentService) GetDocument(title string) (*model.Document, error) {
	document, err := ds.documentRepository.GetDocumentByTitle(title)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.New("Mensaje de error #1")
		case gorm.ErrInvalidField:
			return nil, errors.New("Mensaje de error #2")
		case gorm.ErrDuplicatedKey:
			return nil, errors.New("Mensaje de error #3")
		default:
			return nil, errors.New("Error no identificado")
		}
	}

	return document, nil
}
