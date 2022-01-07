package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	suppressionBasePath = "/suppressions"

	BlockList      string = "blocklist"
	HardBounces    string = "hard-bounces"
	SpamComplaints string = "spam-complaints"
	Unsubscribes   string = "unsubscribes"
)

type SuppressionService service

// suppressionBlockListRoot - recipients response
type suppressionBlockListRoot struct {
	Data  []suppressionBlockListData `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
}

type suppressionBlockListData struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Pattern   string    `json:"pattern"`
	Domain    Domain    `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// suppressionHardBouncesRoot - recipients response
type suppressionHardBouncesRoot struct {
	Data  []suppressionHardBouncesData `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
}

type suppressionHardBouncesData struct {
	ID        string               `json:"id"`
	Reason    string               `json:"reason"`
	CreatedAt time.Time            `json:"created_at"`
	Recipient suppressionRecipient `json:"recipient"`
}

// suppressionSpamComplaintsRoot - recipients response
type suppressionSpamComplaintsRoot struct {
	Data  []suppressionSpamComplaintsData `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
}

type suppressionSpamComplaintsData struct {
	ID        string               `json:"id"`
	Recipient suppressionRecipient `json:"recipient"`
	CreatedAt time.Time            `json:"created_at"`
}

// suppressionUnsubscribesRoot - recipients response
type suppressionUnsubscribesRoot struct {
	Data  []suppressionUnsubscribesData `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
}

type suppressionUnsubscribesData struct {
	ID             string               `json:"id"`
	Reason         string               `json:"reason"`
	ReadableReason string               `json:"readable_reason"`
	Recipient      suppressionRecipient `json:"recipient"`
	CreatedAt      time.Time            `json:"created_at"`
}

type suppressionRecipient struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt string    `json:"deleted_at"`
	Domain    Domain    `json:"domain"`
}

type suppressionBlockResponse struct {
	Data []suppressionBlockData `json:"data"`
}

type suppressionBlockData struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Pattern   string    `json:"pattern"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSuppressionBlockOptions struct {
	DomainID   string   `json:"domain_id"`
	Recipients []string `json:"recipients,omitempty"`
	Patterns   []string `json:"patterns,omitempty"`
}

type CreateSuppressionOptions struct {
	DomainID   string   `json:"domain_id"`
	Recipients []string `json:"recipients"`
}

// SuppressionOptions - modifies the behavior of SuppressionService.List methods
type SuppressionOptions struct {
	DomainID string `url:"domain_id,omitempty"`
	Page     int    `url:"page,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

type DeleteSuppressionOptions struct {
	DomainID string   `json:"domain_id"`
	Ids      []string `json:"ids"`
}

type deleteAll struct {
	DomainID string `json:"domain_id"`
	All      bool   `json:"all"`
}

func (s *SuppressionService) ListBlockList(ctx context.Context, options *SuppressionOptions) (*suppressionBlockListRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, BlockList)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(suppressionBlockListRoot)

	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SuppressionService) ListHardBounces(ctx context.Context, options *SuppressionOptions) (*suppressionHardBouncesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, HardBounces)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(suppressionHardBouncesRoot)

	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SuppressionService) ListSpamComplaints(ctx context.Context, options *SuppressionOptions) (*suppressionSpamComplaintsRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, SpamComplaints)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(suppressionSpamComplaintsRoot)

	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SuppressionService) ListUnsubscribes(ctx context.Context, options *SuppressionOptions) (*suppressionUnsubscribesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, Unsubscribes)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(suppressionUnsubscribesRoot)

	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SuppressionService) CreateBlock(ctx context.Context, options *CreateSuppressionBlockOptions) (*suppressionBlockResponse, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, BlockList)
	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(suppressionBlockResponse)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SuppressionService) CreateHardBounce(ctx context.Context, options *CreateSuppressionOptions) (*suppressionHardBouncesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, HardBounces)
	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(suppressionHardBouncesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SuppressionService) CreateSpamComplaint(ctx context.Context, options *CreateSuppressionOptions) (*suppressionSpamComplaintsRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, SpamComplaints)
	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(suppressionSpamComplaintsRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SuppressionService) CreateUnsubscribe(ctx context.Context, options *CreateSuppressionOptions) (*suppressionUnsubscribesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, Unsubscribes)
	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(suppressionUnsubscribesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SuppressionService) Delete(ctx context.Context, options *DeleteSuppressionOptions, suppressionType string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, suppressionType)

	req, err := s.client.newRequest(http.MethodDelete, path, options)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)

}

func (s *SuppressionService) DeleteAll(ctx context.Context, domainID string, suppressionType string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, suppressionType)

	options := deleteAll{All: true, DomainID: domainID}

	req, err := s.client.newRequest(http.MethodDelete, path, options)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
