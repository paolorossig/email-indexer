package zincsearch

// CreateDocumentsResponse is the response for the CreateDocuments function
type CreateDocumentsResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}

// ErrorReponse is the response of ZincSearch when an error occurs
type ErrorReponse struct {
	ErrorMessage string `json:"error"`
}
