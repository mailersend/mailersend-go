package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const recipientBasePath = "/recipients"

type RecipientService interface {
	List(ctx context.Context, options *ListRecipientOptions) (*RecipientRoot, *Response, error)
	Get(ctx context.Context, recipientID string) (*SingleRecipientRoot, *Response, error)
	Delete(ctx context.Context, recipientID string) (*Response, error)
}

type recipientService struct {
	*service
}

// RecipientRoot - recipients response
type RecipientRoot struct {
	Data  []RecipientObject `json:"data"`
	Links Links             `json:"links"`
	Meta  Meta              `json:"meta"`
}

// SingleRecipientRoot - single recipient response
type SingleRecipientRoot struct {
	Data RecipientData `json:"data"`
}

type RecipientData struct {
	ID        string        `json:"id"`
	Email     string        `json:"email"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
	DeletedAt string        `json:"deleted_at"`
	Emails    []interface{} `json:"emails"`
	Domain    Domain        `json:"domain"`
}

// RecipientObject - a single RecipientObject
type RecipientObject struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// ListRecipientOptions - modifies the behavior of RecipientService.List method
type ListRecipientOptions struct {
	DomainID string `url:"domain_id,omitempty"`
	Page     int    `url:"page,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

func (s *recipientService) List(ctx context.Context, options *ListRecipientOptions) (*RecipientRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, recipientBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RecipientRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *recipientService) Get(ctx context.Context, recipientID string) (*SingleRecipientRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", recipientBasePath, recipientID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleRecipientRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *recipientService) Delete(ctx context.Context, recipientID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", recipientBasePath, recipientID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)

}
