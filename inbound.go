package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const inboundBasePath = "/inbound"

type InboundService interface {
	List(ctx context.Context, options *ListInboundOptions) (*InboundRoot, *Response, error)
	Get(ctx context.Context, inboundID string) (*SingleInboundRoot, *Response, error)
	Create(ctx context.Context, options *CreateInboundOptions) (*SingleInboundRoot, *Response, error)
	Update(ctx context.Context, inboundID string, options *UpdateInboundOptions) (*SingleInboundRoot, *Response, error)
	Delete(ctx context.Context, inboundID string) (*Response, error)
}

type inboundService struct {
	*service
}

// InboundRoot - format of webhook response
type InboundRoot struct {
	Data  []Inbound `json:"data"`
	Links Links     `json:"links"`
	Meta  Meta      `json:"meta"`
}

// SingleInboundRoot - format of Inbound response
type SingleInboundRoot struct {
	Data Inbound `json:"data"`
}

type Inbound struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Address      string      `json:"address"`
	Domain       string      `json:"domain"`
	DNSCheckedAt interface{} `json:"dns_checked_at"`
	Priority     int         `json:"priority"`
	Enabled      bool        `json:"enabled"`
	Filters      []Filters   `json:"filters"`
	Forwards     []Forwards  `json:"forwards"`
	MxValues     mxValues    `json:"mxValues"`
}

type Filters struct {
	Type     string      `json:"type"`
	Key      interface{} `json:"key"`
	Comparer string      `json:"comparer"`
	Value    string      `json:"value"`
}

type Forwards struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Value  string `json:"value"`
	Secret string `json:"secret"`
}

type mxValues struct {
	Priority string `json:"priority"`
	Target   string `json:"target"`
}

// ListInboundOptions - modifies the behavior of *inboundService.List Method
type ListInboundOptions struct {
	DomainID string `url:"domain_id"`
	Page     int    `url:"page,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

// CreateInboundOptions - the Options to set when creating an inbound resource
type CreateInboundOptions struct {
	DomainID         string           `json:"domain_id"`
	Name             string           `json:"name"`
	DomainEnabled    bool             `json:"domain_enabled"`
	InboundDomain    string           `json:"inbound_domain,omitempty"`
	InboundAddress   string           `json:"inbound_address,omitempty"`
	InboundSubdomain string           `json:"inbound_subdomain,omitempty"`
	InboundPriority  int              `json:"inbound_priority,omitempty"`
	MatchFilter      *MatchFilter     `json:"match_filter,omitempty"`
	CatchFilter      *CatchFilter     `json:"catch_filter,omitempty"`
	Forwards         []ForwardsFilter `json:"forwards"`
}

type MatchFilter struct {
	Type string `json:"type,omitempty"`
}

type CatchFilter struct {
	Type    string   `json:"type,omitempty"`
	Filters []Filter `json:"filters,omitempty"`
}

type ForwardsFilter struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// UpdateInboundOptions - the Options to set when creating an inbound resource
type UpdateInboundOptions CreateInboundOptions

func (s *inboundService) List(ctx context.Context, options *ListInboundOptions) (*InboundRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, inboundBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(InboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *inboundService) Get(ctx context.Context, inboundID string) (*SingleInboundRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", inboundBasePath, inboundID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleInboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *inboundService) Create(ctx context.Context, options *CreateInboundOptions) (*SingleInboundRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, inboundBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleInboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *inboundService) Update(ctx context.Context, inboundID string, options *UpdateInboundOptions) (*SingleInboundRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", inboundBasePath, inboundID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleInboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *inboundService) Delete(ctx context.Context, inboundID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", inboundBasePath, inboundID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
