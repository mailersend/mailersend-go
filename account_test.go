package mailersend_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
)

func TestAccountService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/account", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": {
					"name": "Acme Inc",
					"status": "approved",
					"plan_type": "starter",
					"data_retention_days": 7,
					"emails": { "sent": 12000, "received": 430, "limit": -1 },
					"rate_limits": { "general": 60, "email": 120, "bulk_email": 10, "activity": 10 },
					"resources": {
						"domains": { "used": 3, "limit": 10 },
						"users": { "used": 2, "limit": 5 },
						"free_accountant_seat": { "used": 1, "limit": 1 },
						"templates": { "used": 41, "limit": 250 },
						"smtp_users": { "used": 1, "limit": 50 },
						"api_tokens": { "used": 4, "limit": 50 },
						"webhooks": { "used": 2, "limit": 50 }
					},
					"addons": {
						"sms": { "sent": 120, "received": 30, "limit": -1 },
						"whatsapp": { "sent": 800, "limit": 1000 },
						"phone_numbers": 1,
						"domain_spaces": 2,
						"data_retention_days": 30,
						"dedicated_ips": 0
					}
				}
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	response, _, err := ms.Account.Get(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, response)

	account := response.Data
	assert.Equal(t, "Acme Inc", account.Name)
	assert.Equal(t, "approved", account.Status)
	assert.Equal(t, "starter", account.PlanType)
	assert.Equal(t, 7, account.DataRetentionDays)
	assert.Equal(t, mailersend.AccountEmails{Sent: 12000, Received: 430, Limit: -1}, account.Emails)
	assert.Equal(t, mailersend.AccountRateLimits{General: 60, Email: 120, BulkEmail: 10, Activity: 10}, account.RateLimits)
	assert.Equal(t, mailersend.AccountResourceUsage{Used: 3, Limit: 10}, account.Resources.Domains)
	assert.Equal(t, mailersend.AccountResourceUsage{Used: 2, Limit: 5}, account.Resources.Users)
	assert.Equal(t, mailersend.AccountResourceUsage{Used: 1, Limit: 1}, account.Resources.FreeAccountantSeat)
	assert.Equal(t, mailersend.AccountResourceUsage{Used: 41, Limit: 250}, account.Resources.Templates)
	assert.Equal(t, mailersend.AccountResourceUsage{Used: 1, Limit: 50}, account.Resources.SmtpUsers)
	assert.Equal(t, mailersend.AccountResourceUsage{Used: 4, Limit: 50}, account.Resources.ApiTokens)
	assert.Equal(t, 2, account.Resources.Webhooks.Used)
	if assert.NotNil(t, account.Resources.Webhooks.Limit) {
		assert.Equal(t, 50, *account.Resources.Webhooks.Limit)
	}
	assert.Equal(t, mailersend.AccountSmsAddon{Sent: 120, Received: 30, Limit: -1}, account.Addons.Sms)
	if assert.NotNil(t, account.Addons.WhatsApp) {
		assert.Equal(t, mailersend.AccountWhatsAppAddon{Sent: 800, Limit: 1000}, *account.Addons.WhatsApp)
	}
	assert.Equal(t, 1, account.Addons.PhoneNumbers)
	assert.Equal(t, 2, account.Addons.DomainSpaces)
	assert.Equal(t, 30, account.Addons.DataRetentionDays)
	assert.Equal(t, 0, account.Addons.DedicatedIps)
}

func TestAccountService_GetWithoutWhatsAppAddonAndUnlimitedWebhooks(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": {
					"name": "Free Co",
					"status": "pending",
					"plan_type": "free",
					"data_retention_days": 1,
					"emails": { "sent": 0, "received": 0, "limit": 500 },
					"rate_limits": { "general": 60, "email": 120, "bulk_email": 10, "activity": 10 },
					"resources": {
						"domains": { "used": 1, "limit": 1 },
						"users": { "used": 1, "limit": 1 },
						"free_accountant_seat": { "used": 0, "limit": 0 },
						"templates": { "used": 0, "limit": 10 },
						"smtp_users": { "used": 0, "limit": 1 },
						"api_tokens": { "used": 1, "limit": 5 },
						"webhooks": { "used": 0, "limit": null }
					},
					"addons": {
						"sms": { "sent": 0, "received": 0, "limit": 0 },
						"whatsapp": null,
						"phone_numbers": 0,
						"domain_spaces": 0,
						"data_retention_days": 0,
						"dedicated_ips": 0
					}
				}
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	response, _, err := ms.Account.Get(ctx)

	assert.NoError(t, err)
	assert.Equal(t, mailersend.AccountResourceUsage{Used: 0, Limit: 0}, response.Data.Resources.FreeAccountantSeat)
	assert.Nil(t, response.Data.Resources.Webhooks.Limit)
	assert.Nil(t, response.Data.Addons.WhatsApp)
}
