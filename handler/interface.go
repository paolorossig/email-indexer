package handler

import "github.com/paolorossig/go-challenge/domain"

// EmailService is the interface for the EmailService
type EmailService interface {
	GetFileNamesInFolder(folder string) ([]string, error)
	ExtrackEmailsFromUser(userID string) ([]domain.Email, error)
}

// IndexerService is the interface for the IndexerService
type IndexerService interface {
	IndexEmails(indexName string, records interface{}) error
}
