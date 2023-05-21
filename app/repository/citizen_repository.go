package repository

import (
	"quest/model"

	"gorm.io/gorm"
)

type CitizenRepository struct {
	db *gorm.DB
}

func NewCitizenRepository(db *gorm.DB) *CitizenRepository {
	return &CitizenRepository{db: db}
}

func (cr *CitizenRepository) GetAllCitizens() ([]*model.Citizen, error) {
	var citizens []*model.Citizen
	err := cr.db.Find(&citizens).Error
	if err != nil {
		return nil, err
	}

	return citizens, nil
}

func (cr *CitizenRepository) GetCitizenByID(id uint) (*model.Citizen, error) {
	var citizen model.Citizen
	err := cr.db.First(&citizen, id).Error
	if err != nil {
		return nil, err
	}

	return &citizen, nil
}

func (cr *CitizenRepository) GetCitizenDocuments(id uint) ([]*model.Document, error) {
	var documents []*model.Document
	err := cr.db.Where("citizen_id = ?", id).Find(&documents).Error
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (cr *CitizenRepository) CreateCitizen(citizen *model.Citizen) error {
	return cr.db.Create(citizen).Error
}

func (cr *CitizenRepository) UpdateCitizen(citizen *model.Citizen) error {
	return cr.db.Save(citizen).Error
}

func (cr *CitizenRepository) DeleteUser(citizen *model.Citizen) error {
	return cr.db.Delete(citizen).Error
}
