package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const webhookBasePath = "/webhooks"

type WebhookService service

// webhookRoot - format of webhook response
type webhookRoot struct {
	Data  []webhook `json:"data"`
	Links Links     `json:"links"`
	Meta  Meta      `json:"meta"`
}

// singleWebhookRoot - format of webhook response
type singleWebhookRoot struct {
	Data webhook `json:"data"`
}

type webhook struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Events    []string  `json:"events"`
	Name      string    `json:"name"`
	Enabled   bool      `json:"enabled"`
	Editable  bool      `json:"editable"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Domain    Domain    `json:"domain"`
}

// ListWebhookOptions - modifies the behavior of *WebhookService.List Method
type ListWebhookOptions struct {
	DomainID string `url:"domain_id"`
	Limit    int    `url:"limit,omitempty"`
}

// CreateWebhookOptions - modifies the behavior of *WebhookService.Create Method
type CreateWebhookOptions struct {
	Name     string   `json:"name"`
	DomainID string   `json:"domain_id"`
	URL      string   `json:"url"`
	Enabled  *bool    `json:"enabled,omitempty"`
	Events   []string `json:"events"`
}

// UpdateWebhookOptions - modifies the behavior of *WebhookService.Update Method
type UpdateWebhookOptions struct {
	WebhookID string   `json:"-"`
	Name      string   `json:"name,omitempty"`
	URL       string   `json:"url,omitempty"`
	Enabled   *bool    `json:"enabled,omitempty"`
	Events    []string `json:"events,omitempty"`
}

func (s *WebhookService) List(ctx context.Context, options *ListWebhookOptions) (*webhookRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, webhookBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(webhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *WebhookService) Get(ctx context.Context, webhookID string) (*singleWebhookRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", webhookBasePath, webhookID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleWebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *WebhookService) Create(ctx context.Context, options *CreateWebhookOptions) (*singleWebhookRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, webhookBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleWebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *WebhookService) Update(ctx context.Context, options *UpdateWebhookOptions) (*singleWebhookRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", webhookBasePath, options.WebhookID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleWebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *WebhookService) Delete(ctx context.Context, webhookID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", webhookBasePath, webhookID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
