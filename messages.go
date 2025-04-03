package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const messageBasePath = "/messages"

type MessageService interface {
	List(ctx context.Context, options *ListMessageOptions) (*MessageRoot, *Response, error)
	Get(ctx context.Context, messageID string) (*SingleMessageRoot, *Response, error)
}

type messageService struct {
	*service
}

// MessageRoot format of message response
type MessageRoot struct {
	Data  []MessageData `json:"data"`
	Links Links         `json:"links"`
	Meta  Meta          `json:"meta"`
}

type MessageData struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SingleMessageRoot struct {
	Data SingleMessage `json:"data"`
}

type SingleMessage struct {
	ID        string    `json:"id"`
	Emails    []Email   `json:"emails"`
	Domain    Domain    `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Email struct {
	ID        string    `json:"id"`
	From      string    `json:"from"`
	Subject   string    `json:"subject,omitempty"`
	Text      string    `json:"text,omitempty"`
	HTML      string    `json:"html,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	Status    string    `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ListMessageOptions - modifies the behavior of MessageService.List Method
type ListMessageOptions struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func (s *messageService) List(ctx context.Context, options *ListMessageOptions) (*MessageRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, messageBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(MessageRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *messageService) Get(ctx context.Context, messageID string) (*SingleMessageRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", messageBasePath, messageID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleMessageRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
