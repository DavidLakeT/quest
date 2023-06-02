package repository

import (
	"quest/model"

	"gorm.io/gorm"
)

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (dr *DocumentRepository) GetDocumentById(id uint) (*model.Document, error) {
	var document model.Document
	err := dr.db.First(&document, id).Error
	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (dr *DocumentRepository) GetDocumentByTitle(citizenID uint, title string) (*model.Document, error) {
	var document model.Document
	err := dr.db.Where("citizen_id = ? AND title = ?", citizenID, title).First(&document).Error
	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (dr *DocumentRepository) CreateDocument(document *model.Document) error {
	return dr.db.Create(document).Error
}

func (dr *DocumentRepository) UpdateDocument(document *model.Document) error {
	return dr.db.Save(document).Error
}

func (dr *DocumentRepository) DeleteDocument(document *model.Document) error {
	return dr.db.Delete(document).Error
}
