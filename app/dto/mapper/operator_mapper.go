package dto

import (
	"quest/dto"
	"quest/model"
)

func ToOperatorDTO(operator *model.Operator) *dto.OperatorDTO {
	operatorDTO := &dto.OperatorDTO{
		Name: operator.Name,
		URL:  operator.URL,
	}

	return operatorDTO
}
