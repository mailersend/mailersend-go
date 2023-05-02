package mailersend

import (
	"context"
	"net/http"
	"time"
)

const apiQuotaBasePath = "/api-quota"

type ApiQuotaService service

type apiQuotaRoot struct {
	Quota     int       `json:"quota"`
	Remaining int       `json:"remaining"`
	Reset     time.Time `json:"reset"`
}

func (s *ApiQuotaService) Get(ctx context.Context) (*apiQuotaRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, apiQuotaBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(apiQuotaRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
