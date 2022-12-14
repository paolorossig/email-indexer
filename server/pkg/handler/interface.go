package handler

import (
	"github.com/paolorossig/email-indexer/pkg/domain"
)

// EmailService is the interface for the EmailService
type EmailService interface {
	GetAvailableUsers() ([]string, error)
	ExtrackEmailsFromUser(userID string) ([]domain.Email, error)
	SearchInEmails(indexName string, term string) ([]domain.Email, error)
}

// IndexerService is the interface for the IndexerService
type IndexerService interface {
	IndexEmails(indexName string, records []domain.Email) error
}
