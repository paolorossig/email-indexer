package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/render"
	"github.com/paolorossig/email-indexer/pkg/domain"
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
		NewErrResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	var wg sync.WaitGroup

	for _, userID := range userIDs {
		wg.Add(1)
		go ih.indexEmailByUserID(userID, &wg)
	}

	wg.Wait()

	render.Status(r, http.StatusNoContent)
}

func (ih *IndexerHandler) indexEmailByUserID(userID string, wg *sync.WaitGroup) {
	defer wg.Done()
	emailRecords, err := ih.emailService.ExtrackEmailsFromUser(userID)
	if err != nil {
		log.Println("Error extracting emails from user: ", err)
		return
	}

	if err := ih.indexerService.IndexEmails(domain.EmailIndexName, emailRecords); err != nil {
		log.Println("Error indexing emails from user: ", err)
	}
}
