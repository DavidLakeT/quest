package controller

type TransferCitizenRequest struct {
	CitizenID         uint `json:"citizenID"`
	CurrentOperatorID uint `json:"currentOperatorID"`
	NewOperatorID     uint `json:"newOperatorID"`
}
