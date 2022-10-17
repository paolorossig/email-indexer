package service

import (
	"log"

	"github.com/paolorossig/email-indexer/pkg/domain"
)

// IndexerService is the interface for the IndexerService
type IndexerService struct {
	zincsearchAdapter ZincSearchAdapter
}

// NewIndexerService creates a new IndexerService
func NewIndexerService(zsa ZincSearchAdapter) *IndexerService {
	return &IndexerService{
		zincsearchAdapter: zsa,
	}
}

// IndexEmails indexes emails with the ZincSearch API
func (is *IndexerService) IndexEmails(indexName string, records []domain.Email) error {
	res, err := is.zincsearchAdapter.CreateDocuments(indexName, records)
	if err != nil {
		return err
	}

	log.Printf("Indexed %d documents\n", res.RecordCount)

	return nil
}
