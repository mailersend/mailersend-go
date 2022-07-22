package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const smsInboundPath = "/sms-inbounds"

type SmsInboundService service

// smsInboundRoot - format of activity response
type smsInboundRoot struct {
	Data  []smsInbound `json:"data"`
	Links Links        `json:"links"`
	Meta  Meta         `json:"meta"`
}

// singleSmsInboundRoot - format of activity response
type singleSmsInboundRoot struct {
	Data smsInbound `json:"data"`
}

type smsInbound struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Filter     Filter
	ForwardUrl string    `json:"forward_url"`
	Enabled    bool      `json:"enabled"`
	Secret     string    `json:"secret"`
	CreatedAt  time.Time `json:"created_at"`
}

// CreateSmsInboundOptions - modifies the behavior of *WebhookService.Create Method
type CreateSmsInboundOptions struct {
	SmsNumberId string `json:"sms_number_id"`
	Name        string `json:"name"`
	ForwardUrl  string `json:"forward_url"`
	Filter      Filter `json:"filter"`
	Enabled     *bool  `json:"enabled"`
}

// UpdateSmsInboundOptions - modifies the behavior of SmsNumbersService.Update method
type UpdateSmsInboundOptions struct {
	Id          string `json:"-"`
	SmsNumberId string `json:"sms_number_id,omitempty"`
	Name        string `json:"name,omitempty"`
	ForwardUrl  string `json:"forward_url,omitempty"`
	Filter      Filter `json:"filter,omitempty"`
	Enabled     *bool  `json:"enabled,omitempty"`
}

// ListSmsInboundOptions - modifies the behavior of SmsNumbersService.List method
type ListSmsInboundOptions struct {
	SmsNumberId string `url:"sms_number_id,omitempty"`
	Enabled     *bool  `url:"enabled,omitempty"`
	Page        int    `url:"page,omitempty"`
	Limit       int    `url:"limit,omitempty"`
}

func (s *SmsInboundService) List(ctx context.Context, options *ListSmsInboundOptions) (*smsInboundRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, smsInboundPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(smsInboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SmsInboundService) Get(ctx context.Context, smsInboundId string) (*singleSmsInboundRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsInboundPath, smsInboundId)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleSmsInboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SmsInboundService) Create(ctx context.Context, options *CreateSmsInboundOptions) (*singleSmsInboundRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, smsInboundPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleSmsInboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SmsInboundService) Update(ctx context.Context, options *UpdateSmsInboundOptions) (*singleSmsInboundRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsInboundPath, options.Id)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleSmsInboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SmsInboundService) Delete(ctx context.Context, smsInboundId string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", smsInboundPath, smsInboundId)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
