package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const identitiesBasePath = "/identities"

type IdentityService service

// identityRoot - format of identity response
type identityRoot struct {
	Data  []identity `json:"data"`
	Links Links      `json:"links"`
	Meta  Meta       `json:"meta"`
}

// singleIdentityRoot - format of inbound response
type singleIdentityRoot struct {
	Data identity `json:"data"`
}

type identity struct {
	ID           string         `json:"id"`
	Email        string         `json:"email"`
	Name         string         `json:"name"`
	ReplyToEmail interface{}    `json:"reply_to_email"`
	ReplyToName  interface{}    `json:"reply_to_name"`
	IsVerified   bool           `json:"is_verified"`
	Resends      int            `json:"resends"`
	AddNote      bool           `json:"add_note"`
	PersonalNote interface{}    `json:"personal_note"`
	Domain       identityDomain `json:"domain"`
}

type identityDomain struct {
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

func (s *IdentityService) List(ctx context.Context, options *ListIdentityOptions) (*identityRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, identitiesBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(identityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *IdentityService) Get(ctx context.Context, identityID string) (*singleIdentityRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", identitiesBasePath, identityID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleIdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *IdentityService) GetByEmail(ctx context.Context, identityEmail string) (*singleIdentityRoot, *Response, error) {
	path := fmt.Sprintf("%s/email/%s", identitiesBasePath, identityEmail)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleIdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *IdentityService) Create(ctx context.Context, options *CreateIdentityOptions) (*singleIdentityRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, identitiesBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleIdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *IdentityService) Update(ctx context.Context, identityID string, options *UpdateIdentityOptions) (*singleIdentityRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", identitiesBasePath, identityID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleIdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *IdentityService) UpdateByEmail(ctx context.Context, identityEmail string, options *UpdateIdentityOptions) (*singleIdentityRoot, *Response, error) {
	path := fmt.Sprintf("%s/email/%s", identitiesBasePath, identityEmail)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleIdentityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *IdentityService) Delete(ctx context.Context, identityID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", identitiesBasePath, identityID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

func (s *IdentityService) DeleteByEmail(ctx context.Context, identityEmail string) (*Response, error) {
	path := fmt.Sprintf("%s/email/%s", identitiesBasePath, identityEmail)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
