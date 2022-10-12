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

// EmailService is the interface for the EmailService
type EmailService struct{}

// NewEmailService creates a new EmailService
func NewEmailService() *EmailService {
	return &EmailService{}
}

// GetFileNamesInFolder returns the file names in a folder
func (es *EmailService) GetFileNamesInFolder(folder string) ([]string, error) {
	var records []string

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		records = append(records, f.Name())
	}

	return records, nil
}

// ExtrackEmailsFromFilesInFolder extracts emails from files in a folder
func (es *EmailService) ExtrackEmailsFromFilesInFolder(folder string) ([]domain.Email, error) {
	var records []domain.Email

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		filePathString := fmt.Sprintf("%s/%s", folder, f.Name())
		file, err := ioutil.ReadFile(filepath.Clean(filePathString))
		if err != nil {
			return nil, err
		}

		arr := strings.Split(string(file), emailDetailsContentSeparator)
		allDetails, content := arr[0], arr[1]

		detailsArr := strings.Split(allDetails, emailDetailSeparator)

		email := mapEmailDetails(detailsArr)
		email.Content = content
		email.Filepath = filePathString

		records = append(records, *email)
	}

	return records, nil
}

func mapEmailDetails(details []string) *domain.Email {
	email := &domain.Email{}

	for i := 0; i < len(details); i++ {
		keyValue := strings.Split(details[i], ": ")
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
	var records []domain.Email

	userFolderPath := fmt.Sprintf("%s/%s", domain.EmailsRootFolder, userID)
	innerFolders, err := es.GetFileNamesInFolder(userFolderPath)
	if err != nil {
		return nil, err
	}

	for _, innerFolder := range innerFolders {
		filePath := fmt.Sprintf("%s/%s", userFolderPath, innerFolder)
		emails, err := es.ExtrackEmailsFromFilesInFolder(filePath)
		if err != nil {
			continue
		}

		records = append(records, emails...)
	}

	return records, nil
}
