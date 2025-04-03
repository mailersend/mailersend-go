package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const smsMessagesPath = "/sms-messages"

type SmsMessageService interface {
	List(ctx context.Context, options *ListSmsMessageOptions) (*SmsListMessagesRoot, *Response, error)
	Get(ctx context.Context, smsMessageID string) (*SmsSingleMessagesRoot, *Response, error)
}

type smsMessageService struct {
	*service
}

// SmsListMessagesRoot - format of activity response
type SmsListMessagesRoot struct {
	Data  []SmsMessageData `json:"data"`
	Links Links            `json:"links"`
	Meta  Meta             `json:"meta"`
}

// SmsSingleMessagesRoot - format of activity response
type SmsSingleMessagesRoot SmsMessageRoot

// ListSmsMessageOptions - modifies the behavior of SmsMessagesService.List method
type ListSmsMessageOptions struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func (s *smsMessageService) List(ctx context.Context, options *ListSmsMessageOptions) (*SmsListMessagesRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, smsMessagesPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SmsListMessagesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *smsMessageService) Get(ctx context.Context, smsMessageID string) (*SmsSingleMessagesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsMessagesPath, smsMessageID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SmsSingleMessagesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
