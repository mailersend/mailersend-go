package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const inboundBasePath = "/inbound"

type InboundService service

// inboundRoot - format of webhook response
type inboundRoot struct {
	Data  []inbound `json:"data"`
	Links Links     `json:"links"`
	Meta  Meta      `json:"meta"`
}

// singleInboundRoot - format of inbound response
type singleInboundRoot struct {
	Data inbound `json:"data"`
}

type inbound struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Address      string      `json:"address"`
	Domain       string      `json:"domain"`
	DNSCheckedAt interface{} `json:"dns_checked_at"`
	Enabled      bool        `json:"enabled"`
	Filters      []filters   `json:"filters"`
	Forwards     []forwards  `json:"forwards"`
	MxValues     mxValues    `json:"mxValues"`
}

type filters struct {
	Type     string      `json:"type"`
	Key      interface{} `json:"key"`
	Comparer string      `json:"comparer"`
	Value    string      `json:"value"`
}

type forwards struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Value  string `json:"value"`
	Secret string `json:"secret"`
}

type mxValues struct {
	Priority int    `json:"priority"`
	Target   string `json:"target"`
}

// ListInboundOptions - modifies the behavior of *InboundService.List Method
type ListInboundOptions struct {
	DomainID string `url:"domain_id"`
	Page     int    `url:"page,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

// CreateInboundOptions - the Options to set when creating an inbound resource
type CreateInboundOptions struct {
	DomainID         string      `json:"domain_id"`
	Name             string      `json:"name"`
	DomainEnabled    bool        `json:"domain_enabled"`
	InboundDomain    string      `json:"inbound_domain"`
	InboundAddress   string      `json:"inbound_address"`
	InboundSubdomain string      `json:"inbound_subdomain"`
	MatchFilter      MatchFilter `json:"match_filter"`
	CatchFilter      CatchFilter `json:"catch_filter"`
	Forwards         []Forwards  `json:"forwards"`
}

type MatchFilter struct {
	Type string `json:"type"`
}

type Filters struct {
	Comparer string `json:"comparer"`
	Value    string `json:"value"`
}

type CatchFilter struct {
	Type    string    `json:"type"`
	Filters []Filters `json:"filters"`
}

type Forwards struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// UpdateInboundOptions - the Options to set when creating an inbound resource
type UpdateInboundOptions CreateInboundOptions

func (s *InboundService) List(ctx context.Context, options *ListInboundOptions) (*inboundRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, inboundBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(inboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *InboundService) Get(ctx context.Context, inboundID string) (*singleInboundRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", inboundBasePath, inboundID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleInboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *InboundService) Create(ctx context.Context, options *CreateWebhookOptions) (*singleInboundRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, inboundBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleInboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *InboundService) Update(ctx context.Context, options *UpdateInboundOptions) (*singleInboundRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", inboundBasePath, options.DomainID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleInboundRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *InboundService) Delete(ctx context.Context, inboundID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", inboundBasePath, inboundID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
