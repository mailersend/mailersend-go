package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const whatsappItemsPath = "/whatsapp/items"

type WhatsAppItemService interface {
	List(ctx context.Context, options *ListWhatsAppItemOptions) (*WhatsAppListItemsRoot, *Response, error)
	Get(ctx context.Context, whatsAppItemID string) (*WhatsAppSingleItemRoot, *Response, error)
}

type whatsAppItemService struct {
	*service
}

// WhatsAppListItemsRoot - format of WhatsApp items list response
type WhatsAppListItemsRoot struct {
	Data  []WhatsAppItemData `json:"data"`
	Links Links              `json:"links"`
	Meta  Meta               `json:"meta"`
}

// WhatsAppSingleItemRoot - format of single WhatsApp item response
type WhatsAppSingleItemRoot struct {
	Data WhatsAppItemData `json:"data"`
}

// WhatsAppItemData - the delivery of a WhatsApp message to one recipient.
// Recipient and Activity are only returned by WhatsAppItemService.Get
type WhatsAppItemData struct {
	ID                string                 `json:"id"`
	WhatsAppMessageID string                 `json:"whatsapp_message_id"`
	WhatsAppAccountID *string                `json:"whatsapp_account_id"`
	From              *string                `json:"from"`
	To                string                 `json:"to"`
	Status            string                 `json:"status"`
	TemplateID        *string                `json:"template_id"`
	TemplateName      string                 `json:"template_name"`
	ErrorCode         *int                   `json:"error_code"`
	ErrorMessage      *string                `json:"error_message"`
	CreatedAt         time.Time              `json:"created_at"`
	Recipient         *WhatsAppRecipientData `json:"recipient,omitempty"`
	Activity          []WhatsAppItemActivity `json:"activity,omitempty"`
}

// WhatsAppItemActivity - a status the item moved through
type WhatsAppItemActivity struct {
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// ListWhatsAppItemOptions - modifies the behavior of WhatsAppItemService.List method
type ListWhatsAppItemOptions struct {
	Recipient         string   `url:"recipient,omitempty"`
	WhatsAppMessageID string   `url:"whatsapp_message_id,omitempty"`
	Status            []string `url:"status[],omitempty"`
	WhatsAppAccountID string   `url:"whatsapp_account_id,omitempty"`
	From              string   `url:"from,omitempty"`
	TemplateID        string   `url:"template_id,omitempty"`
	DateFrom          int64    `url:"date_from,omitempty"`
	DateTo            int64    `url:"date_to,omitempty"`
	Page              int      `url:"page,omitempty"`
	Limit             int      `url:"limit,omitempty"`
}

func (s *whatsAppItemService) List(ctx context.Context, options *ListWhatsAppItemOptions) (*WhatsAppListItemsRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, whatsappItemsPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppListItemsRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *whatsAppItemService) Get(ctx context.Context, whatsAppItemID string) (*WhatsAppSingleItemRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", whatsappItemsPath, whatsAppItemID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppSingleItemRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
