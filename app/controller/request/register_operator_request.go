package controller

type RegisterOperatorRequest struct {
	Name string `json:"name"`
	URL  string `json:"operatorUrl"`
}
