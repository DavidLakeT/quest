package controller

type RegisterCitizenRequest struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	OperatorID   int    `json:"operatorId"`
	OperatorName string `json:"operatorName"`
}
