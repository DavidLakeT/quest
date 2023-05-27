package repository

import (
	"quest/model"

	"gorm.io/gorm"
)

type OperatorRepository struct {
	db *gorm.DB
}

func NewOperatorRepository(db *gorm.DB) *OperatorRepository {
	return &OperatorRepository{db: db}
}

func (or *OperatorRepository) GetAllOperators() ([]*model.Operator, error) {
	var operators []*model.Operator
	err := or.db.Find(&operators).Error
	if err != nil {
		return nil, err
	}

	return operators, nil
}

func (or *OperatorRepository) GetOperatorByID(id uint) (*model.Operator, error) {
	var operator model.Operator
	err := or.db.First(&operator, id).Error
	if err != nil {
		return nil, err
	}

	return &operator, nil
}

func (or *OperatorRepository) CreateOperator(operator *model.Operator) error {
	return or.db.Create(operator).Error
}
