package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const whatsappTemplatesPath = "/whatsapp/templates"

type WhatsAppTemplateService interface {
	List(ctx context.Context, options *ListWhatsAppTemplateOptions) (*WhatsAppListTemplatesRoot, *Response, error)
	Get(ctx context.Context, whatsAppTemplateID string) (*WhatsAppSingleTemplateRoot, *Response, error)
}

type whatsAppTemplateService struct {
	*service
}

// WhatsAppListTemplatesRoot - format of WhatsApp templates list response
type WhatsAppListTemplatesRoot struct {
	Data  []WhatsAppTemplateData `json:"data"`
	Links Links                  `json:"links"`
	Meta  Meta                   `json:"meta"`
}

// WhatsAppSingleTemplateRoot - format of single WhatsApp template response
type WhatsAppSingleTemplateRoot struct {
	Data WhatsAppTemplateData `json:"data"`
}

// WhatsAppTemplateData - a WhatsApp template. ID is the template_id to send it with
type WhatsAppTemplateData struct {
	ID                string                     `json:"id"`
	WhatsAppAccountID string                     `json:"whatsapp_account_id"`
	Name              string                     `json:"name"`
	ReadableName      string                     `json:"readable_name"`
	Category          string                     `json:"category"`
	Language          string                     `json:"language"`
	Status            string                     `json:"status"`
	Components        WhatsAppTemplateComponents `json:"components"`
	CreatedAt         time.Time                  `json:"created_at"`
	UpdatedAt         time.Time                  `json:"updated_at"`
}

// WhatsAppTemplateComponents - the content of a template. Header and Body are nil when the template has none
type WhatsAppTemplateComponents struct {
	Header  *WhatsAppTemplateTextComponent `json:"header"`
	Body    *WhatsAppTemplateTextComponent `json:"body"`
	Buttons []WhatsAppTemplateButton       `json:"buttons"`
}

// WhatsAppTemplateTextComponent - Text includes any variable placeholders, and is nil for a component without text such as a media header
type WhatsAppTemplateTextComponent struct {
	Text *string `json:"text"`
}

// WhatsAppTemplateButton - URL is only set for a URL button
type WhatsAppTemplateButton struct {
	Type string  `json:"type"`
	Text string  `json:"text"`
	URL  *string `json:"url"`
}

// ListWhatsAppTemplateOptions - modifies the behavior of WhatsAppTemplateService.List method
type ListWhatsAppTemplateOptions struct {
	WhatsAppAccountID string `url:"whatsapp_account_id,omitempty"`
	Status            string `url:"status,omitempty"`
	Category          string `url:"category,omitempty"`
	Language          string `url:"language,omitempty"`
	Page              int    `url:"page,omitempty"`
	Limit             int    `url:"limit,omitempty"`
}

func (s *whatsAppTemplateService) List(ctx context.Context, options *ListWhatsAppTemplateOptions) (*WhatsAppListTemplatesRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, whatsappTemplatesPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppListTemplatesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *whatsAppTemplateService) Get(ctx context.Context, whatsAppTemplateID string) (*WhatsAppSingleTemplateRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", whatsappTemplatesPath, whatsAppTemplateID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppSingleTemplateRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
