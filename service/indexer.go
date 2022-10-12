package service

import (
	"log"

	"github.com/paolorossig/go-challenge/adapter/zincsearch"
)

// ZincSearchAdapter is the adapter for interacting with the ZincSearch API
type ZincSearchAdapter interface {
	CreateDocuments(indexName string, records interface{}) (*zincsearch.CreateDocumentsResponse, error)
}

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
func (is *IndexerService) IndexEmails(indexName string, records interface{}) error {
	res, err := is.zincsearchAdapter.CreateDocuments(indexName, records)
	if err != nil {
		return err
	}

	log.Printf("Indexed %d documents\n", res.RecordCount)

	return nil
}
