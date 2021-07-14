package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const messageBasePath = "/messages"

type MessageService service

// messageRoot format of message response
type messageRoot struct {
	Data  []message `json:"data"`
	Links Links     `json:"links"`
	Meta  Meta      `json:"meta"`
}

type message struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type singleMessageRoot struct {
	Data singleMessage `json:"data"`
}

type singleMessage struct {
	ID        string    `json:"id"`
	Emails    []email   `json:"emails"`
	Domain    Domain    `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type email struct {
	ID        string    `json:"id"`
	From      From      `json:"from"`
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

func (s *MessageService) List(ctx context.Context, options *ListMessageOptions) (*messageRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, messageBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(messageRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *MessageService) Get(ctx context.Context, messageID string) (*singleMessageRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", messageBasePath, messageID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleMessageRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
