package mailersend_test

import (
	"bytes"
	"context"
	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
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
		assert.Equal(t, req.URL.String(), "https://api.mailersend.com/v1/email")
		return &http.Response{
			StatusCode: http.StatusAccepted,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
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
				Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			}
		case "Bearer new-api-key":
			return &http.Response{
				StatusCode: 401,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`ERROR`)),
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
