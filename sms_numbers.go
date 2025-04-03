package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const smsNumbersPath = "/sms-numbers"

type SmsNumberService interface {
	List(ctx context.Context, options *SmsNumberOptions) (*SmsNumberRoot, *Response, error)
	Get(ctx context.Context, numberID string) (*SingleSmsNumberRoot, *Response, error)
	Update(ctx context.Context, options *SmsNumberSettingOptions) (*SingleSmsNumberRoot, *Response, error)
	Delete(ctx context.Context, numberID string) (*Response, error)
}

type smsNumberService struct {
	*service
}

// SmsNumberRoot - format of activity response
type SmsNumberRoot struct {
	Data  []Number `json:"data"`
	Links Links    `json:"links"`
	Meta  Meta     `json:"meta"`
}

// SingleSmsNumberRoot - format of activity response
type SingleSmsNumberRoot struct {
	Data Number `json:"data"`
}

type Number struct {
	Id              string    `json:"id"`
	TelephoneNumber string    `json:"telephone_number"`
	Paused          bool      `json:"paused"`
	CreatedAt       time.Time `json:"created_at"`
}

// SmsNumberSettingOptions - modifies the behavior of SmsNumbersService.Update method
type SmsNumberSettingOptions struct {
	Id     string `json:"-"`
	Paused *bool  `json:"paused,omitempty"`
}

// SmsNumberOptions - modifies the behavior of SmsNumbersService.List method
type SmsNumberOptions struct {
	Paused bool `url:"paused,omitempty"`
	Page   int  `url:"page,omitempty"`
	Limit  int  `url:"limit,omitempty"`
}

func (s *smsNumberService) List(ctx context.Context, options *SmsNumberOptions) (*SmsNumberRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, smsNumbersPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SmsNumberRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *smsNumberService) Get(ctx context.Context, numberID string) (*SingleSmsNumberRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsNumbersPath, numberID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleSmsNumberRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *smsNumberService) Update(ctx context.Context, options *SmsNumberSettingOptions) (*SingleSmsNumberRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsNumbersPath, options.Id)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleSmsNumberRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *smsNumberService) Delete(ctx context.Context, numberID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", smsNumbersPath, numberID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
