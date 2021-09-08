package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const domainBasePath = "/domains"

type DomainService service

// domainRoot format of domain response
type domainRoot struct {
	Data  []Domain `json:"data"`
	Links Links    `json:"links"`
	Meta  Meta     `json:"meta"`
}

// singleDomainRoot format of domain response
type singleDomainRoot struct {
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
	SendPaused              bool   `json:"send_paused,omitempty"`
	TrackClicks             bool   `json:"track_clicks,omitempty"`
	TrackOpens              bool   `json:"track_opens,omitempty"`
	TrackUnsubscribe        bool   `json:"track_unsubscribe,omitempty"`
	TrackUnsubscribeHTML    string `json:"track_unsubscribe_html,omitempty"`
	TrackUnsubscribePlain   string `json:"track_unsubscribe_plain,omitempty"`
	TrackContent            bool   `json:"track_content,omitempty"`
	CustomTrackingEnabled   bool   `json:"custom_tracking_enabled,omitempty"`
	CustomTrackingSubdomain string `json:"custom_tracking_subdomain,omitempty"`
}

type dnsRoot struct {
	Data dns `json:"data"`
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

type dns struct {
	ID             string         `json:"id"`
	Spf            Spf            `json:"spf"`
	Dkim           Dkim           `json:"dkim"`
	ReturnPath     ReturnPath     `json:"return_path"`
	CustomTracking CustomTracking `json:"custom_tracking"`
	InboundRouting InboundRouting `json:"inbound_routing"`
}

type verifyRoot struct {
	Message string `json:"message"`
	Data    verify `json:"data"`
}
type verify struct {
	Dkim     bool `json:"dkim"`
	Spf      bool `json:"spf"`
	Mx       bool `json:"mx"`
	Tracking bool `json:"tracking"`
	Cname    bool `json:"cname"`
	RpCname  bool `json:"rp_cname"`
}

// domainRecipientRoot format of domain response
type domainRecipientRoot struct {
	Data  []domainRecipient `json:"data"`
	Links Links             `json:"links"`
	Meta  Meta              `json:"meta"`
}

// domainRecipient list of domain recipients
type domainRecipient struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// ListDomainOptions - modifies the behavior of DomainService.List Method
type ListDomainOptions struct {
	Page     int   `url:"page,omitempty"`
	Limit    int   `url:"limit,omitempty"`
	Verified *bool `url:"verified,omitempty"`
}

// DomainSettingOptions - modifies the behavior of DomainService.Update Method
type DomainSettingOptions struct {
	DomainID                string `json:"-"`
	SendPaused              *bool  `json:"send_paused,omitempty"`
	TrackClicks             *bool  `json:"track_clicks,omitempty"`
	TrackOpens              *bool  `json:"track_opens,omitempty"`
	TrackUnsubscribe        *bool  `json:"track_unsubscribe,omitempty"`
	TrackUnsubscribeHTML    string `json:"track_unsubscribe_html,omitempty"`
	TrackUnsubscribePlain   string `json:"track_unsubscribe_plain,omitempty"`
	TrackContent            *bool  `json:"track_content,omitempty"`
	CustomTrackingEnabled   *bool  `json:"custom_tracking_enabled,omitempty"`
	CustomTrackingSubdomain string `json:"custom_tracking_subdomain,omitempty"`
}

type CreateDomainOptions struct {
	Name                    string `json:"name"`
	ReturnPathSubdomain     string `json:"return_path_subdomain,omitempty"`
	CustomTrackingSubdomain string `json:"custom_tracking_subdomain,omitempty"`
	InboundRoutingSubdomain string `json:"inbound_routing_subdomain,omitempty"`
}

// GetRecipientsOptions - modifies the behavior of DomainService.GetRecipients Method
type GetRecipientsOptions struct {
	DomainID string `url:"-"`
	Page     int    `url:"page,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

func (s *DomainService) List(ctx context.Context, options *ListDomainOptions) (*domainRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, domainBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(domainRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *DomainService) Get(ctx context.Context, domainID string) (*singleDomainRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", domainBasePath, domainID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleDomainRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *DomainService) Update(ctx context.Context, options *DomainSettingOptions) (*singleDomainRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/settings", domainBasePath, options.DomainID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleDomainRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *DomainService) Delete(ctx context.Context, domainID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", domainBasePath, domainID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

func (s *DomainService) Create(ctx context.Context, options *CreateDomainOptions) (*singleDomainRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, domainBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleDomainRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *DomainService) GetDNS(ctx context.Context, domainID string) (*dnsRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/dns-records", domainBasePath, domainID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(dnsRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *DomainService) Verify(ctx context.Context, domainID string) (*verifyRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/verify", domainBasePath, domainID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(verifyRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *DomainService) GetRecipients(ctx context.Context, options *GetRecipientsOptions) (*domainRecipientRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/recipients", domainBasePath, options.DomainID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(domainRecipientRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
