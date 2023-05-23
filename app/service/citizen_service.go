package service

import (
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
