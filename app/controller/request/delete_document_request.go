package controller

type DeleteDocumentRequest struct {
	CitizenID uint   `json:"citizenId"`
	Title     string `json:"title"`
}
