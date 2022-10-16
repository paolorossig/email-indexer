package domain

const (
	// EmailsRootFolder is the root folder for the emails
	EmailsRootFolder = "./enron_mail_20110402/maildir"
	// EmailIndexName is the name of the index for the emails
	EmailIndexName = "emails"
)

// Email is the struct for the email
type Email struct {
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
	Filepath  string `json:"filepath"`
}
