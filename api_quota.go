package mailersend

import (
	"context"
	"net/http"
	"time"
)

const apiQuotaBasePath = "/api-quota"

type ApiQuotaService interface {
	Get(ctx context.Context) (*ApiQuotaRoot, *Response, error)
}

type apiQuotaService struct {
	*service
}

type ApiQuotaRoot struct {
	Quota     int       `json:"quota"`
	Remaining int       `json:"remaining"`
	Reset     time.Time `json:"reset"`
}

func (s *apiQuotaService) Get(ctx context.Context) (*ApiQuotaRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, apiQuotaBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ApiQuotaRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
