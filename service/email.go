package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/paolorossig/go-challenge/domain"
)

const (
	emailDetailSeparator         = "\r\n"
	emailDetailsContentSeparator = "\r\n\r\n"
)

// EmailService is the interface for the EmailService
type EmailService struct{}

// NewEmailService creates a new EmailService
func NewEmailService() *EmailService {
	return &EmailService{}
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
		fmt.Println("Error reading file")
		return nil, err
	}

	arr := strings.SplitN(string(file), emailDetailsContentSeparator, 2)
	if len(arr) != 2 {
		fmt.Printf("Wrong email file found at %s\n", filepathString)
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
		panic(err)
	}

	return emails, nil
}

func (es *EmailService) visitAndProcessEmailFiles(emails *[]domain.Email) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
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
