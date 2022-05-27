package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const smsActivityPath = "/sms-activity"
const smsMessagesPath = "/sms-messages"

type SmsActivityService service

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

// SmsActivityOptions - modifies the behavior of SmsService.Activity method
type SmsActivityOptions struct {
	SmsNumberId string   `url:"sms_number_id,omitempty"`
	Status      []string `url:"status[],omitempty"`
	Page        int      `url:"page,omitempty"`
	DateFrom    int64    `url:"date_from,omitempty"`
	DateTo      int64    `url:"date_to,omitempty"`
	Limit       int      `url:"limit,omitempty"`
}

func (s *SmsActivityService) List(ctx context.Context, options *SmsActivityOptions) (*smsListActivityRoot, *Response, error) {
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

func (s *SmsActivityService) Get(ctx context.Context, smsMessageID string) (*SmsMessageRoot, *Response, error) {
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
