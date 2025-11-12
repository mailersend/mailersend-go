package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const usersBasePath = "/users"

// UserService defines the interface for user management operations
type UserService interface {
	List(ctx context.Context, options *ListUserOptions) (*UserRoot, *Response, error)
	Get(ctx context.Context, userID string) (*SingleUserRoot, *Response, error)
	Invite(ctx context.Context, options *InviteUserOptions) (*SingleUserRoot, *Response, error)
	Update(ctx context.Context, userID string, options *UpdateUserOptions) (*SingleUserRoot, *Response, error)
	Delete(ctx context.Context, userID string) (*Response, error)
}

type userService struct {
	*service
}

// UserRoot - format of user response
type UserRoot struct {
	Data  []User `json:"data"`
	Links Links  `json:"links"`
	Meta  Meta   `json:"meta"`
}

// SingleUserRoot - format of user response
type SingleUserRoot struct {
	Data User `json:"data"`
}

// User represents a MailerSend account user
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role"`
	Status    string `json:"status,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ListUserOptions defines options for listing users
type ListUserOptions struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// InviteUserOptions defines options for inviting a user
type InviteUserOptions struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

// UpdateUserOptions defines options for updating a user
type UpdateUserOptions struct {
	Role string `json:"role,omitempty"`
}

// List retrieves a list of account users
func (s *userService) List(ctx context.Context, options *ListUserOptions) (*UserRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, usersBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	users := new(UserRoot)
	res, err := s.client.do(ctx, req, users)
	if err != nil {
		return nil, res, err
	}

	return users, res, nil
}

// Get retrieves a single user by ID
func (s *userService) Get(ctx context.Context, userID string) (*SingleUserRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", usersBasePath, userID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(SingleUserRoot)
	res, err := s.client.do(ctx, req, user)
	if err != nil {
		return nil, res, err
	}

	return user, res, nil
}

// Invite invites a new user to the account
func (s *userService) Invite(ctx context.Context, options *InviteUserOptions) (*SingleUserRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, usersBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	user := new(SingleUserRoot)
	res, err := s.client.do(ctx, req, user)
	if err != nil {
		return nil, res, err
	}

	return user, res, nil
}

// Update updates an existing user
func (s *userService) Update(ctx context.Context, userID string, options *UpdateUserOptions) (*SingleUserRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", usersBasePath, userID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	user := new(SingleUserRoot)
	res, err := s.client.do(ctx, req, user)
	if err != nil {
		return nil, res, err
	}

	return user, res, nil
}

// Delete deletes a user from the account
func (s *userService) Delete(ctx context.Context, userID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", usersBasePath, userID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
