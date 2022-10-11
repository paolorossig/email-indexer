package service

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/paolorossig/go-challenge/domain"
)

const (
	emailDetailSeparator         = "\r\n"
	emailDetailsContentSeparator = "\r\n\r\n"
)

// IndexerService is the interface for the IndexerService
type IndexerService struct{}

// NewIndexerService creates a new IndexerService
func NewIndexerService() *IndexerService {
	return &IndexerService{}
}

// IndexFileFromFolder indexes the files from a folder
func (is *IndexerService) IndexFileFromFolder(folder string) ([]*domain.Email, error) {
	var records []*domain.Email

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		filePath := filepath.Clean(fmt.Sprintf("%s/%s", folder, f.Name()))
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		arr := strings.Split(string(file), emailDetailsContentSeparator)
		allDetails, content := arr[0], arr[1]

		detailsArr := strings.Split(allDetails, emailDetailSeparator)

		email := &domain.Email{
			MessageID: strings.TrimPrefix(detailsArr[0], "Message-ID: "),
			Date:      strings.TrimPrefix(detailsArr[1], "Date: "),
			From:      strings.TrimPrefix(detailsArr[2], "From: "),
			To:        strings.TrimPrefix(detailsArr[3], "To: "),
			Subject:   strings.TrimPrefix(detailsArr[4], "Subject: "),
			Content:   content,
		}

		records = append(records, email)
	}

	return records, nil
}
