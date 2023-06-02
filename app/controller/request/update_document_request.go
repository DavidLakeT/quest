package controller

type UpdateDocumentRequest struct {
	CitizenID     uint   `json:"citizenId"`
	DocumentUrl   string `json:"documentUrl"`
	DocumentTitle string `json:"documentTitle"`
}
