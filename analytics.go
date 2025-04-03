package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const analyticsBasePath = "/analytics"

type AnalyticsService interface {
	GetActivityByDate(ctx context.Context, options *AnalyticsOptions) (*AnalyticsActivityRoot, *Response, error)
	GetOpensByCountry(ctx context.Context, options *AnalyticsOptions) (*OpensRoot, *Response, error)
	GetOpensByUserAgent(ctx context.Context, options *AnalyticsOptions) (*OpensRoot, *Response, error)
	GetOpensByReadingEnvironment(ctx context.Context, options *AnalyticsOptions) (*OpensRoot, *Response, error)
}

type analyticsService struct {
	*service
}

// AnalyticsActivityRoot - format of analytics response
type AnalyticsActivityRoot struct {
	Data AnalyticsData `json:"data"`
}

type AnalyticsData struct {
	DateFrom string           `json:"date_from"`
	DateTo   string           `json:"date_to"`
	GroupBy  string           `json:"group_by"`
	Stats    []AnalyticsStats `json:"stats"`
}

type AnalyticsStats struct {
	Date           string `json:"date"`
	Queued         int    `json:"queued,omitempty"`
	Sent           int    `json:"sent,omitempty"`
	Delivered      int    `json:"delivered,omitempty"`
	SoftBounced    int    `json:"soft_bounced,omitempty"`
	HardBounced    int    `json:"hard_bounced,omitempty"`
	Junk           int    `json:"junk,omitempty"`
	Opened         int    `json:"opened,omitempty"`
	Clicked        int    `json:"clicked,omitempty"`
	Unsubscribed   int    `json:"unsubscribed,omitempty"`
	SpamComplaints int    `json:"spam_complaints,omitempty"`
}

type OpensRoot struct {
	Data OpenData `json:"data"`
}

type OpenData struct {
	DateFrom int         `json:"date_from"`
	DateTo   int         `json:"date_to"`
	Stats    []OpenStats `json:"stats"`
}

type OpenStats struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// AnalyticsOptions - modifies the behavior of AnalyticsService methods
type AnalyticsOptions struct {
	DomainID    string   `url:"domain_id,omitempty"`
	RecipientID []int64  `url:"recipient_id,omitempty"`
	DateFrom    int64    `url:"date_from"`
	DateTo      int64    `url:"date_to"`
	GroupBy     string   `url:"group_by,omitempty"`
	Tags        []string `url:"tags[],omitempty"`
	Event       []string `url:"event[],omitempty"`
}

func (s *analyticsService) GetActivityByDate(ctx context.Context, options *AnalyticsOptions) (*AnalyticsActivityRoot, *Response, error) {
	path := fmt.Sprintf("%s/date", analyticsBasePath)

	req, err := s.client.newRequest("GET", path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(AnalyticsActivityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *analyticsService) GetOpensByCountry(ctx context.Context, options *AnalyticsOptions) (*OpensRoot, *Response, error) {
	path := fmt.Sprintf("%s/country", analyticsBasePath)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(OpensRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *analyticsService) GetOpensByUserAgent(ctx context.Context, options *AnalyticsOptions) (*OpensRoot, *Response, error) {
	path := fmt.Sprintf("%s/ua-name", analyticsBasePath)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(OpensRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *analyticsService) GetOpensByReadingEnvironment(ctx context.Context, options *AnalyticsOptions) (*OpensRoot, *Response, error) {
	path := fmt.Sprintf("%s/ua-type", analyticsBasePath)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(OpensRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
