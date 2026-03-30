package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const dmarcMonitoringBasePath = "/dmarc-monitoring"

// DmarcMonitoringService defines the interface for DMARC Monitoring API operations.
type DmarcMonitoringService interface {
	List(ctx context.Context, options *ListDmarcMonitorOptions) (*DmarcMonitorRoot, *Response, error)
	Create(ctx context.Context, options *CreateDmarcMonitorOptions) (*SingleDmarcMonitorRoot, *Response, error)
	Update(ctx context.Context, options *UpdateDmarcMonitorOptions) (*SingleDmarcMonitorRoot, *Response, error)
	Delete(ctx context.Context, monitorID string) (*Response, error)
	GetAggregatedReport(ctx context.Context, options *ListDmarcReportOptions) (*DmarcAggregatedReportRoot, *Response, error)
	GetIPReport(ctx context.Context, monitorID string, ip string) (*DmarcIPReportRoot, *Response, error)
	GetReportSources(ctx context.Context, options *ListDmarcReportSourcesOptions) (*DmarcReportSourcesRoot, *Response, error)
	MarkIPFavorite(ctx context.Context, monitorID string, ip string) (*Response, error)
	RemoveIPFavorite(ctx context.Context, monitorID string, ip string) (*Response, error)
}

type dmarcMonitoringService struct {
	*service
}

// DmarcMonitorRoot - list of DMARC monitors response
type DmarcMonitorRoot struct {
	Data  []DmarcMonitor `json:"data"`
	Links Links          `json:"links"`
	Meta  Meta           `json:"meta"`
}

// SingleDmarcMonitorRoot - single DMARC monitor response
type SingleDmarcMonitorRoot struct {
	Data DmarcMonitor `json:"data"`
}

// DmarcMonitor - a single DMARC monitor
type DmarcMonitor struct {
	ID                string `json:"id"`
	DomainID          string `json:"domain_id"`
	Domain            Domain `json:"domain"`
	DmarcRecord       string `json:"dmarc_record"`
	WantedDmarcRecord string `json:"wanted_dmarc_record"`
	IsDmarcVerified   bool   `json:"is_dmarc_verified"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

// DmarcAggregatedReportRoot - aggregated DMARC report response
type DmarcAggregatedReportRoot struct {
	Data  []DmarcAggregatedReport `json:"data"`
	Links Links                   `json:"links"`
	Meta  Meta                    `json:"meta"`
}

// DmarcAggregatedReport - a single aggregated report entry
type DmarcAggregatedReport struct {
	IP             string `json:"ip"`
	Count          int    `json:"count"`
	DmarcPassCount int    `json:"dmarc_pass_count"`
	DmarcFailCount int    `json:"dmarc_fail_count"`
	SpfPassCount   int    `json:"spf_pass_count"`
	SpfFailCount   int    `json:"spf_fail_count"`
	DkimPassCount  int    `json:"dkim_pass_count"`
	DkimFailCount  int    `json:"dkim_fail_count"`
	Disposition    string `json:"disposition"`
	Country        string `json:"country"`
	IsFavorite     bool   `json:"is_favorite"`
	CreatedAt      string `json:"created_at"`
}

// DmarcIPReportRoot - IP-specific DMARC report response
type DmarcIPReportRoot struct {
	Data  []DmarcIPReport `json:"data"`
	Links Links           `json:"links"`
	Meta  Meta            `json:"meta"`
}

// DmarcIPReport - a single IP report entry
type DmarcIPReport struct {
	IP          string `json:"ip"`
	Count       int    `json:"count"`
	Disposition string `json:"disposition"`
	DkimResult  string `json:"dkim_result"`
	SpfResult   string `json:"spf_result"`
	CreatedAt   string `json:"created_at"`
}

// DmarcReportSourcesRoot - DMARC report sources response
type DmarcReportSourcesRoot struct {
	Data  []DmarcReportSource `json:"data"`
	Links Links               `json:"links"`
	Meta  Meta                `json:"meta"`
}

// DmarcReportSource - a single report source
type DmarcReportSource struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// ListDmarcMonitorOptions - modifies the behavior of dmarcMonitoringService.List
type ListDmarcMonitorOptions struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// CreateDmarcMonitorOptions - modifies the behavior of dmarcMonitoringService.Create
type CreateDmarcMonitorOptions struct {
	DomainID string `json:"domain_id"`
}

// UpdateDmarcMonitorOptions - modifies the behavior of dmarcMonitoringService.Update
type UpdateDmarcMonitorOptions struct {
	MonitorID         string `json:"-"`
	WantedDmarcRecord string `json:"wanted_dmarc_record"`
}

// ListDmarcReportOptions - modifies the behavior of dmarcMonitoringService.GetAggregatedReport
type ListDmarcReportOptions struct {
	MonitorID string `url:"-"`
	Page      int    `url:"page,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

// ListDmarcReportSourcesOptions - modifies the behavior of dmarcMonitoringService.GetReportSources
type ListDmarcReportSourcesOptions struct {
	MonitorID string `url:"-"`
	Page      int    `url:"page,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

func (s *dmarcMonitoringService) List(ctx context.Context, options *ListDmarcMonitorOptions) (*DmarcMonitorRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, dmarcMonitoringBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(DmarcMonitorRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *dmarcMonitoringService) Create(ctx context.Context, options *CreateDmarcMonitorOptions) (*SingleDmarcMonitorRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, dmarcMonitoringBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleDmarcMonitorRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *dmarcMonitoringService) Update(ctx context.Context, options *UpdateDmarcMonitorOptions) (*SingleDmarcMonitorRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", dmarcMonitoringBasePath, options.MonitorID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleDmarcMonitorRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *dmarcMonitoringService) Delete(ctx context.Context, monitorID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", dmarcMonitoringBasePath, monitorID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

func (s *dmarcMonitoringService) GetAggregatedReport(ctx context.Context, options *ListDmarcReportOptions) (*DmarcAggregatedReportRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/report", dmarcMonitoringBasePath, options.MonitorID)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(DmarcAggregatedReportRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *dmarcMonitoringService) GetIPReport(ctx context.Context, monitorID string, ip string) (*DmarcIPReportRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/report/%s", dmarcMonitoringBasePath, monitorID, ip)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(DmarcIPReportRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *dmarcMonitoringService) GetReportSources(ctx context.Context, options *ListDmarcReportSourcesOptions) (*DmarcReportSourcesRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s/report-sources", dmarcMonitoringBasePath, options.MonitorID)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(DmarcReportSourcesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *dmarcMonitoringService) MarkIPFavorite(ctx context.Context, monitorID string, ip string) (*Response, error) {
	path := fmt.Sprintf("%s/%s/favorite/%s", dmarcMonitoringBasePath, monitorID, ip)

	req, err := s.client.newRequest(http.MethodPut, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

func (s *dmarcMonitoringService) RemoveIPFavorite(ctx context.Context, monitorID string, ip string) (*Response, error) {
	path := fmt.Sprintf("%s/%s/favorite/%s", dmarcMonitoringBasePath, monitorID, ip)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
