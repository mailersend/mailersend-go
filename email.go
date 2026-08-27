package mailersend

import (
	"context"
	"fmt"
	"net/http"
)

const emailBasePath = "/email"
const emailsBasePath = "/emails"

type EmailService interface {
	NewMessage() *Message
	Send(ctx context.Context, message *Message) (*Response, error)
	List(ctx context.Context, options *ListEmailOptions) (*EmailRoot, *Response, error)
	Get(ctx context.Context, emailID string) (*SingleEmailRoot, *Response, error)
}

type emailService struct {
	*service
}

const (
	DispositionInline     = "inline"
	DispositionAttachment = "attachment"
)

// Message structures contain both the message text and the envelop for an e-mail message.
type Message struct {
	Recipients  []Recipient  `json:"to"`
	From        From         `json:"from"`
	CC          []Recipient  `json:"cc,omitempty"`
	Bcc         []Recipient  `json:"bcc,omitempty"`
	ReplyTo     ReplyTo      `json:"reply_to,omitempty"`
	InReplyTo   string       `json:"in_reply_to,omitempty"`
	Subject     string       `json:"subject,omitempty"`
	Text        string       `json:"text,omitempty"`
	HTML        string       `json:"html,omitempty"`
	TemplateID  string       `json:"template_id,omitempty"`
	Language    string       `json:"language,omitempty"`
	SendAt      int64        `json:"send_at,omitempty"`
	Tags        []string     `json:"tags,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`

	TemplateVariables []Variables       `json:"variables"`
	Personalization   []Personalization `json:"personalization"`
	Headers           []Header          `json:"headers"`
	ListUnsubscribe   string            `json:"list_unsubscribe"`
	PrecedenceBulk    bool              `json:"precedence_bulk,omitempty"`
	References        []string          `json:"references,omitempty"`
	Settings          Settings          `json:"settings,omitempty"`
}

// From - simple struct to declare from name/ email
type From = Recipient

// ReplyTo - simple struct to declare from name/ email
type ReplyTo = Recipient

// Recipient - you can set multiple recipients
type Recipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Deprecated: Variables - you can set multiple Substitutions for each Recipient
type Variables struct {
	Email         string         `json:"email"`
	Substitutions []Substitution `json:"substitutions"`
}

// Deprecated: Substitution - you can set multiple Substitutions for each Recipient
type Substitution struct {
	Var   string `json:"var"`
	Value string `json:"value"`
}

// Personalization - you can set multiple Personalization for each Recipient
type Personalization struct {
	Email string                 `json:"email"`
	Data  map[string]interface{} `json:"data"`
}

// Header - you can set multiple Personalization for each Recipient
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Attachment - you can set multiple Attachments
type Attachment struct {
	Content     string `json:"content"`
	Filename    string `json:"filename"`
	Disposition string `json:"disposition,omitempty"`
	ID          string `json:"id,omitempty"`
}

// Settings - you can set email Settings
type Settings struct {
	TrackClicks  bool `json:"track_clicks"`
	TrackOpens   bool `json:"track_opens"`
	TrackContent bool `json:"track_content"`
}

// Deprecated: NewMessage - Setup a new message ready to be sent
func (ms *Mailersend) NewMessage() *Message {
	return &Message{}
}

// NewMessage - Setup a new email message ready to be sent.
func (s *emailService) NewMessage() *Message {
	return &Message{}
}

// SetFrom - Set from.
func (m *Message) SetFrom(from From) {
	m.From = from
}

// SetRecipients - Set all the recipients.
func (m *Message) SetRecipients(recipients []Recipient) {
	m.Recipients = recipients
}

// SetCc - Set CC.
func (m *Message) SetCc(cc []Recipient) {
	m.CC = cc
}

// SetBcc - Set Bcc.
func (m *Message) SetBcc(bcc []Recipient) {
	m.Bcc = bcc
}

// SetReplyTo - Set ReplyTo.
func (m *Message) SetReplyTo(replyTo Recipient) {
	m.ReplyTo = replyTo
}

// SetInReplyTo - Set InReplyTo.
func (m *Message) SetInReplyTo(inReplyTo string) {
	m.InReplyTo = inReplyTo
}

// SetSubject - Set the subject of the email, required if not using a template.
func (m *Message) SetSubject(subject string) {
	m.Subject = subject
}

// SetHTML - Set the html content of the email, required if not using a template.
func (m *Message) SetHTML(html string) {
	m.HTML = html
}

// SetText - Set the text content of the email, required if not using a template.
func (m *Message) SetText(text string) {
	m.Text = text
}

// SetTemplateID - Set the template ID.
func (m *Message) SetTemplateID(templateID string) {
	m.TemplateID = templateID
}

// SetLanguage - Set the language used for the template (language code, e.g. "de", "pt-BR").
// Only meaningful when a template is set; ignored for raw html/text sends.
func (m *Message) SetLanguage(language string) {
	m.Language = language
}

// Deprecated: SetSubstitutions - Set the template substitutions.
func (m *Message) SetSubstitutions(variables []Variables) {
	m.TemplateVariables = variables
}

// SetPersonalization - Set the template personalization.
func (m *Message) SetPersonalization(personalization []Personalization) {
	m.Personalization = personalization
}

// SetHeaders - Set the custom headers.
func (m *Message) SetHeaders(headers []Header) {
	m.Headers = headers
}

// SetListUnsubscribe - Set the custom list unsubscribe header (Professional and Enterprise accounts only)
func (m *Message) SetListUnsubscribe(listUnsubscribe string) {
	m.ListUnsubscribe = listUnsubscribe
}

// SetTags - Set all the tags.
func (m *Message) SetTags(tags []string) {
	m.Tags = tags
}

