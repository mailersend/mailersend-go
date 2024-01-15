package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const emailVerificationBasePath = "/email-verification"

type EmailVerificationService service

// emailVerificationRoot format of verification response
type emailVerificationRoot struct {
	Data  []emailVerification `json:"data"`
	Links Links               `json:"links"`
	Meta  Meta                `json:"meta"`
}

// singleDomainRoot format of single verification response
type singleEmailVerificationRoot struct {
	Data emailVerification `json:"data"`
}

type resultEmailVerificationRoot struct {
	Data  []result `json:"data"`
	Links Links    `json:"links"`
	Meta  Meta     `json:"meta"`
}

type resultSingleEmailVerification struct {
	Status string `json:"status"`
}

type emailVerification struct {
	Id                  string      `json:"id"`
	Name                string      `json:"name"`
	Total               int         `json:"total"`
	VerificationStarted interface{} `json:"verification_started"`
	VerificationEnded   interface{} `json:"verification_ended"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
	Status              status      `json:"status"`
	Source              string      `json:"source"`
	Statistics          statistics  `json:"statistics"`
}

type status struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type statistics struct {
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

type result struct {
	Address string `json:"address"`
	Result  string `json:"result"`
}

// ListEmailVerificationOptions - modifies the behavior of EmailVerificationService.List Method
type ListEmailVerificationOptions struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// CreateEmailVerificationOptions -  modifies the behavior of EmailVerificationService.Create Method
type CreateEmailVerificationOptions struct {
	Name   string   `json:"name"`
	Emails []string `json:"emails"`
}

// GetEmailVerificationOptions - modifies the behavior of EmailVerificationService.List and EmailVerificationService.GetResult Method
type GetEmailVerificationOptions struct {
	EmailVerificationId string `url:"-"`
	Page                int    `url:"page,omitempty"`
	Limit               int    `url:"limit,omitempty"`
}

type SingleEmailVerificationOptions struct {
	Email string `json:"email"`
}

func (s *EmailVerificationService) List(ctx context.Context, options *ListEmailVerificationOptions) (*emailVerificationRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, emailVerificationBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(emailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *EmailVerificationService) Get(ctx context.Context, emailVerificationId string) (*singleEmailVerificationRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", emailVerificationBasePath, emailVerificationId)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleEmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *EmailVerificationService) Update(ctx context.Context, options *DomainSettingOptions) (*singleEmailVerificationRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/settings", emailVerificationBasePath, options.DomainID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleEmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *EmailVerificationService) Delete(ctx context.Context, domainID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", emailVerificationBasePath, domainID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

func (s *EmailVerificationService) Create(ctx context.Context, options *CreateEmailVerificationOptions) (*singleEmailVerificationRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, emailVerificationBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleEmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *EmailVerificationService) Verify(ctx context.Context, emailVerificationId string) (*singleEmailVerificationRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/verify", emailVerificationBasePath, emailVerificationId)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleEmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *EmailVerificationService) VerifySingle(ctx context.Context, options *SingleEmailVerificationOptions) (*resultSingleEmailVerification, *Response, error) {
	path := fmt.Sprintf("%s/verify", emailVerificationBasePath)

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	verification := new(resultSingleEmailVerification)
	res, err := s.client.do(ctx, req, verification)
	if err != nil {
		return nil, res, err
	}

	return verification, res, nil
}

func (s *EmailVerificationService) GetResults(ctx context.Context, options *GetEmailVerificationOptions) (*resultEmailVerificationRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/results", emailVerificationBasePath, options.EmailVerificationId)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(resultEmailVerificationRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
