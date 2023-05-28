package service

import (
	"errors"
	"quest/model"
	"quest/repository"
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
	return ds.documentRepository.CreateDocument(document)
}

func (ds *DocumentService) GetDocument(citizenID uint, name string) (*model.Document, error) {
	document, err := ds.documentRepository.GetDocumentByTitle(citizenID, name)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (ds *DocumentService) AuthenticateDocument(citizenID uint, name string) error {
	document, err := ds.documentRepository.GetDocumentByTitle(citizenID, name)
	if err != nil {
		return err
	}

	if document.Validated {
		return errors.New("Document is already authenticated.")
	}

	document.Validated = true
	err = ds.documentRepository.UpdateDocument(document)
	if err != nil {
		return err
	}

	return nil
}