// AddAttachment - Add an attachment base64 encoded content.
func (m *Message) AddAttachment(attachment Attachment) {
	m.Attachments = append(m.Attachments, attachment)
}

// SetSendAt - Set send_at.
func (m *Message) SetSendAt(sendAt int64) {
	m.SendAt = sendAt
}

// SetPrecedenceBulk - Set precedence_bulk
func (m *Message) SetPrecedenceBulk(precedenceBulk bool) {
	m.PrecedenceBulk = precedenceBulk
}

// SetReferences - Set references
func (m *Message) SetReferences(references []string) {
	m.References = references
}

// AddReference - Add a reference
func (m *Message) AddReference(reference string) {
	m.References = append(m.References, reference)
}

// SetSettings - Set settings
func (m *Message) SetSettings(settings Settings) {
	m.Settings = settings
}

// Deprecated: Send - send the message
func (ms *Mailersend) Send(ctx context.Context, message *Message) (*Response, error) {
	req, err := ms.newRequest(http.MethodPost, emailBasePath, message)
	if err != nil {
		return nil, err
	}

	return ms.do(ctx, req, nil)
}

// Send - send the message
func (s *emailService) Send(ctx context.Context, message *Message) (*Response, error) {
	req, err := s.client.newRequest(http.MethodPost, emailBasePath, message)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

// EmailRoot - format of the emails list response. Links.Last is always empty
// because the API does not report a last page for this endpoint.
type EmailRoot struct {
	Data  []EmailData `json:"data"`
	Links Links       `json:"links"`
	Meta  Meta        `json:"meta"`
}

// EmailData - a single email in the emails list. Text and HTML are always nil in
// list rows, use EmailService.Get to read the content. Tags is nil when the email
// was sent without tags, TemplateID is nil when no template was used and
// SuppressionReason is only set when Status is "rejected".
//
// Headers is a list of {name, value} objects, the same shape as Message.Headers,
// not a flat name-to-value map.
type EmailData struct {
	ID                string      `json:"id"`
	From              string      `json:"from"`
	To                string      `json:"to"`
	Subject           string      `json:"subject"`
	Text              *string     `json:"text"`
	HTML              *string     `json:"html"`
	TemplateID        *string     `json:"template_id"`
	DomainID          string      `json:"domain_id"`
	MessageID         string      `json:"message_id"`
	Status            string      `json:"status"`
	Tags              []string    `json:"tags"`
	Interaction       []string    `json:"interaction"`
	SuppressionReason *string     `json:"suppression_reason"`
	CreatedAt         string      `json:"created_at"`
	UpdatedAt         string      `json:"updated_at"`
	Headers           interface{} `json:"headers"`
}

// SingleEmailRoot - format of the single email response
type SingleEmailRoot struct {
	Data SingleEmail `json:"data"`
}

// SingleEmail - a single email with its content and its activity events.
// Text and HTML are nil when content tracking is disabled for the domain and
// TemplateID is nil when no template was used; Activity is returned either way.
//
// Headers is a list of {name, value} objects, the same shape as Message.Headers,
// not a flat name-to-value map.
type SingleEmail struct {
	ID                string            `json:"id"`
	From              string            `json:"from"`
	To                string            `json:"to"`
	Subject           string            `json:"subject"`
	Text              *string           `json:"text"`
	HTML              *string           `json:"html"`
	TemplateID        *string           `json:"template_id"`
	DomainID          string            `json:"domain_id"`
	MessageID         string            `json:"message_id"`
	Status            string            `json:"status"`
	Tags              []string          `json:"tags"`
	Interaction       []string          `json:"interaction"`
	SuppressionReason *string           `json:"suppression_reason"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	Recipient         ActivityRecipient `json:"recipient"`
	Headers           interface{}       `json:"headers"`
	Activity          []EmailActivity   `json:"activity"`
}

// EmailActivity - an activity event recorded for an email, newest first and capped
// at 200 events per email. SuppressionReason is only present on "suppressed" events
// and is one of on_hold, hard_bounced, unsubscribed, spam_complained or blocklisted.
type EmailActivity struct {
	ID                string  `json:"id"`
	Type              string  `json:"type"`
	CreatedAt         string  `json:"created_at"`
	SuppressionReason *string `json:"suppression_reason,omitempty"`
}

// ListEmailOptions - modifies the behavior of EmailService.List method.
//
// DomainID, DateFrom and DateTo are required. Status and Interaction are always
// sent as arrays, as the API rejects scalar values for them.
//
// Page is 1-based and capped at 1000, Limit is between 10 and 100 and defaults
// to 25.
type ListEmailOptions struct {
	DomainID       string   `url:"domain_id"`
	DateFrom       int64    `url:"date_from"`
	DateTo         int64    `url:"date_to"`
	Page           int      `url:"page,omitempty"`
	Limit          int      `url:"limit,omitempty"`
	Status         []string `url:"status[],omitempty"`
	Interaction    []string `url:"interaction[],omitempty"`
	RecipientEmail string   `url:"recipient_email,omitempty"`
	MessageID      string   `url:"message_id,omitempty"`
	TemplateID     string   `url:"template_id,omitempty"`
	Subject        string   `url:"subject,omitempty"`
	Tag            string   `url:"tag,omitempty"`
}

// List - get a list of emails sent from a domain, newest first.
func (s *emailService) List(ctx context.Context, options *ListEmailOptions) (*EmailRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, emailsBasePath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(EmailRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Get - get a single email together with its activity events.
func (s *emailService) Get(ctx context.Context, emailID string) (*SingleEmailRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", emailBasePath, emailID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(SingleEmailRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
