package mailersend

// Message structures contain both the message text and the envelop for an e-mail message.
type Message struct {
	Recipients []Recipient `json:"to"`
	From       `json:"from"`
	Subject    string `json:"subject,omitempty"`
	Text       string `json:"text,omitempty"`
	HTML       string `json:"html,omitempty"`
	TemplateID string `json:"template_id,omitempty"`

	TemplateVariables []Variables `json:"variables"`
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

// NewMessage - Setup a new message ready to be sent.
func (ms *Mailersend) NewMessage() *Message {
	return &Message{}
}

// SetFrom - Set all teh recipients.
func (m *Message) SetFrom(from From) {
	m.From = from
}

// SetRecipients - Set all the recipients.
func (m *Message) SetRecipients(recipients []Recipient) {
	m.Recipients = recipients
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

// SetTemplateID - Set all teh recipients.
func (m *Message) SetTemplateID(templateid string) {
	m.TemplateID = templateid
}

// SetSubstitutions - Set all teh recipients.
func (m *Message) SetSubstitutions(variables []Variables) {
	m.TemplateVariables = variables
}
