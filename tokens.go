package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const tokenBasePath = "/token"

type TokenService service

// tokenRoot - format of token response
type tokenRoot struct {
	Data token `json:"data"`
}

type token struct {
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

func (s *TokenService) Create(ctx context.Context, options *CreateTokenOptions) (*tokenRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, tokenBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(tokenRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *TokenService) Update(ctx context.Context, options *UpdateTokenOptions) (*tokenRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/settings", tokenBasePath, options.TokenID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(tokenRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *TokenService) Delete(ctx context.Context, tokenID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", tokenBasePath, tokenID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
