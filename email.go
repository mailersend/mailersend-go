package mailersend

import (
	"context"
	"net/http"
)

const emailBasePath = "/email"

type EmailService interface {
	NewMessage() *Message
	Send(ctx context.Context, message *Message) (*Response, error)
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
