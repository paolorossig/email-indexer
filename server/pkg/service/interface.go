package service

import "github.com/paolorossig/email-indexer/pkg/adapter/zincsearch"

// ZincSearchAdapter is the adapter for interacting with the ZincSearch API
type ZincSearchAdapter interface {
	CreateDocuments(indexName string, records interface{}) (*zincsearch.CreateDocumentsResponse, error)
	SearchDocuments(indexName string, body zincsearch.SearchDocumentsRequest) (*zincsearch.SearchDocumentsResponse, error)
}
