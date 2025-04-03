package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const emailVerificationBasePath = "/email-verification"

type EmailVerificationService interface {
	List(ctx context.Context, options *ListEmailVerificationOptions) (*EmailVerificationRoot, *Response, error)
	Get(ctx context.Context, emailVerificationId string) (*SingleEmailVerificationRoot, *Response, error)
	Update(ctx context.Context, options *DomainSettingOptions) (*SingleEmailVerificationRoot, *Response, error)
	Delete(ctx context.Context, domainID string) (*Response, error)
	Create(ctx context.Context, options *CreateEmailVerificationOptions) (*SingleEmailVerificationRoot, *Response, error)
	Verify(ctx context.Context, emailVerificationId string) (*SingleEmailVerificationRoot, *Response, error)
	VerifySingle(ctx context.Context, options *SingleEmailVerificationOptions) (*ResultSingleEmailVerification, *Response, error)
	GetResults(ctx context.Context, options *GetEmailVerificationOptions) (*ResultEmailVerificationRoot, *Response, error)
}

type emailVerificationService struct {
	*service
}

// EmailVerificationRoot format of verification response
type EmailVerificationRoot struct {
	Data  []EmailVerification `json:"data"`
	Links Links               `json:"links"`
	Meta  Meta                `json:"meta"`
}

// singleDomainRoot format of single verification response
type SingleEmailVerificationRoot struct {
	Data EmailVerification `json:"data"`
}

type ResultEmailVerificationRoot struct {
	Data  []Result `json:"data"`
	Links Links    `json:"links"`
	Meta  Meta     `json:"meta"`
}

type ResultSingleEmailVerification struct {
	Status string `json:"status"`
}

type EmailVerification struct {
	Id                  string      `json:"id"`
	Name                string      `json:"name"`
	Total               int         `json:"total"`
	VerificationStarted interface{} `json:"verification_started"`
	VerificationEnded   interface{} `json:"verification_ended"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
	Status              Status      `json:"status"`
	Source              string      `json:"source"`
	Statistics          Statistics  `json:"statistics"`
}

type Status struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Statistics struct {
	Valid           int `json:"valid"`
	CatchAll        int `json:"catch_all"`
	MailboxFull     int `json:"mailbox_full"`
	RoleBased       int `json:"role_based"`
	Unknown         int `json:"unknown"`
	SyntaxError     int `json:"syntax_error"`
	Typo            int `json:"typo"`
	MailboxNotFound int `json:"mailbox_not_found"`
	Disposable      int `json:"disposable"`
	MailboxBlocked  int `json:"mailbox_blocked"`
	Failed          int `json:"failed"`
}

type Result struct {
	Address string `json:"address"`
	Result  string `json:"result"`
}

// ListEmailVerificationOptions - modifies the behavior of emailVerificationService.List Method
type ListEmailVerificationOptions struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// CreateEmailVerificationOptions -  modifies the behavior of emailVerificationService.Create Method
type CreateEmailVerificationOptions struct {
	Name   string   `json:"name"`
	Emails []string `json:"emails"`
}

// GetEmailVerificationOptions - modifies the behavior of emailVerificationService.List and emailVerificationService.GetResult Method
type GetEmailVerificationOptions struct {
	EmailVerificationId string `url:"-"`
	Page                int    `url:"page,omitempty"`
	Limit               int    `url:"limit,omitempty"`
}

type SingleEmailVerificationOptions struct {
	Email string `json:"email"`
}

func (s *emailVerificationService) List(ctx context.Context, options *ListEmailVerificationOptions) (*EmailVerificationRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, emailVerificationBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(EmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *emailVerificationService) Get(ctx context.Context, emailVerificationId string) (*SingleEmailVerificationRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", emailVerificationBasePath, emailVerificationId)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleEmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *emailVerificationService) Update(ctx context.Context, options *DomainSettingOptions) (*SingleEmailVerificationRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/settings", emailVerificationBasePath, options.DomainID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleEmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *emailVerificationService) Delete(ctx context.Context, domainID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", emailVerificationBasePath, domainID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

func (s *emailVerificationService) Create(ctx context.Context, options *CreateEmailVerificationOptions) (*SingleEmailVerificationRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, emailVerificationBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleEmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *emailVerificationService) Verify(ctx context.Context, emailVerificationId string) (*SingleEmailVerificationRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/verify", emailVerificationBasePath, emailVerificationId)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleEmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *emailVerificationService) VerifySingle(ctx context.Context, options *SingleEmailVerificationOptions) (*ResultSingleEmailVerification, *Response, error) {
	path := fmt.Sprintf("%s/verify", emailVerificationBasePath)

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	verification := new(ResultSingleEmailVerification)
	res, err := s.client.do(ctx, req, verification)
	if err != nil {
		return nil, res, err
	}

	return verification, res, nil
}

func (s *emailVerificationService) GetResults(ctx context.Context, options *GetEmailVerificationOptions) (*ResultEmailVerificationRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/results", emailVerificationBasePath, options.EmailVerificationId)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ResultEmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
