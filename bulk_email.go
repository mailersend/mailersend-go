package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const bulkEmailBasePath = "/bulk-email"

type BulkEmailService service

type bulkEmailResponse struct {
	Message     string `json:"message"`
	BulkEmailID string `json:"bulk_email_id"`
}

type bulkEmailRoot struct {
	Data bulkEmailData `json:"data"`
}
type bulkEmailData struct {
	ID                        string      `json:"id"`
	State                     string      `json:"state"`
	TotalRecipientsCount      int         `json:"total_recipients_count"`
	SuppressedRecipientsCount int         `json:"suppressed_recipients_count"`
	SuppressedRecipients      interface{} `json:"suppressed_recipients"`
	ValidationErrorsCount     int         `json:"validation_errors_count"`
	ValidationErrors          interface{} `json:"validation_errors"`
	MessagesID                []string    `json:"messages_id"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt                 time.Time   `json:"updated_at"`
}

// Send - send bulk messages
func (s *BulkEmailService) Send(ctx context.Context, message []*Message) (*bulkEmailResponse, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, bulkEmailBasePath, message)
	if err != nil {
		return nil, nil, err
	}

	root := new(bulkEmailResponse)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *BulkEmailService) Status(ctx context.Context, bulkEmailID string) (*bulkEmailRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", bulkEmailBasePath, bulkEmailID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(bulkEmailRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
