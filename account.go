package mailersend

import (
	"context"
	"net/http"
)

const accountBasePath = "/account"

type AccountService interface {
	Get(ctx context.Context) (*AccountRoot, *Response, error)
}

type accountService struct {
	*service
}

// AccountRoot - format of account response
type AccountRoot struct {
	Data Account `json:"data"`
}

// Account is the account behind the API token.
type Account struct {
	Name              string            `json:"name"`
	Status            string            `json:"status"`
	PlanType          string            `json:"plan_type"`
	DataRetentionDays int               `json:"data_retention_days"`
	Emails            AccountEmails     `json:"emails"`
	RateLimits        AccountRateLimits `json:"rate_limits"`
	Resources         AccountResources  `json:"resources"`
	Addons            AccountAddons     `json:"addons"`
}

type AccountEmails struct {
	Sent     int `json:"sent"`
	Received int `json:"received"`
	Limit    int `json:"limit"`
}

type AccountRateLimits struct {
	General   int `json:"general"`
	Email     int `json:"email"`
	BulkEmail int `json:"bulk_email"`
	Activity  int `json:"activity"`
}

type AccountResources struct {
	Domains            AccountResourceUsage         `json:"domains"`
	Users              AccountResourceUsage         `json:"users"`
	FreeAccountantSeat AccountResourceUsage         `json:"free_accountant_seat"`
	Templates          AccountResourceUsage         `json:"templates"`
	SmtpUsers          AccountResourceUsage         `json:"smtp_users"`
	ApiTokens          AccountResourceUsage         `json:"api_tokens"`
	Webhooks           AccountNullableResourceUsage `json:"webhooks"`
}

type AccountResourceUsage struct {
	Used  int `json:"used"`
	Limit int `json:"limit"`
}

type AccountNullableResourceUsage struct {
	Used  int  `json:"used"`
	Limit *int `json:"limit"`
}

type AccountAddons struct {
	Sms               AccountSmsAddon       `json:"sms"`
	WhatsApp          *AccountWhatsAppAddon `json:"whatsapp"`
	PhoneNumbers      int                   `json:"phone_numbers"`
	DomainSpaces      int                   `json:"domain_spaces"`
	DataRetentionDays int                   `json:"data_retention_days"`
	DedicatedIps      int                   `json:"dedicated_ips"`
}

type AccountSmsAddon struct {
	Sent     int `json:"sent"`
	Received int `json:"received"`
	Limit    int `json:"limit"`
}

type AccountWhatsAppAddon struct {
	Sent  int `json:"sent"`
	Limit int `json:"limit"`
}

// Get retrieves the account behind the API token
func (s *accountService) Get(ctx context.Context) (*AccountRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, accountBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AccountRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
