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

func TestCanCreateSmtpUserListOptions(t *testing.T) {
	options := mailersend.ListSmtpUserOptions{
		Page:  1,
		Limit: 25,
	}

	assert.Equal(t, 1, options.Page)
	assert.Equal(t, 25, options.Limit)
}

func TestCanCreateSmtpUserCreateOptions(t *testing.T) {
	options := mailersend.CreateSmtpUserOptions{
		Name:    "SMTP User",
		Enabled: mailersend.Bool(true),
	}

	assert.Equal(t, "SMTP User", options.Name)
	assert.Equal(t, mailersend.Bool(true), options.Enabled)
}

func TestCanCreateSmtpUserUpdateOptions(t *testing.T) {
	options := mailersend.UpdateSmtpUserOptions{
		Name:    "Updated SMTP User",
		Enabled: mailersend.Bool(false),
	}

	assert.Equal(t, "Updated SMTP User", options.Name)
	assert.Equal(t, mailersend.Bool(false), options.Enabled)
}

func TestSmtpUserService_List(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/domains/domain-id/smtp-users?limit=25&page=1", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": [
					{
						"id": "smtp-user-id",
						"name": "SMTP User",
						"username": "smtp_user",
						"enabled": true,
						"created_at": "2023-01-01T00:00:00.000000Z"
					}
				]
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	options := &mailersend.ListSmtpUserOptions{
		Page:  1,
		Limit: 25,
	}

	response, _, err := ms.SmtpUser.List(ctx, "domain-id", options)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, "smtp-user-id", response.Data[0].ID)
}

func TestSmtpUserService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/domains/domain-id/smtp-users/smtp-user-id", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": {
					"id": "smtp-user-id",
					"name": "SMTP User",
					"username": "smtp_user",
					"enabled": true,
					"created_at": "2023-01-01T00:00:00.000000Z"
				}
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	response, _, err := ms.SmtpUser.Get(ctx, "domain-id", "smtp-user-id")

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "smtp-user-id", response.Data.ID)
}

func TestSmtpUserService_Create(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/domains/domain-id/smtp-users", req.URL.String())
		assert.Equal(t, http.MethodPost, req.Method)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": {
					"id": "new-smtp-user-id",
					"name": "New SMTP User",
					"username": "new_smtp_user",
					"enabled": true,
					"created_at": "2023-01-01T00:00:00.000000Z"
				}
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	options := &mailersend.CreateSmtpUserOptions{
		Name:    "New SMTP User",
		Enabled: mailersend.Bool(true),
	}

	response, _, err := ms.SmtpUser.Create(ctx, "domain-id", options)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "new-smtp-user-id", response.Data.ID)
}

func TestSmtpUserService_Update(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/domains/domain-id/smtp-users/smtp-user-id", req.URL.String())
		assert.Equal(t, http.MethodPut, req.Method)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": {
					"id": "smtp-user-id",
					"name": "Updated SMTP User",
					"username": "smtp_user",
					"enabled": false,
					"created_at": "2023-01-01T00:00:00.000000Z"
				}
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	options := &mailersend.UpdateSmtpUserOptions{
		Name:    "Updated SMTP User",
		Enabled: mailersend.Bool(false),
	}

	response, _, err := ms.SmtpUser.Update(ctx, "domain-id", "smtp-user-id", options)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "smtp-user-id", response.Data.ID)
	assert.Equal(t, "Updated SMTP User", response.Data.Name)
}

func TestSmtpUserService_Delete(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/domains/domain-id/smtp-users/smtp-user-id", req.URL.String())
		assert.Equal(t, http.MethodDelete, req.Method)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	_, err := ms.SmtpUser.Delete(ctx, "domain-id", "smtp-user-id")

	assert.NoError(t, err)
}
