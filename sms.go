package mailersend

import (
	"context"
	"net/http"
)

const smsBasePath = "/sms"

type SMSService service

type SMS struct {
	From string   `json:"from"`
	To   []string `json:"to"`
	Text string   `json:"text"`
}

// NewMessage - Setup a new email message ready to be sent.
func (s *SMSService) NewMessage() *SMS {
	return &SMS{}
}

// SetFrom - Set from.
func (m *SMS) SetFrom(from string) {
	m.From = from
}

// SetTo - Set to.
func (m *SMS) SetTo(to []string) {
	m.To = to
}

// SetText - Set the text content of the email, required if not using a template.
func (m *SMS) SetText(text string) {
	m.Text = text
}

// Send - send the message
func (s *SMSService) Send(ctx context.Context, sms *SMS) (*Response, error) {
	req, err := s.client.newRequest(http.MethodPost, smsBasePath, sms)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
