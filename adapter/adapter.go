package adapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dghubble/sling"
)

const (
	headerContentType   = "Content-Type"
	mimeApplicationJSON = "application/json"
)

var errUndefinedMethod = errors.New("undefined method")

// Adapter is the struct needed for interacting with via http.
type Adapter struct {
	client *http.Client
	Sling  *sling.Sling
}

// NewAdapter initializes a new adapter for the provided baseURL.
func NewAdapter(c *http.Client, baseURL string) *Adapter {
	s := sling.New().Client(c).Base(baseURL)

	return &Adapter{
		client: c,
		Sling:  s,
	}
}

// SetHeader sets a header for the adapter.
func (a *Adapter) SetHeader(header, value string) {
	a.Sling.Set(header, value)
}

// SetBasicAuth sets the basic auth for the adapter.
func (a *Adapter) SetBasicAuth(username, password string) {
	a.Sling.SetBasicAuth(username, password)
}

// BuildRequest performs a request to the provided path with the provided method and body
func (a *Adapter) BuildRequest(method string, path string, body interface{}) (*http.Request, error) {
	s, err := a.getRequestSling(method, path)
	if err != nil {
		return &http.Request{}, err
	}

	if err := addBodyToSling(s, body); err != nil {
		return &http.Request{}, err
	}

	req, err := s.Request()
	if err != nil {
		return &http.Request{}, err
	}

	return req, nil
}

func (a *Adapter) getRequestSling(method string, path string) (*sling.Sling, error) {
	switch method {
	case http.MethodPost:
		return a.Sling.New().Post(path), nil
	case http.MethodGet:
		return a.Sling.New().Get(path), nil
	case http.MethodPut:
		return a.Sling.New().Put(path), nil
	case http.MethodDelete:
		return a.Sling.New().Delete(path), nil
	case http.MethodPatch:
		return a.Sling.New().Patch(path), nil
	default:
		return a.Sling, errUndefinedMethod
	}
}

func addBodyToSling(s *sling.Sling, body interface{}) error {
	if body == nil {
		return nil
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return err
	}

	s.Set(headerContentType, mimeApplicationJSON).Body(bytes.NewReader(bodyBytes))

	return nil
}
