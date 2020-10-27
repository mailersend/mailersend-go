package mailersend

import (
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
