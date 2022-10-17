package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/paolorossig/email-indexer/pkg/domain"
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

type getAvailableUsersResponse struct {
	Users []string `json:"users"`
}

// GetAvailableUsers is the method that returns the available users
func (eh *EmailHandler) GetAvailableUsers(w http.ResponseWriter, r *http.Request) {
	records, err := eh.emailService.GetAvailableUsers()
	if err != nil {
		NewErrResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	response := &getAvailableUsersResponse{
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
		NewErrResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	response := &getEmailsFromUserResponse{
		Emails: records,
	}

	render.JSON(w, r, response)
}

// SearchInEmailsResponse is the response for the SearchInEmails method
type SearchInEmailsResponse struct {
	Emails []domain.Email `json:"emails"`
}

// SearchInEmails is the method that searches in the emails
func (eh *EmailHandler) SearchInEmails(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("q")

	records, err := eh.emailService.SearchInEmails(domain.EmailIndexName, term)
	if err != nil {
		NewErrResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	response := &SearchInEmailsResponse{
		Emails: records,
	}

	render.JSON(w, r, response)
}

// ErrResponse is the response for the errors
type ErrResponse struct {
	Status    int    `json:"status"`          // user-level status message
	ErrorText string `json:"error,omitempty"` // application-level error message, for debugging
}

// NewErrResponse is the method that creates a new ErrResponse
func NewErrResponse(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	errReponse := &ErrResponse{
		Status:    statusCode,
		ErrorText: err.Error(),
	}

	render.Status(r, statusCode)
	render.JSON(w, r, errReponse)
}
