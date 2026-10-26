package mailersend

import (
	"context"
	"net/http"
	"time"
)

const whatsappActivitiesPath = "/whatsapp/activities"

type WhatsAppActivityService interface {
	List(ctx context.Context, options *ListWhatsAppActivityOptions) (*WhatsAppListActivitiesRoot, *Response, error)
}

type whatsAppActivityService struct {
	*service
}

// WhatsAppListActivitiesRoot - format of WhatsApp activities list response
type WhatsAppListActivitiesRoot struct {
	Data  []WhatsAppActivityData `json:"data"`
	Links Links                  `json:"links"`
	Meta  Meta                   `json:"meta"`
}

// WhatsAppActivityData - one status recorded for an item
type WhatsAppActivityData struct {
	ID                string    `json:"id"`
	WhatsAppItemID    string    `json:"whatsapp_item_id"`
	WhatsAppMessageID string    `json:"whatsapp_message_id"`
	WhatsAppAccountID *string   `json:"whatsapp_account_id"`
	From              *string   `json:"from"`
	To                string    `json:"to"`
	Type              string    `json:"type"`
	CreatedAt         time.Time `json:"created_at"`
}

// ListWhatsAppActivityOptions - modifies the behavior of WhatsAppActivityService.List method
type ListWhatsAppActivityOptions struct {
	WhatsAppItemID    string   `url:"whatsapp_item_id,omitempty"`
	WhatsAppMessageID string   `url:"whatsapp_message_id,omitempty"`
	Recipient         string   `url:"recipient,omitempty"`
	Type              []string `url:"type[],omitempty"`
	DateFrom          int64    `url:"date_from,omitempty"`
	DateTo            int64    `url:"date_to,omitempty"`
	Page              int      `url:"page,omitempty"`
	Limit             int      `url:"limit,omitempty"`
}

func (s *whatsAppActivityService) List(ctx context.Context, options *ListWhatsAppActivityOptions) (*WhatsAppListActivitiesRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, whatsappActivitiesPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppListActivitiesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
