package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const whatsappMessagesPath = "/whatsapp/messages"

type WhatsAppMessageService interface {
	List(ctx context.Context, options *ListWhatsAppMessageOptions) (*WhatsAppListMessagesRoot, *Response, error)
	Get(ctx context.Context, whatsAppMessageID string) (*WhatsAppSingleMessageRoot, *Response, error)
}

type whatsAppMessageService struct {
	*service
}

// WhatsAppListMessagesRoot - format of WhatsApp messages list response
type WhatsAppListMessagesRoot struct {
	Data  []WhatsAppMessageData `json:"data"`
	Links Links                 `json:"links"`
	Meta  Meta                  `json:"meta"`
}

// WhatsAppSingleMessageRoot - format of single WhatsApp message response
type WhatsAppSingleMessageRoot struct {
	Data WhatsAppMessageData `json:"data"`
}

// WhatsAppMessageData - a sent WhatsApp message. Items is only returned by WhatsAppMessageService.Get
type WhatsAppMessageData struct {
	ID                string             `json:"id"`
	From              *string            `json:"from"`
	To                []string           `json:"to"`
	WhatsAppAccountID *string            `json:"whatsapp_account_id"`
	TemplateID        *string            `json:"template_id"`
	CreatedAt         time.Time          `json:"created_at"`
	Items             []WhatsAppItemData `json:"items,omitempty"`
}

// ListWhatsAppMessageOptions - modifies the behavior of WhatsAppMessageService.List method
type ListWhatsAppMessageOptions struct {
	WhatsAppAccountID string `url:"whatsapp_account_id,omitempty"`
	From              string `url:"from,omitempty"`
	TemplateID        string `url:"template_id,omitempty"`
	DateFrom          int64  `url:"date_from,omitempty"`
	DateTo            int64  `url:"date_to,omitempty"`
	Page              int    `url:"page,omitempty"`
	Limit             int    `url:"limit,omitempty"`
}

func (s *whatsAppMessageService) List(ctx context.Context, options *ListWhatsAppMessageOptions) (*WhatsAppListMessagesRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, whatsappMessagesPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppListMessagesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *whatsAppMessageService) Get(ctx context.Context, whatsAppMessageID string) (*WhatsAppSingleMessageRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", whatsappMessagesPath, whatsAppMessageID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppSingleMessageRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
