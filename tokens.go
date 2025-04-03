package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const tokenBasePath = "/token"

type TokenService interface {
	Create(ctx context.Context, options *CreateTokenOptions) (*TokenRoot, *Response, error)
	Update(ctx context.Context, options *UpdateTokenOptions) (*TokenRoot, *Response, error)
	Delete(ctx context.Context, tokenID string) (*Response, error)
}

type tokenService struct {
	*service
}

// TokenRoot - format of token response
type TokenRoot struct {
	Data Token `json:"data"`
}

type Token struct {
	ID          string `json:"id,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}

// CreateTokenOptions - modifies the behavior of TokenService.Create Method
type CreateTokenOptions struct {
	Name     string   `json:"name"`
	DomainID string   `json:"domain_id"`
	Scopes   []string `json:"scopes"`
}

// UpdateTokenOptions - modifies the behavior of TokenService.Update Method
type UpdateTokenOptions struct {
	TokenID string `json:"-"`
	Status  string `json:"status"`
}

func (s *tokenService) Create(ctx context.Context, options *CreateTokenOptions) (*TokenRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, tokenBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(TokenRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *tokenService) Update(ctx context.Context, options *UpdateTokenOptions) (*TokenRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/settings", tokenBasePath, options.TokenID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(TokenRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *tokenService) Delete(ctx context.Context, tokenID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", tokenBasePath, tokenID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
