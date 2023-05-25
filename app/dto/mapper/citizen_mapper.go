package dto

import (
	"quest/dto"
	"quest/model"
)

func ToCitizenDTO(citizen *model.Citizen) *dto.CitizenDTO {
	citizenDTO := &dto.CitizenDTO{
		ID:         citizen.ID,
		Name:       citizen.Name,
		Address:    citizen.Address,
		Email:      citizen.Email,
		OperatorID: citizen.OperatorID,
	}

	return citizenDTO
}
