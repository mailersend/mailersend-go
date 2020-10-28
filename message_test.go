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

	templateID = "testtemplateid12"
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

	variables := []mailersend.Variables{
		{
			Email: toEmail,
			Substitutions: []mailersend.Substitution{
				{
					Var:   "test",
					Value: "Dave",
				},
			},
		},
	}

	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTemplateID(templateID)
	message.SetSubstitutions(variables)

	assert.Equal(t, from, message.From)
	assert.Equal(t, recipients[0], message.Recipients[0])
	assert.Equal(t, recipients[1], message.Recipients[1])

	assert.Equal(t, variables, message.TemplateVariables)
	assert.Equal(t, templateID, message.TemplateID)

	assert.Equal(t, subject, message.Subject)
	assert.Equal(t, html, message.HTML)
	assert.Equal(t, text, message.Text)

}
