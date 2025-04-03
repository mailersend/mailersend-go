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

type SuppressionService interface {
	ListBlockList(ctx context.Context, options *SuppressionOptions) (*SuppressionBlockListRoot, *Response, error)
	ListHardBounces(ctx context.Context, options *SuppressionOptions) (*SuppressionHardBouncesRoot, *Response, error)
	ListSpamComplaints(ctx context.Context, options *SuppressionOptions) (*SuppressionSpamComplaintsRoot, *Response, error)
	ListUnsubscribes(ctx context.Context, options *SuppressionOptions) (*SuppressionUnsubscribesRoot, *Response, error)
	CreateBlock(ctx context.Context, options *CreateSuppressionBlockOptions) (*SuppressionBlockResponse, *Response, error)
	CreateHardBounce(ctx context.Context, options *CreateSuppressionOptions) (*SuppressionHardBouncesRoot, *Response, error)
	CreateSpamComplaint(ctx context.Context, options *CreateSuppressionOptions) (*SuppressionSpamComplaintsRoot, *Response, error)
	CreateUnsubscribe(ctx context.Context, options *CreateSuppressionOptions) (*SuppressionUnsubscribesRoot, *Response, error)
	Delete(ctx context.Context, options *DeleteSuppressionOptions, suppressionType string) (*Response, error)
	DeleteAll(ctx context.Context, domainID string, suppressionType string) (*Response, error)
}

type suppressionService struct {
	*service
}

// SuppressionBlockListRoot - recipients response
type SuppressionBlockListRoot struct {
	Data  []SuppressionBlockListData `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
}

type SuppressionBlockListData struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Pattern   string    `json:"pattern"`
	Domain    Domain    `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SuppressionHardBouncesRoot - recipients response
type SuppressionHardBouncesRoot struct {
	Data  []SuppressionHardBouncesData `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
}

type SuppressionHardBouncesData struct {
	ID        string               `json:"id"`
	Reason    string               `json:"reason"`
	CreatedAt time.Time            `json:"created_at"`
	Recipient SuppressionRecipient `json:"recipient"`
}

// SuppressionSpamComplaintsRoot - recipients response
type SuppressionSpamComplaintsRoot struct {
	Data  []SuppressionSpamComplaintsData `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
}

type SuppressionSpamComplaintsData struct {
	ID        string               `json:"id"`
	Recipient SuppressionRecipient `json:"recipient"`
	CreatedAt time.Time            `json:"created_at"`
}

// SuppressionUnsubscribesRoot - recipients response
type SuppressionUnsubscribesRoot struct {
	Data  []SuppressionUnsubscribesData `json:"data"`
	Links `json:"links"`
	Meta  `json:"meta"`
}

type SuppressionUnsubscribesData struct {
	ID             string               `json:"id"`
	Reason         string               `json:"reason"`
	ReadableReason string               `json:"readable_reason"`
	Recipient      SuppressionRecipient `json:"recipient"`
	CreatedAt      time.Time            `json:"created_at"`
}

type SuppressionRecipient struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt string    `json:"deleted_at"`
	Domain    Domain    `json:"domain"`
}

type SuppressionBlockResponse struct {
	Data []SuppressionBlockData `json:"data"`
}

type SuppressionBlockData struct {
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

type DeleteAll struct {
	DomainID string `json:"domain_id"`
	All      bool   `json:"all"`
}

func (s *suppressionService) ListBlockList(ctx context.Context, options *SuppressionOptions) (*SuppressionBlockListRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, BlockList)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuppressionBlockListRoot)

	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *suppressionService) ListHardBounces(ctx context.Context, options *SuppressionOptions) (*SuppressionHardBouncesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, HardBounces)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuppressionHardBouncesRoot)

	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *suppressionService) ListSpamComplaints(ctx context.Context, options *SuppressionOptions) (*SuppressionSpamComplaintsRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, SpamComplaints)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuppressionSpamComplaintsRoot)

	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *suppressionService) ListUnsubscribes(ctx context.Context, options *SuppressionOptions) (*SuppressionUnsubscribesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, Unsubscribes)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuppressionUnsubscribesRoot)

	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *suppressionService) CreateBlock(ctx context.Context, options *CreateSuppressionBlockOptions) (*SuppressionBlockResponse, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, BlockList)
	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuppressionBlockResponse)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *suppressionService) CreateHardBounce(ctx context.Context, options *CreateSuppressionOptions) (*SuppressionHardBouncesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, HardBounces)
	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuppressionHardBouncesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *suppressionService) CreateSpamComplaint(ctx context.Context, options *CreateSuppressionOptions) (*SuppressionSpamComplaintsRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, SpamComplaints)
	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuppressionSpamComplaintsRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *suppressionService) CreateUnsubscribe(ctx context.Context, options *CreateSuppressionOptions) (*SuppressionUnsubscribesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, Unsubscribes)
	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuppressionUnsubscribesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *suppressionService) Delete(ctx context.Context, options *DeleteSuppressionOptions, suppressionType string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, suppressionType)

	req, err := s.client.newRequest(http.MethodDelete, path, options)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)

}

func (s *suppressionService) DeleteAll(ctx context.Context, domainID string, suppressionType string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", suppressionBasePath, suppressionType)

	options := DeleteAll{All: true, DomainID: domainID}

	req, err := s.client.newRequest(http.MethodDelete, path, options)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
