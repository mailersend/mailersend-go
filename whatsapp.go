package mailersend

import (
	"context"
	"net/http"
)

const whatsappSendPath = "/whatsapp/send"

// WhatsAppService defines the interface for sending WhatsApp messages.
type WhatsAppService interface {
	NewMessage() *WhatsAppMessage
	Send(ctx context.Context, message *WhatsAppMessage) (*Response, error)
}

type whatsAppService struct {
	*service
}

// WhatsAppMessage represents a WhatsApp message to be sent.
type WhatsAppMessage struct {
	From            string                    `json:"from"`
	To              []string                  `json:"to"`
	TemplateId      string                    `json:"template_id"`
	Personalization []WhatsAppPersonalization `json:"personalization,omitempty"`
}

// WhatsAppPersonalization holds per-recipient template variable values.
type WhatsAppPersonalization struct {
	To   string                      `json:"to"`
	Data WhatsAppPersonalizationData `json:"data"`
}

// WhatsAppPersonalizationData contains positional variable arrays for each template section.
type WhatsAppPersonalizationData struct {
	Header  []string `json:"header,omitempty"`
	Body    []string `json:"body,omitempty"`
	Buttons []string `json:"buttons,omitempty"`
}

// NewMessage - Setup a new WhatsAppMessage ready to be sent.
func (s *whatsAppService) NewMessage() *WhatsAppMessage {
	return &WhatsAppMessage{}
}

// SetFrom - Set the sender phone number in international format without the + symbol.
func (m *WhatsAppMessage) SetFrom(from string) {
	m.From = from
}

// SetTo - Set the recipient phone numbers in international format without the + symbol.
func (m *WhatsAppMessage) SetTo(to []string) {
	m.To = to
}

// SetTemplateId - Set the approved WhatsApp template ID.
func (m *WhatsAppMessage) SetTemplateId(templateId string) {
	m.TemplateId = templateId
}

// SetPersonalization - Set per-recipient template variable personalization.
func (m *WhatsAppMessage) SetPersonalization(personalization []WhatsAppPersonalization) {
	m.Personalization = personalization
}

// Send - Send the WhatsApp message.
func (s *whatsAppService) Send(ctx context.Context, message *WhatsAppMessage) (*Response, error) {
	req, err := s.client.newRequest(http.MethodPost, whatsappSendPath, message)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
