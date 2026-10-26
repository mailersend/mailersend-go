package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const whatsappInboundMessagesPath = "/whatsapp/inbound-messages"

type WhatsAppInboundMessageService interface {
	List(ctx context.Context, options *ListWhatsAppInboundMessageOptions) (*WhatsAppListInboundMessagesRoot, *Response, error)
	Get(ctx context.Context, whatsAppInboundMessageID string) (*WhatsAppSingleInboundMessageRoot, *Response, error)
}

type whatsAppInboundMessageService struct {
	*service
}

// WhatsAppListInboundMessagesRoot - format of WhatsApp inbound messages list response
type WhatsAppListInboundMessagesRoot struct {
	Data  []WhatsAppInboundMessageData `json:"data"`
	Links Links                        `json:"links"`
	Meta  Meta                         `json:"meta"`
}

// WhatsAppSingleInboundMessageRoot - format of single WhatsApp inbound message response
type WhatsAppSingleInboundMessageRoot struct {
	Data WhatsAppInboundMessageData `json:"data"`
}

// WhatsAppInboundMessageData - a WhatsApp message received by one of your phone numbers.
// Meta holds the content of the message and its shape depends on Type, for example
// "body" for text, or "url", "status", "mime_type", "size" and "caption" for media. Meta is nil when there is no content.
// Context is set only when the message is a reply.
type WhatsAppInboundMessageData struct {
	ID                string                  `json:"id"`
	WhatsAppAccountID string                  `json:"whatsapp_account_id"`
	From              string                  `json:"from"`
	To                string                  `json:"to"`
	Type              string                  `json:"type"`
	ReceivedAt        time.Time               `json:"received_at"`
	Meta              map[string]interface{}  `json:"meta"`
	Context           *WhatsAppInboundContext `json:"context,omitempty"`
}

// WhatsAppInboundContext - present when the inbound message is a reply
type WhatsAppInboundContext struct {
	MessageID string `json:"message_id"`
}

// ListWhatsAppInboundMessageOptions - modifies the behavior of WhatsAppInboundMessageService.List method
type ListWhatsAppInboundMessageOptions struct {
	WhatsAppAccountID string   `url:"whatsapp_account_id,omitempty"`
	From              string   `url:"from,omitempty"`
	Type              []string `url:"type[],omitempty"`
	DateFrom          int64    `url:"date_from,omitempty"`
	DateTo            int64    `url:"date_to,omitempty"`
	Page              int      `url:"page,omitempty"`
	Limit             int      `url:"limit,omitempty"`
}

func (s *whatsAppInboundMessageService) List(ctx context.Context, options *ListWhatsAppInboundMessageOptions) (*WhatsAppListInboundMessagesRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, whatsappInboundMessagesPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppListInboundMessagesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *whatsAppInboundMessageService) Get(ctx context.Context, whatsAppInboundMessageID string) (*WhatsAppSingleInboundMessageRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", whatsappInboundMessagesPath, whatsAppInboundMessageID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppSingleInboundMessageRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
