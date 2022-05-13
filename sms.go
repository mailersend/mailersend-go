package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const smsBasePath = "/sms"
const smsActivityPath = "/sms-activity"
const smsMessagesPath = "/sms-messages"

type SmsService service

type Sms struct {
	From string   `json:"from"`
	To   []string `json:"to"`
	Text string   `json:"text"`
}

// smsListActivityRoot - format of activity response
type smsListActivityRoot struct {
	Data  []smsActivityData `json:"data"`
	Links Links             `json:"links"`
	Meta  Meta              `json:"meta"`
}

// smsActivityData - format of sms activity data
type smsActivityData struct {
	From         string    `json:"from"`
	To           string    `json:"to"`
	CreatedAt    time.Time `json:"created_at"`
	Status       string    `json:"status"`
	SmsMessageId string    `json:"sms_message_id"`
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
	SmsActivityData []smsActivityData `json:"sms_activity"`
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

// SmsActivityOptions - modifies the behavior of SmsService.Activity method
type SmsActivityOptions struct {
	SmsNumberId string   `url:"sms_number_id,omitempty"`
	Status      []string `url:"status[],omitempty"`
	Page        int      `url:"page,omitempty"`
	DateFrom    int64    `url:"date_from,omitempty"`
	DateTo      int64    `url:"date_to,omitempty"`
	Limit       int      `url:"limit,omitempty"`
}

// NewMessage - Setup a new Sms message ready to be sent.
func (s *SmsService) NewMessage() *Sms {
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

// Send - send the message
func (s *SmsService) Send(ctx context.Context, sms *Sms) (*Response, error) {
	req, err := s.client.newRequest(http.MethodPost, smsBasePath, sms)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

func (s *SmsService) ListActivity(ctx context.Context, options *SmsActivityOptions) (*smsListActivityRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, smsActivityPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(smsListActivityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SmsService) Activity(ctx context.Context, smsMessageID string) (*SmsMessageRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsMessagesPath, smsMessageID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SmsMessageRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
