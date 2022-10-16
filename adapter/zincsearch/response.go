package zincsearch

// CreateDocumentsResponse is the response for the CreateDocuments function
type CreateDocumentsResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}

// SearchDocumentsResponse is the response for the SearchDocuments function
type SearchDocumentsResponse struct {
	Hits struct {
		Hits []struct {
			ID        string                 `json:"_id"`
			Timestamp string                 `json:"@timestamp"`
			Score     float64                `json:"_score"`
			Source    map[string]interface{} `json:"_source"`
		} `json:"hits"`
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
	} `json:"hits"`
	TimedOut bool    `json:"timed_out"`
	Took     float64 `json:"took"`
}

// ErrorReponse is the response of ZincSearch when an error occurs
type ErrorReponse struct {
	ErrorMessage string `json:"error"`
}
