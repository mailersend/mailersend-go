package mailersend

import (
	"context"
	"net/http"
	"time"
)

const smsBasePath = "/sms"

type SmsService interface {
	NewMessage() *Sms
	Send(ctx context.Context, sms *Sms) (*Response, error)
}

type smsService struct {
	*service
}

type Sms struct {
	From            string               `json:"from"`
	To              []string             `json:"to"`
	Text            string               `json:"text"`
	Personalization []SmsPersonalization `json:"personalization,omitempty"`
}

// SmsPersonalization - you can set multiple SmsPersonalization for each Recipient
type SmsPersonalization struct {
	PhoneNumber string                 `json:"phone_number"`
	Data        map[string]interface{} `json:"data"`
}

type SmsMessageRoot struct {
	Data SmsMessageData `json:"data"`
}

type SmsMessageData struct {
	Id              string            `json:"id"`
	From            string            `json:"from"`
	To              []string          `json:"to"`
	Text            string            `json:"text"`
	Paused          bool              `json:"paused"`
	CreatedAt       time.Time         `json:"created_at"`
	SmsMessage      []SmsMessage      `json:"sms"`
	SmsActivityData []SmsActivityData `json:"sms_activity"`
}

type SmsMessage struct {
	Id               string      `json:"id"`
	From             string      `json:"from"`
	To               string      `json:"to"`
	Text             string      `json:"text"`
	Status           string      `json:"status"`
	SegmentCount     int         `json:"segment_count"`
	ErrorType        interface{} `json:"error_type"`
	ErrorDescription interface{} `json:"error_description"`
}

// NewMessage - Setup a new Sms message ready to be sent.
func (s *smsService) NewMessage() *Sms {
	return &Sms{}
}

// SetFrom - Set from.
func (m *Sms) SetFrom(from string) {
	m.From = from
}

// SetTo - Set to.
func (m *Sms) SetTo(to []string) {
	m.To = to
}

// SetText - Set the text content of the email, required if not using a template.
func (m *Sms) SetText(text string) {
	m.Text = text
}

// SetPersonalization - Set the template personalization.
func (m *Sms) SetPersonalization(personalization []SmsPersonalization) {
	m.Personalization = personalization
}

// Send - send the message
func (s *smsService) Send(ctx context.Context, sms *Sms) (*Response, error) {
	req, err := s.client.newRequest(http.MethodPost, smsBasePath, sms)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
