package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const smsRecipientPath = "/sms-recipients"

type SmsRecipientService interface {
	List(ctx context.Context, options *SmsRecipientOptions) (*SmsRecipientRoot, *Response, error)
	Get(ctx context.Context, smsRecipientId string) (*SingleSmsRecipientRoot, *Response, error)
	Update(ctx context.Context, options *SmsRecipientSettingOptions) (*SingleSmsRecipientUpdateRoot, *Response, error)
}

type smsRecipientService struct {
	*service
}

// SmsRecipientRoot - format of activity response
type SmsRecipientRoot struct {
	Data  []SmsRecipient `json:"data"`
	Links Links          `json:"links"`
	Meta  Meta           `json:"meta"`
}

// singleSmsNumberRoot - format of activity response
type SingleSmsRecipientRoot struct {
	Data SmsRecipientData `json:"data"`
}

// SingleSmsRecipientUpdateRoot - format of activity response
type SingleSmsRecipientUpdateRoot struct {
	Data SmsRecipientDataUpdate `json:"data"`
}

type SmsRecipient struct {
	Id        string    `json:"id"`
	Number    string    `json:"number"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type SmsRecipientData struct {
	Id        string    `json:"id"`
	Number    string    `json:"number"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Sms       []SmsMessage
}

type SmsRecipientDataUpdate struct {
	Id        string    `json:"id"`
	Number    string    `json:"number"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// SmsRecipientSettingOptions - modifies the behavior of SmsNumbersService.Update method
type SmsRecipientSettingOptions struct {
	Id     string `json:"-"`
	Status string `json:"status,omitempty"`
}

// SmsRecipientOptions - modifies the behavior of SmsNumbersService.List method
type SmsRecipientOptions struct {
	Status      bool   `url:"status,omitempty"`
	SmsNumberId string `url:"sms_number_id,omitempty"`
	Page        int    `url:"page,omitempty"`
	Limit       int    `url:"limit,omitempty"`
}

func (s *smsRecipientService) List(ctx context.Context, options *SmsRecipientOptions) (*SmsRecipientRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, smsRecipientPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SmsRecipientRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *smsRecipientService) Get(ctx context.Context, smsRecipientId string) (*SingleSmsRecipientRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsRecipientPath, smsRecipientId)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleSmsRecipientRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *smsRecipientService) Update(ctx context.Context, options *SmsRecipientSettingOptions) (*SingleSmsRecipientUpdateRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsRecipientPath, options.Id)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleSmsRecipientUpdateRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
