package controller

type UpdateCitizenRequest struct {
	CitizenID  uint   `json:"citizenId"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	OperatorID int    `json:"operatorId"`
}
