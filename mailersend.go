package mailersend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Debug http client
var Debug = false

const (
	// APIBase - base mailersend url
	APIBase = "https://api.mailersend.com/v1"
)

// Mailersend - base mailersend api client
type Mailersend struct {
	apiBase string
	apiKey  string
	client  *http.Client
}

// NewMailersend creates a new client instance.
func NewMailersend(apiKey string) *Mailersend {
	return &Mailersend{
		apiBase: APIBase,
		apiKey:  apiKey,
		client:  http.DefaultClient,
	}
}

// APIKey - Get api key after it has been created
func (ms *Mailersend) APIKey() string {
	return ms.apiKey
}

// Client - Get the current client
func (ms *Mailersend) Client() *http.Client {
	return ms.client
}

// SetClient - Set the client if you want more control over the client implementation
func (ms *Mailersend) SetClient(c *http.Client) {
	ms.client = c
}

// Send - send the message
func (ms *Mailersend) Send(ctx context.Context, message *Message) (res *http.Response, err error) {
	req, err := ms.newRequest("POST", "/email", message)
	if err != nil {
		return nil, err
	}
	res, err = ms.do(ctx, req)
	return res, err
}

func (ms *Mailersend) newRequest(method, path string, message *Message) (*http.Request, error) {
	u := fmt.Sprintf("%s%s", ms.apiBase, path)

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(message)

	req, err := http.NewRequest(method, u, reqBodyBytes)

	if err != nil {
		return nil, err
	}

	if message != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Add("Authorization", "Bearer "+ms.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mailersend-Client-Golang-v1")

	return req, nil
}

func (ms *Mailersend) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := ms.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}
