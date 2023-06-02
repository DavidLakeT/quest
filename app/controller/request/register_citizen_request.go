package controller

type RegisterCitizenRequest struct {
	CitizenID  uint   `json:"citizenId"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	OperatorID int    `json:"operatorId"`
}
