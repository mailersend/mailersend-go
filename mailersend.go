package mailersend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const APIBase string = "https://api.mailersend.com/v1"

// Mailersend - base mailersend api client
type Mailersend struct {
	apiBase string
	apiKey  string
	client  *http.Client

	common service // Reuse a single struct.

	// Services
	Activity          ActivityService
	Analytics         AnalyticsService
	Domain            DomainService
	Email             EmailService
	BulkEmail         BulkEmailService
	Message           MessageService
	ScheduleMessage   ScheduleMessageService
	Recipient         RecipientService
	Template          TemplateService
	Token             TokenService
	Webhook           WebhookService
	Suppression       SuppressionService
	Inbound           InboundService
	Sms               SmsService
	SmsActivity       SmsActivityService
	SmsNumber         SmsNumberService
	SmsRecipient      SmsRecipientService
	SmsWebhook        SmsWebhookService
	SmsMessage        SmsMessageService
	SmsInbound        SmsInboundService
	EmailVerification EmailVerificationService
	Identity          IdentityService
	ApiQuota          ApiQuotaService
}

type service struct {
	client *Mailersend
}

// Response is a Mailersend API response. This wraps the standard http.Response
// returned from Mailersend and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

// AuthError occurs when using HTTP Authentication fails
type AuthError ErrorResponse

func (r *AuthError) Error() string { return (*ErrorResponse)(r).Error() }

// Meta - used for api responses
type Meta struct {
	CurrentPage json.Number `json:"current_page"`
	From        json.Number `json:"from"`
	Path        string      `json:"path"`
	PerPage     json.Number `json:"per_page"`
	To          json.Number `json:"to"`
}

// Links - used for api responses
type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

// Filter - used to filter resources
type Filter struct {
	Comparer string `json:"comparer"`
	Value    string `json:"value"`
	Key      string `json:"key,omitempty"`
}

// NewMailersend - creates a new client instance.
func NewMailersend(apiKey string) *Mailersend {
	ms := &Mailersend{
		apiBase: APIBase,
		apiKey:  apiKey,
		client:  http.DefaultClient,
	}

	ms.common.client = ms
	ms.Activity = &activityService{&ms.common}
	ms.Analytics = &analyticsService{&ms.common}
	ms.Domain = &domainService{&ms.common}
	ms.Email = &emailService{&ms.common}
	ms.BulkEmail = &bulkEmailService{&ms.common}
	ms.Message = &messageService{&ms.common}
	ms.ScheduleMessage = &scheduleMessageService{&ms.common}
	ms.Recipient = &recipientService{&ms.common}
	ms.Template = &templateService{&ms.common}
	ms.Token = &tokenService{&ms.common}
	ms.Webhook = &webhookService{&ms.common}
	ms.Suppression = &suppressionService{&ms.common}
	ms.Inbound = &inboundService{&ms.common}
	ms.Sms = &smsService{&ms.common}
	ms.SmsActivity = &smsActivityService{&ms.common}
	ms.SmsNumber = &smsNumberService{&ms.common}
	ms.SmsRecipient = &smsRecipientService{&ms.common}
	ms.SmsWebhook = &smsWebhookService{&ms.common}
	ms.SmsMessage = &smsMessageService{&ms.common}
	ms.SmsInbound = &smsInboundService{&ms.common}
	ms.EmailVerification = &emailVerificationService{&ms.common}
	ms.Identity = &identityService{&ms.common}
	ms.ApiQuota = &apiQuotaService{&ms.common}

	return ms
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

// SetAPIKey - Set the client api key
func (ms *Mailersend) SetAPIKey(apikey string) {
	ms.apiKey = apikey
}

func (ms *Mailersend) newRequest(method, path string, body interface{}) (*http.Request, error) {
	reqURL := fmt.Sprintf("%s%s", ms.apiBase, path)
	reqBodyBytes := new(bytes.Buffer)

	if method == http.MethodPost ||
		method == http.MethodPut ||
		method == http.MethodDelete {
		err := json.NewEncoder(reqBodyBytes).Encode(body)
		if err != nil {
			return nil, err
		}
	} else if method == http.MethodGet {
		reqURL, _ = addOptions(reqURL, body)
	}

	req, err := http.NewRequest(method, reqURL, reqBodyBytes)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ms.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mailersend-Client-Golang-v1")

	return req, nil
}

func (ms *Mailersend) do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
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

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		defer resp.Body.Close()
		_, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return response, readErr
		}
	}

	return response, err
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range or equal to 202 Accepted.
func CheckResponse(r *http.Response) error {
	if r.StatusCode == http.StatusAccepted {
		return nil
	}
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	switch {
	case r.StatusCode == http.StatusUnauthorized:
		return (*AuthError)(errorResponse)
	default:
		return errorResponse
	}
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()

	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		origValues[k] = v
	}

	origURL.RawQuery = origValues.Encode()
	return origURL.String(), nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
