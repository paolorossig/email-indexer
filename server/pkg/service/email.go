package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/paolorossig/email-indexer/pkg/adapter/zincsearch"
	"github.com/paolorossig/email-indexer/pkg/domain"
)

const (
	emailDetailSeparator         = "\r\n"
	emailDetailsContentSeparator = "\r\n\r\n"
	defaultEmailSearchType       = "matchphrase"
	defaultEmailMaxResults       = 5
)

// EmailService is the interface for the EmailService
type EmailService struct {
	zincsearchAdapter ZincSearchAdapter
}

// NewEmailService creates a new EmailService
func NewEmailService(zsa ZincSearchAdapter) *EmailService {
	return &EmailService{
		zincsearchAdapter: zsa,
	}
}

// GetAvailableUsers returns the User IDs available
func (es *EmailService) GetAvailableUsers() ([]string, error) {
	var records []string

	files, err := os.ReadDir(domain.EmailsRootFolder)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			records = append(records, f.Name())
		}
	}

	return records, nil
}

// ProcessEmailFile processes an email file and returns an Email
func (es *EmailService) ProcessEmailFile(filepathString string) (*domain.Email, error) {
	file, err := os.ReadFile(filepath.Clean(filepathString))
	if err != nil {
		log.Println("Error in ProcessEmailFile - reading file: ", err)
		return nil, err
	}

	arr := strings.SplitN(string(file), emailDetailsContentSeparator, 2)
	if len(arr) != 2 {
		log.Println("Error in ProcessEmailFile - Wrong email file found at: ", filepathString)
		return nil, err
	}

	allDetails, content := arr[0], arr[1]

	detailsArr := strings.Split(allDetails, emailDetailSeparator)

	email := mapEmailDetails(detailsArr)
	email.Content = content
	email.Filepath = filepathString

	return email, nil

}

func mapEmailDetails(details []string) *domain.Email {
	email := &domain.Email{}

	for i := 0; i < len(details); i++ {
		keyValue := strings.SplitN(details[i], ": ", 2)
		switch keyValue[0] {
		case "Message-ID":
			email.MessageID = keyValue[1]
		case "Date":
			email.Date = keyValue[1]
		case "From":
			email.From = keyValue[1]
		case "To":
			email.To = keyValue[1]
		case "Subject":
			email.Subject = keyValue[1]
		default:
			continue
		}
	}

	return email
}

// ExtrackEmailsFromUser extracts emails from a user
func (es *EmailService) ExtrackEmailsFromUser(userID string) ([]domain.Email, error) {
	var emails []domain.Email

	userFolderPath := fmt.Sprintf("%s/%s", domain.EmailsRootFolder, userID)

	err := filepath.Walk(userFolderPath, es.visitAndProcessEmailFiles(&emails))
	if err != nil {
		log.Println("Error in ExtrackEmailsFromUser: ", err)
	}

	return emails, nil
}

func (es *EmailService) visitAndProcessEmailFiles(emails *[]domain.Email) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error in visitAndProcessEmailFiles: ", err)
		}

		if !info.IsDir() {
			email, err := es.ProcessEmailFile(path)
			if err != nil || email == nil {
				return nil
			}

			*emails = append(*emails, *email)
		}

		return nil
	}
}

// SearchInEmails searches in emails
func (es *EmailService) SearchInEmails(indexName string, term string) ([]domain.Email, error) {
	now := time.Now()
	startTime := now.AddDate(0, 0, -7).Format("2006-01-02T15:04:05Z")
	endTime := now.Format("2006-01-02T15:04:05Z")

	body := zincsearch.SearchDocumentsRequest{
		SearchType: defaultEmailSearchType,
		Query: zincsearch.SearchDocumentsRequestQuery{
			Term:      term,
			StartTime: startTime,
			EndTime:   endTime,
		},
		SortFields: []string{"-@timestamp"},
		From:       0,
		MaxResults: defaultEmailMaxResults,
	}

	response, err := es.zincsearchAdapter.SearchDocuments(indexName, body)
	if err != nil {
		log.Println("Error in SearchInEmails: ", err)
		return nil, err
	}

	return mapZincSearchResponseToEmails(response), nil
}

func mapZincSearchResponseToEmails(response *zincsearch.SearchDocumentsResponse) []domain.Email {
	var emails []domain.Email

	for _, hit := range response.Hits.Hits {
		var email domain.Email

		emailBytes, _ := json.Marshal(hit.Source)

		err := json.Unmarshal(emailBytes, &email)
		if err != nil {
			log.Println("Error in mapZincSearchResponseToEmails: ", err)
			continue
		}

		emails = append(emails, email)
	}

	return emails
}
