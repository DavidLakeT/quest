package service

import "quest/repository"

type OperatorService struct {
	operatorRepository *repository.OperatorRepository
}

func NewOperatorService(operatorRepository *repository.OperatorRepository) *OperatorService {
	return &OperatorService{
		operatorRepository: operatorRepository,
	}
}
