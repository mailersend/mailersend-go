package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const webhookBasePath = "/webhooks"

type WebhookService interface {
	List(ctx context.Context, options *ListWebhookOptions) (*WebhookRoot, *Response, error)
	Get(ctx context.Context, webhookID string) (*SingleWebhookRoot, *Response, error)
	Create(ctx context.Context, options *CreateWebhookOptions) (*SingleWebhookRoot, *Response, error)
	Update(ctx context.Context, options *UpdateWebhookOptions) (*SingleWebhookRoot, *Response, error)
	Delete(ctx context.Context, webhookID string) (*Response, error)
}

type webhookService struct {
	*service
}

// WebhookRoot - format of webhook response
type WebhookRoot struct {
	Data  []Webhook `json:"data"`
	Links Links     `json:"links"`
	Meta  Meta      `json:"meta"`
}

// SingleWebhookRoot - format of webhook response
type SingleWebhookRoot struct {
	Data Webhook `json:"data"`
}

type Webhook struct {
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

func (s *webhookService) List(ctx context.Context, options *ListWebhookOptions) (*WebhookRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, webhookBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(WebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *webhookService) Get(ctx context.Context, webhookID string) (*SingleWebhookRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", webhookBasePath, webhookID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleWebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *webhookService) Create(ctx context.Context, options *CreateWebhookOptions) (*SingleWebhookRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, webhookBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleWebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *webhookService) Update(ctx context.Context, options *UpdateWebhookOptions) (*SingleWebhookRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", webhookBasePath, options.WebhookID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleWebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *webhookService) Delete(ctx context.Context, webhookID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", webhookBasePath, webhookID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
