package service

import (
	"errors"
	"quest/model"
	"quest/repository"
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
	return cs.citizenRepository.CreateCitizen(citizen)
}

func (cs *CitizenService) GetCitizen(id uint) (*model.Citizen, error) {
	return cs.citizenRepository.GetCitizenByID(id)
}

func (cs *CitizenService) GetCitizenDocuments(id uint) ([]*model.Document, error) {
	return cs.citizenRepository.GetCitizenDocuments(id)
}

func (cs *CitizenService) TransferCitizen(citizenID uint, currentOperatorID uint, newOperatorID uint) error {
	citizen, err := cs.GetCitizen(citizenID)
	if err != nil {
		return err
	}

	if currentOperatorID == newOperatorID {
		return errors.New("Current and New operator are the same.")
	}

	citizen.OperatorID = int(newOperatorID)
	err = cs.citizenRepository.UpdateCitizen(citizen)
	if err != nil {
		return err
	}

	return nil
}
