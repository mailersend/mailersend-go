package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const messageScheduleBasePath = "/message-schedules"

type ScheduleMessageService interface {
	List(ctx context.Context, options *ListScheduleMessageOptions) (*ScheduleMessageRoot, *Response, error)
	Get(ctx context.Context, messageID string) (*ScheduleMessageSingleRoot, *Response, error)
	Delete(ctx context.Context, messageID string) (*Response, error)
}

type scheduleMessageService struct {
	*service
}

type ScheduleMessageRoot struct {
	Data  []ScheduleMessageData `json:"data"`
	Links Links                 `json:"links"`
	Meta  Meta                  `json:"meta"`
}

type ScheduleMessageData struct {
	MessageID     string      `json:"message_id"`
	Subject       string      `json:"subject"`
	SendAt        time.Time   `json:"send_at"`
	Status        string      `json:"status"`
	StatusMessage interface{} `json:"status_message"`
	CreatedAt     string      `json:"created_at"`
}

type ScheduleMessageSingleRoot struct {
	Data ScheduleMessageSingleData `json:"data"`
}

type ScheduleDomain struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ScheduleMessage struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ScheduleMessageSingleData struct {
	MessageID     string          `json:"message_id"`
	Subject       string          `json:"subject"`
	SendAt        time.Time       `json:"send_at"`
	Status        string          `json:"status"`
	StatusMessage interface{}     `json:"status_message"`
	CreatedAt     time.Time       `json:"created_at"`
	Domain        ScheduleDomain  `json:"domain"`
	Message       ScheduleMessage `json:"message"`
}

// ListScheduleMessageOptions - modifies the behavior of MessageService.List Method
type ListScheduleMessageOptions struct {
	DomainID string `url:"domain_id,omitempty"`
	Status   string `url:"status,omitempty"`
	Page     int    `url:"page,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

func (s *scheduleMessageService) List(ctx context.Context, options *ListScheduleMessageOptions) (*ScheduleMessageRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, messageScheduleBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(ScheduleMessageRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *scheduleMessageService) Get(ctx context.Context, messageID string) (*ScheduleMessageSingleRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", messageScheduleBasePath, messageID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ScheduleMessageSingleRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *scheduleMessageService) Delete(ctx context.Context, messageID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", messageScheduleBasePath, messageID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
