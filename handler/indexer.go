package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/paolorossig/go-challenge/domain"
)

const rootFolder = "./enron_mail_20110402/maildir"

// IndexerService is the interface for the IndexerService
type IndexerService interface {
	IndexFileFromFolder(folder string) ([]*domain.Email, error)
}

// IndexerHandler is the handler for the Indexer
type IndexerHandler struct {
	indexerService IndexerService
}

// NewIndexerHandler creates a new IndexerHandler
func NewIndexerHandler(is IndexerService) *IndexerHandler {
	return &IndexerHandler{
		indexerService: is,
	}
}

// Index is the method that indexes the data
func (ih *IndexerHandler) Index(w http.ResponseWriter, r *http.Request) {
	folderParam := "allen-p/_sent_mail"
	folder := fmt.Sprintf("%s/%s", rootFolder, folderParam)

	records, err := ih.indexerService.IndexFileFromFolder(folder)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, NewErrResponse(err))
		return
	}

	render.JSON(w, r, &IndexResponse{
		Index:   folderParam,
		Records: records,
	})
}

// NewErrResponse is the method that creates a new ErrResponse
func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal Server Error",
		ErrorText:      err.Error(),
	}
}
