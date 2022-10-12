package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/paolorossig/go-challenge/domain"
)

// EmailHandler is the handler for the Email requests
type EmailHandler struct {
	emailService EmailService
}

// NewEmailHandler creates a new EmailHandler
func NewEmailHandler(es EmailService) *EmailHandler {
	return &EmailHandler{
		emailService: es,
	}
}

type getAvailableUsers struct {
	Users []string `json:"users"`
}

// GetAvailableUsers is the method that returns the available users
func (eh *EmailHandler) GetAvailableUsers(w http.ResponseWriter, r *http.Request) {
	records, err := eh.emailService.GetFileNamesInFolder(domain.EmailsRootFolder)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, NewErrResponse(err))
		return
	}

	response := &getAvailableUsers{
		Users: records,
	}

	render.JSON(w, r, response)
}

type getEmailsFromUserResponse struct {
	Emails []domain.Email `json:"emails"`
}

// GetEmailsFromUser is the method that returns the emails from a user
func (eh *EmailHandler) GetEmailsFromUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	records, err := eh.emailService.ExtrackEmailsFromUser(userID)
	if err != nil {
		return
	}

	response := &getEmailsFromUserResponse{
		Emails: records,
	}

	render.JSON(w, r, response)
}

// ErrResponse is the response for the errors
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// NewErrResponse is the method that creates a new ErrResponse
func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal Server Error",
		ErrorText:      err.Error(),
	}
}
