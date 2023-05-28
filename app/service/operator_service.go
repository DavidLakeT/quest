package service

import (
	"quest/model"
	"quest/repository"
)

type OperatorService struct {
	operatorRepository *repository.OperatorRepository
}

func NewOperatorService(operatorRepository *repository.OperatorRepository) *OperatorService {
	return &OperatorService{
		operatorRepository: operatorRepository,
	}
}

func (os *OperatorService) GetOperator(id uint) (*model.Operator, error) {
	return os.operatorRepository.GetOperatorByID(id)
}

func (os *OperatorService) RegisterOperator(operator *model.Operator) error {
	return os.operatorRepository.CreateOperator(operator)
}
