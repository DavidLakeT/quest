package service

import (
	"errors"
	"quest/model"
	"quest/repository"

	"gorm.io/gorm"
)

type CitizenService struct {
	citizenRepository *repository.CitizenRepository
}

func NewCitizenService(citizenRepository *repository.CitizenRepository) *CitizenService {
	return &CitizenService{
		citizenRepository: citizenRepository,
	}
}

func (cs *CitizenService) RegisterCitizen(citizen *model.Citizen) error {
	switch cs.citizenRepository.CreateCitizen(citizen) {
	case gorm.ErrRecordNotFound:
		return errors.New("Mensaje de error #1")
	case gorm.ErrInvalidField:
		return errors.New("Mensaje de error #2")
	case gorm.ErrDuplicatedKey:
		return errors.New("Mensaje de error #3")
	default:
		return errors.New("Error no identificado")
	}
}

func (cs *CitizenService) GetCitizen(id uint) (*model.Citizen, error) {
	return cs.citizenRepository.GetCitizenByID(id)
}
