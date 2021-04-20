package mailersend_test

import (
	"testing"

	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
)

const (
	fromName  = "Your Name"
	fromEmail = "your@domain.com"

	toName  = "Your Client"
	toEmail = "your@client.com"

	subject = "Subject"
	text    = "This is the text content"
	html    = "<p>This is the HTML content</p>"

	templateID = "testtemplateid"
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
					Var:   "foo",
					Value: "bar",
				},
			},
		},
	}

	tags := []string{"foo", "bar"}

	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTemplateID(templateID)
	message.SetSubstitutions(variables)
	message.SetTags(tags)

	assert.Equal(t, from, message.From)
	assert.Equal(t, recipients[0], message.Recipients[0])
	assert.Equal(t, recipients[1], message.Recipients[1])

	assert.Equal(t, variables, message.TemplateVariables)
	assert.Equal(t, templateID, message.TemplateID)

	assert.Equal(t, subject, message.Subject)
	assert.Equal(t, html, message.HTML)
	assert.Equal(t, text, message.Text)

	assert.Equal(t, tags, message.Tags)

}
