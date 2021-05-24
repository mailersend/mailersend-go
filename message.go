package mailersend

// Message structures contain both the message text and the envelop for an e-mail message.
type Message struct {
	Recipients []Recipient `json:"to"`
	From       From        `json:"from"`
	CC         []Recipient `json:"cc,omitempty"`
	Bcc        []Recipient `json:"bcc,omitempty"`
	Subject    string      `json:"subject,omitempty"`
	Text       string      `json:"text,omitempty"`
	HTML       string      `json:"html,omitempty"`
	TemplateID string      `json:"template_id,omitempty"`
	Tags       []string    `json:"tags,omitempty"`

	TemplateVariables []Variables       `json:"variables"`
	Personalization   []Personalization `json:"personalization"`
}

// From - simple struct to declare from name/ email
type From struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

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

// NewMessage - Setup a new message ready to be sent.
func (ms *Mailersend) NewMessage() *Message {
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

// SetBcc - Set BCC.
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
func (m *Message) SetTemplateID(templateid string) {
	m.TemplateID = templateid
}

// SetSubstitutions - Set the template substitutions(.
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
