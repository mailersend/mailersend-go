package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const domainBasePath = "/domains"

type DomainService interface {
	List(ctx context.Context, options *ListDomainOptions) (*DomainRoot, *Response, error)
	Get(ctx context.Context, domainID string) (*SingleDomainRoot, *Response, error)
	Update(ctx context.Context, options *DomainSettingOptions) (*SingleDomainRoot, *Response, error)
	Delete(ctx context.Context, domainID string) (*Response, error)
	Create(ctx context.Context, options *CreateDomainOptions) (*SingleDomainRoot, *Response, error)
	GetDNS(ctx context.Context, domainID string) (*DnsRoot, *Response, error)
	Verify(ctx context.Context, domainID string) (*VerifyRoot, *Response, error)
	GetRecipients(ctx context.Context, options *GetRecipientsOptions) (*DomainRecipientRoot, *Response, error)
}

type domainService struct {
	*service
}

// DomainRoot format of domain response
type DomainRoot struct {
	Data  []Domain `json:"data"`
	Links Links    `json:"links"`
	Meta  Meta     `json:"meta"`
}

// SingleDomainRoot format of domain response
type SingleDomainRoot struct {
	Data Domain `json:"data"`
}

type Domain struct {
	ID                     string         `json:"id"`
	Name                   string         `json:"name"`
	Dkim                   bool           `json:"dkim"`
	Spf                    bool           `json:"spf"`
	Tracking               bool           `json:"tracking"`
	IsVerified             bool           `json:"is_verified"`
	IsCnameVerified        bool           `json:"is_cname_verified"`
	IsDNSActive            bool           `json:"is_dns_active"`
	IsCnameActive          bool           `json:"is_cname_active"`
	IsTrackingAllowed      bool           `json:"is_tracking_allowed"`
	HasNotQueuedMessages   bool           `json:"has_not_queued_messages"`
	NotQueuedMessagesCount int            `json:"not_queued_messages_count"`
	DomainSettings         DomainSettings `json:"domain_settings"`
	CreatedAt              string         `json:"created_at"`
	UpdatedAt              string         `json:"updated_at"`
}

type DomainSettings struct {
	SendPaused                 bool   `json:"send_paused,omitempty"`
	TrackClicks                bool   `json:"track_clicks,omitempty"`
	TrackOpens                 bool   `json:"track_opens,omitempty"`
	TrackUnsubscribe           bool   `json:"track_unsubscribe,omitempty"`
	TrackUnsubscribeHTML       string `json:"track_unsubscribe_html,omitempty"`
	TrackUnsubscribePlain      string `json:"track_unsubscribe_plain,omitempty"`
	TrackContent               bool   `json:"track_content,omitempty"`
	CustomTrackingEnabled      bool   `json:"custom_tracking_enabled,omitempty"`
	CustomTrackingSubdomain    string `json:"custom_tracking_subdomain,omitempty"`
	IgnoreDuplicatedRecipients bool   `json:"ignore_duplicated_recipients,omitempty"`
	PrecedenceBulk             bool   `json:"precedence_bulk,omitempty"`
}

type DnsRoot struct {
	Data Dns `json:"data"`
}

type Spf struct {
	Hostname string `json:"hostname"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

type Dkim struct {
	Hostname string `json:"hostname"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

type ReturnPath struct {
	Hostname string `json:"hostname"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

type CustomTracking struct {
	Hostname string `json:"hostname"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

type InboundRouting struct {
	Hostname string `json:"hostname"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Priority string `json:"priority"`
}

type Dns struct {
	ID             string         `json:"id"`
	Spf            Spf            `json:"spf"`
	Dkim           Dkim           `json:"dkim"`
	ReturnPath     ReturnPath     `json:"return_path"`
	CustomTracking CustomTracking `json:"custom_tracking"`
	InboundRouting InboundRouting `json:"inbound_routing"`
}

type VerifyRoot struct {
	Message string `json:"message"`
	Data    Verify `json:"data"`
}
type Verify struct {
	Dkim     bool `json:"dkim"`
	Spf      bool `json:"spf"`
	Mx       bool `json:"mx"`
	Tracking bool `json:"tracking"`
	Cname    bool `json:"cname"`
	RpCname  bool `json:"rp_cname"`
}

// DomainRecipientRoot format of domain response
type DomainRecipientRoot struct {
	Data  []DomainRecipient `json:"data"`
	Links Links             `json:"links"`
	Meta  Meta              `json:"meta"`
}

// DomainRecipient list of domain recipients
type DomainRecipient struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// ListDomainOptions - modifies the behavior of domainService.List Method
type ListDomainOptions struct {
	Page     int   `url:"page,omitempty"`
	Limit    int   `url:"limit,omitempty"`
	Verified *bool `url:"verified,omitempty"`
}

// DomainSettingOptions - modifies the behavior of domainService.Update Method
type DomainSettingOptions struct {
	DomainID                   string `json:"-"`
	SendPaused                 *bool  `json:"send_paused,omitempty"`
	TrackClicks                *bool  `json:"track_clicks,omitempty"`
	TrackOpens                 *bool  `json:"track_opens,omitempty"`
	TrackUnsubscribe           *bool  `json:"track_unsubscribe,omitempty"`
	TrackUnsubscribeHTML       string `json:"track_unsubscribe_html,omitempty"`
	TrackUnsubscribePlain      string `json:"track_unsubscribe_plain,omitempty"`
	TrackContent               *bool  `json:"track_content,omitempty"`
	CustomTrackingEnabled      *bool  `json:"custom_tracking_enabled,omitempty"`
	CustomTrackingSubdomain    string `json:"custom_tracking_subdomain,omitempty"`
	IgnoreDuplicatedRecipients *bool  `json:"ignore_duplicated_recipients,omitempty"`
	PrecedenceBulk             *bool  `json:"precedence_bulk,omitempty"`
}

type CreateDomainOptions struct {
	Name                    string `json:"name"`
	ReturnPathSubdomain     string `json:"return_path_subdomain,omitempty"`
	CustomTrackingSubdomain string `json:"custom_tracking_subdomain,omitempty"`
	InboundRoutingSubdomain string `json:"inbound_routing_subdomain,omitempty"`
}

// GetRecipientsOptions - modifies the behavior of domainService.GetRecipients Method
type GetRecipientsOptions struct {
	DomainID string `url:"-"`
	Page     int    `url:"page,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

func (s *domainService) List(ctx context.Context, options *ListDomainOptions) (*DomainRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, domainBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(DomainRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *domainService) Get(ctx context.Context, domainID string) (*SingleDomainRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", domainBasePath, domainID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleDomainRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *domainService) Update(ctx context.Context, options *DomainSettingOptions) (*SingleDomainRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/settings", domainBasePath, options.DomainID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleDomainRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *domainService) Delete(ctx context.Context, domainID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", domainBasePath, domainID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

func (s *domainService) Create(ctx context.Context, options *CreateDomainOptions) (*SingleDomainRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, domainBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleDomainRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *domainService) GetDNS(ctx context.Context, domainID string) (*DnsRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/dns-records", domainBasePath, domainID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(DnsRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *domainService) Verify(ctx context.Context, domainID string) (*VerifyRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/verify", domainBasePath, domainID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(VerifyRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *domainService) GetRecipients(ctx context.Context, options *GetRecipientsOptions) (*DomainRecipientRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/recipients", domainBasePath, options.DomainID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(DomainRecipientRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
