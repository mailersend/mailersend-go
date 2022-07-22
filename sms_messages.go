package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const smsMessagesPath = "/sms-messages"

type SmsMessageService service

// smsListMessagesRoot - format of activity response
type smsListMessagesRoot struct {
	Data  []SmsMessageData `json:"data"`
	Links Links            `json:"links"`
	Meta  Meta             `json:"meta"`
}

// smsSingleMessagesRoot - format of activity response
type smsSingleMessagesRoot SmsMessageRoot

// ListSmsMessageOptions - modifies the behavior of SmsMessagesService.List method
type ListSmsMessageOptions struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func (s *SmsMessageService) List(ctx context.Context, options *ListSmsMessageOptions) (*smsListMessagesRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, smsMessagesPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(smsListMessagesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SmsMessageService) Get(ctx context.Context, smsMessageID string) (*smsSingleMessagesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsMessagesPath, smsMessageID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(smsSingleMessagesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
