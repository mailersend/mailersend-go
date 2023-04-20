package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const templateBasePath = "/templates"

type TemplateService service

// templateRoot format of template response
type templateRoot struct {
	Data  []template `json:"data"`
	Links Links      `json:"links"`
	Meta  Meta       `json:"meta"`
}

type template struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	ImagePath string `json:"image_path"`
	CreatedAt string `json:"created_at"`
}

type singleTemplateRoot struct {
	Data singleTemplate `json:"data"`
}

type singleTemplate struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Type          string        `json:"type"`
	ImagePath     string        `json:"image_path"`
	CreatedAt     time.Time     `json:"created_at"`
	Category      interface{}   `json:"category"`
	Domain        Domain        `json:"domain"`
	TemplateStats templateStats `json:"template_stats"`
}

type templateStats struct {
	Total           int       `json:"total"`
	Queued          int       `json:"queued"`
	Sent            int       `json:"sent"`
	Rejected        int       `json:"rejected"`
	Delivered       int       `json:"delivered"`
	LastEmailSentAt time.Time `json:"last_email_sent_at"`
}

// ListTemplateOptions - modifies the behavior of TemplateService.List Method
type ListTemplateOptions struct {
	DomainID string `url:"domain_id,omitempty"`
	Page     int    `url:"page,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

func (s *TemplateService) List(ctx context.Context, options *ListTemplateOptions) (*templateRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, templateBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(templateRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *TemplateService) Get(ctx context.Context, templateID string) (*singleTemplateRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", templateBasePath, templateID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleTemplateRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *TemplateService) Delete(ctx context.Context, templateID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", templateBasePath, templateID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
