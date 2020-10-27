package mailersend_test

import (
	"testing"

	"github.com/mailersend/mailersend-go/v1"
	"github.com/stretchr/testify/assert"
)

const (
	fromName  = "MailerSend"
	fromEmail = "example@mailersend.com"

	toName  = "Bob Gordon"
	toEmail = "robert@mailerlite.com"

	subject = "Test Email"
	text    = "This is the message content"
	html    = "<p>This is the html content</p>"
)

func TestMessage(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	from := mailersend.From{
		Name:  fromName,
		Email: fromEmail,
	}

	recipients := []mailersend.Recipient{
		{
			Name:  toName,
			Email: toEmail,
		},
		{
			Name:  toName + "2",
			Email: toEmail,
		},
	}

	message := ms.NewMessage(from, subject, text, html)

	message.SetRecipients(recipients)

	assert.Equal(t, from, message.Sender)
	assert.Equal(t, recipients[0], message.Recipients[0])
	assert.Equal(t, recipients[1], message.Recipients[1])

}
