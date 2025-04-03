package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const bulkEmailBasePath = "/bulk-email"

type BulkEmailService interface {
	Send(ctx context.Context, message []*Message) (*BulkEmailResponse, *Response, error)
	Status(ctx context.Context, bulkEmailID string) (*BulkEmailRoot, *Response, error)
}

type bulkEmailService struct {
	*service
}

type BulkEmailResponse struct {
	Message     string `json:"message"`
	BulkEmailID string `json:"bulk_email_id"`
}

type BulkEmailRoot struct {
	Data BulkEmailData `json:"data"`
}
type BulkEmailData struct {
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
func (s *bulkEmailService) Send(ctx context.Context, message []*Message) (*BulkEmailResponse, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, bulkEmailBasePath, message)
	if err != nil {
		return nil, nil, err
	}

	root := new(BulkEmailResponse)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *bulkEmailService) Status(ctx context.Context, bulkEmailID string) (*BulkEmailRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", bulkEmailBasePath, bulkEmailID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(BulkEmailRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
