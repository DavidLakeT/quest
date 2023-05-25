package controller

type RegisterCitizenRequest struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	OperatorID int    `json:"operatorId"`
}
