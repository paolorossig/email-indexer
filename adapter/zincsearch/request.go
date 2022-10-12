package zincsearch

// CreateDocumentsRequest is the request for the CreateDocuments function
type CreateDocumentsRequest struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}
