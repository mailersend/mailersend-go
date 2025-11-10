package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const smtpUsersBasePath = "/smtp-users"

// SmtpUserService defines the interface for SMTP user operations
type SmtpUserService interface {
	List(ctx context.Context, options *ListSmtpUserOptions) (*SmtpUserRoot, *Response, error)
	Get(ctx context.Context, smtpUserID string) (*SingleSmtpUserRoot, *Response, error)
	Create(ctx context.Context, options *CreateSmtpUserOptions) (*SingleSmtpUserRoot, *Response, error)
	Update(ctx context.Context, smtpUserID string, options *UpdateSmtpUserOptions) (*SingleSmtpUserRoot, *Response, error)
	Delete(ctx context.Context, smtpUserID string) (*Response, error)
}

type smtpUserService struct {
	*service
}

// SmtpUserRoot - format of smtp user response
type SmtpUserRoot struct {
	Data  []SmtpUser `json:"data"`
	Links Links      `json:"links"`
	Meta  Meta       `json:"meta"`
}

// SingleSmtpUserRoot - format of smtp user response
type SingleSmtpUserRoot struct {
	Data SmtpUser `json:"data"`
}

// SmtpUser represents an SMTP user
type SmtpUser struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Enabled   bool   `json:"enabled"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ListSmtpUserOptions defines options for listing SMTP users
type ListSmtpUserOptions struct {
	DomainID string `url:"domain_id,omitempty"`
	Page     int    `url:"page,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

// CreateSmtpUserOptions defines options for creating an SMTP user
type CreateSmtpUserOptions struct {
	DomainID string `json:"domain_id"`
	Name     string `json:"name"`
	Enabled  *bool  `json:"enabled,omitempty"`
}

// UpdateSmtpUserOptions defines options for updating an SMTP user
type UpdateSmtpUserOptions struct {
	Name    string `json:"name,omitempty"`
	Enabled *bool  `json:"enabled,omitempty"`
}

// List retrieves a list of SMTP users
func (s *smtpUserService) List(ctx context.Context, options *ListSmtpUserOptions) (*SmtpUserRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, smtpUsersBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	smtpUsers := new(SmtpUserRoot)
	res, err := s.client.do(ctx, req, smtpUsers)
	if err != nil {
		return nil, res, err
	}

	return smtpUsers, res, nil
}

// Get retrieves a single SMTP user by ID
func (s *smtpUserService) Get(ctx context.Context, smtpUserID string) (*SingleSmtpUserRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smtpUsersBasePath, smtpUserID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	smtpUser := new(SingleSmtpUserRoot)
	res, err := s.client.do(ctx, req, smtpUser)
	if err != nil {
		return nil, res, err
	}

	return smtpUser, res, nil
}

// Create creates a new SMTP user
func (s *smtpUserService) Create(ctx context.Context, options *CreateSmtpUserOptions) (*SingleSmtpUserRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, smtpUsersBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	smtpUser := new(SingleSmtpUserRoot)
	res, err := s.client.do(ctx, req, smtpUser)
	if err != nil {
		return nil, res, err
	}

	return smtpUser, res, nil
}

// Update updates an existing SMTP user
func (s *smtpUserService) Update(ctx context.Context, smtpUserID string, options *UpdateSmtpUserOptions) (*SingleSmtpUserRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smtpUsersBasePath, smtpUserID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	smtpUser := new(SingleSmtpUserRoot)
	res, err := s.client.do(ctx, req, smtpUser)
	if err != nil {
		return nil, res, err
	}

	return smtpUser, res, nil
}

// Delete deletes an SMTP user
func (s *smtpUserService) Delete(ctx context.Context, smtpUserID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", smtpUsersBasePath, smtpUserID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
