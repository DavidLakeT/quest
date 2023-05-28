package controller

type TransferCitizenRequest struct {
	CitizenID         uint `json:"citizenId"`
	CurrentOperatorID uint `json:"currentOperatorId"`
	NewOperatorID     uint `json:"newOperatorId"`
}
