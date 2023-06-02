package service

import (
	"errors"
	"quest/model"
	"quest/repository"
)

type CitizenService struct {
	citizenRepository *repository.CitizenRepository
	operatorService   *OperatorService
}

func NewCitizenService(citizenRepository *repository.CitizenRepository, operatorService *OperatorService) *CitizenService {
	return &CitizenService{
		citizenRepository: citizenRepository,
		operatorService:   operatorService,
	}
}

func (cs *CitizenService) RegisterCitizen(citizen *model.Citizen) error {
	_, err := cs.operatorService.GetOperator(citizen.OperatorID)
	if err != nil {
		return errors.New("Could not find an operator with that ID")
	}
	err = cs.citizenRepository.CreateCitizenOnAPI(citizen)
	if err != nil {
		return err
	}
	return cs.citizenRepository.CreateCitizen(citizen)
}

func (cs *CitizenService) UpdateCitizen(citizen *model.Citizen) error {
	return cs.citizenRepository.UpdateCitizen(citizen)
}

func (cs *CitizenService) DeleteCitizen(id uint) error {
	_, err := cs.citizenRepository.GetCitizenByID(id)
	if err != nil {
		return errors.New("Could not find a citizen with that ID")
	}

	return cs.DeleteCitizen(id)
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

	_, err = cs.operatorService.GetOperator(citizen.OperatorID)
	if err != nil {
		return errors.New("Could not find an operator with the new operator ID")
	}

	citizen.OperatorID = newOperatorID
	err = cs.citizenRepository.UpdateCitizen(citizen)
	if err != nil {
		return err
	}

	return nil
}

func (cc *CitizenService) LoginCitizen(citizenID uint, password string) (error, bool) {
	return cc.citizenRepository.LoginCitizen(citizenID, password)
}
