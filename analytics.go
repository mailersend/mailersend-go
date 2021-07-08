package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const analyticsBasePath = "/analytics"

type AnalyticsService service

// analyticsActivityRoot - format of analytics response
type analyticsActivityRoot struct {
	Data analyticsData `json:"data"`
}

type analyticsData struct {
	DateFrom string           `json:"date_from"`
	DateTo   string           `json:"date_to"`
	GroupBy  string           `json:"group_by"`
	Stats    []analyticsStats `json:"stats"`
}

type analyticsStats struct {
	Date           string `json:"date"`
	Processed      int    `json:"processed,omitempty"`
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

type opensRoot struct {
	Data openData `json:"data"`
}

type openData struct {
	DateFrom int         `json:"date_from"`
	DateTo   int         `json:"date_to"`
	Stats    []openStats `json:"stats"`
}

type openStats struct {
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

func (s *AnalyticsService) GetActivityByDate(ctx context.Context, options *AnalyticsOptions) (*analyticsActivityRoot, *Response, error) {
	path := fmt.Sprintf("%s/date", analyticsBasePath)

	req, err := s.client.newRequest("GET", path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(analyticsActivityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *AnalyticsService) GetOpensByCountry(ctx context.Context, options *AnalyticsOptions) (*opensRoot, *Response, error) {
	path := fmt.Sprintf("%s/country", analyticsBasePath)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(opensRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *AnalyticsService) GetOpensByUserAgent(ctx context.Context, options *AnalyticsOptions) (*opensRoot, *Response, error) {
	path := fmt.Sprintf("%s/ua-name", analyticsBasePath)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(opensRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *AnalyticsService) GetOpensByReadingEnvironment(ctx context.Context, options *AnalyticsOptions) (*opensRoot, *Response, error) {
	path := fmt.Sprintf("%s/ua-type", analyticsBasePath)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(opensRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
