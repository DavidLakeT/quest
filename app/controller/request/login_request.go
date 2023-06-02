package controller

type LoginRequest struct {
	CitizenID 	uint   `json:"citizenId"`
	Password    string `json:"password"`
}