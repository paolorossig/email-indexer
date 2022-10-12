package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

const (
	emailIndexName = "emails"
)

// IndexerHandler is the handler for the Indexer requests
type IndexerHandler struct {
	indexerService IndexerService
	emailService   EmailService
}

// NewIndexerHandler creates a new IndexerHandler
func NewIndexerHandler(is IndexerService, es EmailService) *IndexerHandler {
	return &IndexerHandler{
		indexerService: is,
		emailService:   es,
	}
}

// IndexEmails is the method that indexes the emails
func (ih *IndexerHandler) IndexEmails(w http.ResponseWriter, r *http.Request) {
	userIDs, err := ih.emailService.GetAvailableUsers()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, NewErrResponse(err))
		return
	}

	// TODO: Waitgroup here!

	for i := 0; i <= 5; i++ {
		userID := userIDs[i]

		emailRecords, err := ih.emailService.ExtrackEmailsFromUser(userID)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, NewErrResponse(err))
			return
		}

		if err := ih.indexerService.IndexEmails(emailIndexName, emailRecords); err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, NewErrResponse(err))
			return
		}
	}

	render.Status(r, http.StatusNoContent)
}
