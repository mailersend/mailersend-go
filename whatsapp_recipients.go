package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const whatsappRecipientsPath = "/whatsapp/recipients"

type WhatsAppRecipientService interface {
	List(ctx context.Context, options *ListWhatsAppRecipientOptions) (*WhatsAppListRecipientsRoot, *Response, error)
	Get(ctx context.Context, whatsAppRecipientID string) (*WhatsAppSingleRecipientRoot, *Response, error)
}

type whatsAppRecipientService struct {
	*service
}

// WhatsAppListRecipientsRoot - format of WhatsApp recipients list response.
// Pagination is simple: there is no total, and Links.Last is always empty, so page until Links.Next is empty.
type WhatsAppListRecipientsRoot struct {
	Data  []WhatsAppRecipientData `json:"data"`
	Links Links                   `json:"links"`
	Meta  Meta                    `json:"meta"`
}

// WhatsAppSingleRecipientRoot - format of single WhatsApp recipient response
type WhatsAppSingleRecipientRoot struct {
	Data WhatsAppRecipientData `json:"data"`
}

// WhatsAppRecipientData - a WhatsApp recipient. Items is only returned by WhatsAppRecipientService.Get
type WhatsAppRecipientData struct {
	ID          string             `json:"id"`
	PhoneNumber *string            `json:"phone_number"`
	Username    *string            `json:"username"`
	Status      string             `json:"status"`
	CreatedAt   time.Time          `json:"created_at"`
	Items       []WhatsAppItemData `json:"items,omitempty"`
}

// ListWhatsAppRecipientOptions - modifies the behavior of WhatsAppRecipientService.List method
type ListWhatsAppRecipientOptions struct {
	Status      string `url:"status,omitempty"`
	PhoneNumber string `url:"phone_number,omitempty"`
	Page        int    `url:"page,omitempty"`
	Limit       int    `url:"limit,omitempty"`
}

func (s *whatsAppRecipientService) List(ctx context.Context, options *ListWhatsAppRecipientOptions) (*WhatsAppListRecipientsRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, whatsappRecipientsPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppListRecipientsRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *whatsAppRecipientService) Get(ctx context.Context, whatsAppRecipientID string) (*WhatsAppSingleRecipientRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", whatsappRecipientsPath, whatsAppRecipientID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppSingleRecipientRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
