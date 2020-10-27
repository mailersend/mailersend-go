package mailersend

// Message structures contain both the message text and the envelop for an e-mail message.
type Message struct {
	Recipients []Recipient
	Sender     From
	subject    string
	text       string
	html       string
	templateID int

	// todo templateVariables map[string]interface{}
	ms *Mailersend
}

// From - simple struct to declare from name/ email
type From struct {
	Name  string
	Email string
}

// Recipient - you can set multiple recipients
type Recipient struct {
	Name  string
	Email string
}

// NewMessage - Setup a new message ready to be sent.
func (ms *Mailersend) NewMessage(from From, subject string, text string, html string) *Message {
	return &Message{
		Sender:  from,
		subject: subject,
		text:    text,
		html:    html,
		ms:      ms,
	}
}

// SetRecipients - Set all teh recipients.
func (m *Message) SetRecipients(recipients []Recipient) {
	m.Recipients = recipients
}

// Send - send the message (TODO)
func (m *Message) Send() bool {
	// Implement sending the message here
	return true
}
