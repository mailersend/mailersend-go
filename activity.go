package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const activityBasePath = "/activity"

type ActivityService service

// activityRoot - format of activity response
type activityRoot struct {
	Data  []activityData `json:"data"`
	Links Links          `json:"links"`
	Meta  Meta           `json:"meta"`
}

type activityData struct {
	ID        string        `json:"id"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
	Type      string        `json:"type"`
	Email     activityEmail `json:"email"`
}

type activityEmail struct {
	ID        string            `json:"id"`
	From      string            `json:"from"`
	Subject   string            `json:"subject"`
	Text      string            `json:"text"`
	HTML      string            `json:"html"`
	Status    string            `json:"status"`
	Tags      interface{}       `json:"tags"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
	Recipient activityRecipient `json:"recipient"`
}

type activityRecipient struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// ActivityOptions - modifies the behavior of ActivityService.List method
type ActivityOptions struct {
	DomainID string   `url:"-,"`
	Page     int      `url:"page,omitempty"`
	DateFrom int64    `url:"date_from,omitempty"`
	DateTo   int64    `url:"date_to,omitempty"`
	Limit    int      `url:"limit,omitempty"`
	Event    []string `url:"event[],omitempty"`
}

func (s *ActivityService) List(ctx context.Context, options *ActivityOptions) (*activityRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", activityBasePath, options.DomainID)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(activityRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
