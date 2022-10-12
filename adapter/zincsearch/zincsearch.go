package zincsearch

import (
	"fmt"
	"net/http"
	"os"

	"github.com/paolorossig/go-challenge/adapter"
)

const (
	defaultZincSearchHost = "http://localhost:4080"
)

// Client is the client for interacting with the ZincSearch API
type Client struct {
	adapter *adapter.Adapter
}

// NewClient initializes the ZincSearch client.
func NewClient(c *http.Client) *Client {
	host := os.Getenv("ZINCSEARCH_HOST")
	if host == "" {
		host = defaultZincSearchHost
	}

	a := adapter.NewAdapter(c, host)
	setBasicHeaders(a)

	return &Client{
		adapter: a,
	}
}

func setBasicHeaders(a *adapter.Adapter) {
	username := os.Getenv("ZINC_FIRST_ADMIN_USER")
	password := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")

	if username == "" || password == "" {
		panic("ZINC_FIRST_ADMIN_USER and ZINC_FIRST_ADMIN_PASSWORD must be set")
	}

	a.SetBasicAuth(username, password)
}

// CreateDocuments creates documents with the Bulkv2 ZincSearch API
func (c *Client) CreateDocuments(indexName string, records interface{}) (*CreateDocumentsResponse, error) {
	response := &CreateDocumentsResponse{}
	apiError := &ErrorReponse{}

	path := "/api/_bulkv2"
	body := CreateDocumentsRequest{
		Index:   indexName,
		Records: records,
	}

	req, err := c.adapter.BuildRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	res, err := c.adapter.Sling.Do(req, response, apiError)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error creating documents: %s", apiError.ErrorMessage)
	}

	return response, nil
}
