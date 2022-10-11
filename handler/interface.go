package handler

import "github.com/paolorossig/go-challenge/domain"

// IndexResponse is the response for the Index function
type IndexResponse struct {
	Index   string          `json:"index"`
	Records []*domain.Email `json:"records"`
}

// ErrResponse is the response for the errors
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}
