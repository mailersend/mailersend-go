package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const identitiesBasePath = "/identities"

type IdentityService interface {
	List(ctx context.Context, options *ListIdentityOptions) (*IdentityRoot, *Response, error)
	Get(ctx context.Context, identityID string) (*SingleIdentityRoot, *Response, error)
	GetByEmail(ctx context.Context, identityEmail string) (*SingleIdentityRoot, *Response, error)
	Create(ctx context.Context, options *CreateIdentityOptions) (*SingleIdentityRoot, *Response, error)
	Update(ctx context.Context, identityID string, options *UpdateIdentityOptions) (*SingleIdentityRoot, *Response, error)
	UpdateByEmail(ctx context.Context, identityEmail string, options *UpdateIdentityOptions) (*SingleIdentityRoot, *Response, error)
	Delete(ctx context.Context, identityID string) (*Response, error)
	DeleteByEmail(ctx context.Context, identityEmail string) (*Response, error)
}

type identityService struct {
	*service
}

// IdentityRoot - format of identity response
type IdentityRoot struct {
	Data  []Identity `json:"data"`
	Links Links      `json:"links"`
	Meta  Meta       `json:"meta"`
}

// SingleIdentityRoot - format of inbound response
type SingleIdentityRoot struct {
	Data Identity `json:"data"`
}

type Identity struct {
	ID           string         `json:"id"`
	Email        string         `json:"email"`
	Name         string         `json:"name"`
	ReplyToEmail interface{}    `json:"reply_to_email"`
	ReplyToName  interface{}    `json:"reply_to_name"`
	IsVerified   bool           `json:"is_verified"`
	Resends      int            `json:"resends"`
	AddNote      bool           `json:"add_note"`
	PersonalNote interface{}    `json:"personal_note"`
	Domain       IdentityDomain `json:"domain"`
}

type IdentityDomain struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ListIdentityOptions - modifies the behavior of *IdentityService.List Method
type ListIdentityOptions struct {
	DomainID string `url:"domain_id"`
	Page     int    `url:"page,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

type CreateIdentityOptions struct {
	DomainID     string `json:"domain_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PersonalNote string `json:"personal_note"`
	ReplyToName  string `json:"reply_to_name"`
	ReplyToEmail string `json:"reply_to_email"`
	AddNote      bool   `json:"add_note"`
}

// UpdateIdentityOptions - the Options to set when creating an Identity resource
type UpdateIdentityOptions CreateIdentityOptions

func (s *identityService) List(ctx context.Context, options *ListIdentityOptions) (*IdentityRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, identitiesBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(IdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *identityService) Get(ctx context.Context, identityID string) (*SingleIdentityRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", identitiesBasePath, identityID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleIdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *identityService) GetByEmail(ctx context.Context, identityEmail string) (*SingleIdentityRoot, *Response, error) {
	path := fmt.Sprintf("%s/email/%s", identitiesBasePath, identityEmail)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleIdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *identityService) Create(ctx context.Context, options *CreateIdentityOptions) (*SingleIdentityRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, identitiesBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleIdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *identityService) Update(ctx context.Context, identityID string, options *UpdateIdentityOptions) (*SingleIdentityRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", identitiesBasePath, identityID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleIdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *identityService) UpdateByEmail(ctx context.Context, identityEmail string, options *UpdateIdentityOptions) (*SingleIdentityRoot, *Response, error) {
	path := fmt.Sprintf("%s/email/%s", identitiesBasePath, identityEmail)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleIdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *identityService) Delete(ctx context.Context, identityID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", identitiesBasePath, identityID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

func (s *identityService) DeleteByEmail(ctx context.Context, identityEmail string) (*Response, error) {
	path := fmt.Sprintf("%s/email/%s", identitiesBasePath, identityEmail)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
