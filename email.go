package mailersend

import (
	"context"
	"net/http"
)

const emailBasePath = "/email"

type EmailService service

// Message structures contain both the message text and the envelope for an e-mail message.
type Message struct {
	Recipients  []Recipient  `json:"to"`
	From        From         `json:"from"`
	CC          []Recipient  `json:"cc,omitempty"`
	Bcc         []Recipient  `json:"bcc,omitempty"`
	Subject     string       `json:"subject,omitempty"`
	Text        string       `json:"text,omitempty"`
	HTML        string       `json:"html,omitempty"`
	TemplateID  string       `json:"template_id,omitempty"`
	Tags        []string     `json:"tags,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`

	TemplateVariables []Variables       `json:"variables"`
	Personalization   []Personalization `json:"personalization"`
}

// From - simple struct to declare from name/ email
type From = Recipient

// Recipient - you can set multiple recipients
type Recipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Variables - you can set multiple Substitutions for each Recipient
type Variables struct {
	Email         string         `json:"email"`
	Substitutions []Substitution `json:"substitutions"`
}

// Substitution - you can set multiple Substitutions for each Recipient
type Substitution struct {
	Var   string `json:"var"`
	Value string `json:"value"`
}

// Personalization - you can set multiple Personalization for each Recipient
type Personalization struct {
	Email string                 `json:"email"`
	Data  map[string]interface{} `json:"data"`
}

// Attachment - you can set multiple Attachments
type Attachment struct {
	Content  string `json:"content"`
	Filename string `json:"filename"`
	ID       string `json:"id,omitempty"`
}

// Deprecated: NewMessage - Setup a new message ready to be sent
func (ms *Mailersend) NewMessage() *Message {
	return &Message{}
}

// NewMessage - Setup a new email message ready to be sent.
func (s *EmailService) NewMessage() *Message {
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

// SetSubstitutions - Set the template substitutions.
func (m *Message) SetSubstitutions(variables []Variables) {
	m.TemplateVariables = variables
}

// SetPersonalization - Set the template personalization.
func (m *Message) SetPersonalization(personalization []Personalization) {
	m.Personalization = personalization
}

// SetTags - Set all the tags.
func (m *Message) SetTags(tags []string) {
	m.Tags = tags
}

// AddAttachment - Add an attachment base64 encoded content.
func (m *Message) AddAttachment(attachment Attachment) {
	m.Attachments = append(m.Attachments, attachment)
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
func (s *EmailService) Send(ctx context.Context, message *Message) (*Response, error) {
	req, err := s.client.newRequest(http.MethodPost, emailBasePath, message)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
