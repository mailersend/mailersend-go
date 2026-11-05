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

const (
	testKey = "valid-mailersend-api-key"
)

func TestNewMailersend(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	assert.Equal(t, ms.APIKey(), testKey)
	assert.Equal(t, ms.Client(), http.DefaultClient)

	client := new(http.Client)
	ms.SetClient(client)
	assert.Equal(t, client, ms.Client())

}

func TestCanMakeMockApiCall(t *testing.T) {

	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.mailersend.com/2026-10/email")
		return &http.Response{
			StatusCode: http.StatusAccepted,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	ms.SetClient(client)

	message := ms.Email.NewMessage()

	res, err := ms.Email.Send(ctx, message)
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusAccepted)

}

func TestWillHandleError(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		// return nil to force error from mock server
		return nil
	})

	ctx := context.TODO()

	ms.SetClient(client)

	message := ms.Email.NewMessage()

	_, err := ms.Email.Send(ctx, message)

	assert.Error(t, err)

}

func TestCanSetApiKey(t *testing.T) {

	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		switch req.Header.Get("Authorization") {
		case "Bearer valid-mailersend-api-key":
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
			}
		case "Bearer new-api-key":
			return &http.Response{
				StatusCode: 401,
				Body:       io.NopCloser(bytes.NewBufferString(`ERROR`)),
			}
		}

		return nil
	})

	ctx := context.TODO()

	ms.SetClient(client)

	message := ms.Email.NewMessage()

	res, _ := ms.Email.Send(ctx, message)

	assert.Equal(t, res.StatusCode, http.StatusOK)

	ms.SetAPIKey("new-api-key")

	resError, _ := ms.Email.Send(ctx, message)

	assert.NotEqual(t, resError.StatusCode, http.StatusAccepted)
	assert.Equal(t, resError.StatusCode, http.StatusUnauthorized)

}

func TestSetAPIVersion(t *testing.T) {
	tests := []struct {
		name    string
		version string
		wantURL string
	}{
		{name: "default", wantURL: "https://api.mailersend.com/2026-10/email"},
		{name: "v1", version: "v1", wantURL: "https://api.mailersend.com/v1/email"},
		{name: "dated", version: "2027-01", wantURL: "https://api.mailersend.com/2027-01/email"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := mailersend.NewMailersend(testKey)

			if tt.version != "" {
				assert.NoError(t, ms.SetAPIVersion(tt.version))
				assert.Equal(t, tt.version, ms.APIVersion())
			} else {
				assert.Equal(t, mailersend.DefaultAPIVersion, ms.APIVersion())
			}

			var requestURL string
			ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
				requestURL = req.URL.String()
				return &http.Response{
					StatusCode: http.StatusAccepted,
					Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
				}
			}))

			_, err := ms.Email.Send(context.TODO(), ms.Email.NewMessage())

			assert.NoError(t, err)
			assert.Equal(t, tt.wantURL, requestURL)
		})
	}
}

func TestSetAPIVersionRejectsInvalidVersion(t *testing.T) {
	for _, version := range []string{"2026-13", "latest", "", "v2", "2026-1", "/2026-10"} {
		t.Run(version, func(t *testing.T) {
			ms := mailersend.NewMailersend(testKey)

			assert.Error(t, ms.SetAPIVersion(version))
			assert.Equal(t, mailersend.DefaultAPIVersion, ms.APIVersion())
		})
	}
}
