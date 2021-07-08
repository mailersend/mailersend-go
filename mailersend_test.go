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

//RoundTripFunc
type RoundTripFunc func(req *http.Request) *http.Response

//RoundTrip
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

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
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	ctx := context.TODO()

	ms.SetClient(client)

	message := ms.NewMessage()

	_, err := ms.Send(ctx, message)
	if err != nil {
		return
	}

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

	_, err := ms.Send(ctx, message)

	assert.Error(t, err)

}

func TestCanSetApiKey(t *testing.T) {

	ms := mailersend.NewMailersend("api-key")

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.Header.Get("Authorization"), "Bearer api-key")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	ctx := context.TODO()

	ms.SetClient(client)

	message := ms.NewMessage()

	_, err := ms.Send(ctx, message)
	if err != nil {
		return
	}

}

func TestCanSetApiKeyNew(t *testing.T) {

	ms := mailersend.NewMailersend("api-key")

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.Header.Get("Authorization"), "Bearer api-key")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	ctx := context.TODO()

	ms.SetClient(client)

	message := ms.Email.NewMessage()

	_, err := ms.Email.Send(ctx, message)
	if err != nil {
		return
	}

}
