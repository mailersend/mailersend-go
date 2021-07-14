package mailersend_test

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"
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

	templateID = "123"
)

var from = mailersend.From{
	Name:  fromName,
	Email: fromEmail,
}

var recipients = []mailersend.Recipient{
	{
		Name:  toName,
		Email: toEmail,
	},
}

var cc = []mailersend.Recipient{
	{
		Name:  "CC 1" + toName,
		Email: "cc1-" + toEmail,
	},
	{
		Name:  "CC 2" + toName,
		Email: "cc2-" + toEmail,
	},
}

var bcc = []mailersend.Recipient{
	{
		Name:  "BCC " + toName,
		Email: "bcc-" + toEmail,
	},
}

func basicEmail() *mailersend.Message {
	ms := mailersend.NewMailersend(testKey)

	message := ms.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)

	return message
}

func basicEmailNew() *mailersend.Message {
	ms := mailersend.NewMailersend(testKey)

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)

	return message
}

func TestSimpleMessage(t *testing.T) {
	message := basicEmail()

	assert.Equal(t, from, message.From)
	assert.Equal(t, recipients[0], message.Recipients[0])
	assert.Equal(t, subject, message.Subject)
	assert.Equal(t, html, message.HTML)
	assert.Equal(t, text, message.Text)
}

func TestSimpleMessageNew(t *testing.T) {
	message := basicEmailNew()

	assert.Equal(t, from, message.From)
	assert.Equal(t, recipients[0], message.Recipients[0])
	assert.Equal(t, subject, message.Subject)
	assert.Equal(t, html, message.HTML)
	assert.Equal(t, text, message.Text)
}

func TestCanCCMessage(t *testing.T) {
	message := basicEmail()
	message.SetCc(cc)

	assert.Equal(t, cc, message.CC)
}

func TestCanBCCMessage(t *testing.T) {
	message := basicEmail()
	message.SetBcc(bcc)

	assert.Equal(t, bcc, message.Bcc)
}

func TestCanCCBCCMessage(t *testing.T) {
	message := basicEmailNew()
	message.SetCc(cc)
	message.SetBcc(bcc)

	assert.Equal(t, cc, message.CC)
	assert.Equal(t, bcc, message.Bcc)
}

func TestTemplateMessage(t *testing.T) {
	message := basicEmail()

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

	personalization := []mailersend.Personalization{
		{
			Email: toEmail,
			Data: map[string]interface{}{
				"Var":   "foo",
				"Value": "bar",
			},
		},
	}

	tags := []string{"foo", "bar"}

	message.SetTemplateID(templateID)
	message.SetSubstitutions(variables)
	message.SetPersonalization(personalization)
	message.SetTags(tags)

	assert.Equal(t, variables, message.TemplateVariables)
	assert.Equal(t, templateID, message.TemplateID)
	assert.Equal(t, personalization, message.Personalization)
	assert.Equal(t, tags, message.Tags)
}

func TestFullMessage(t *testing.T) {
	message := basicEmail()

	message.SetCc(cc)
	message.SetBcc(bcc)

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

	personalization := []mailersend.Personalization{
		{
			Email: toEmail,
			Data: map[string]interface{}{
				"Var":   "foo",
				"Value": "bar",
			},
		},
	}

	tags := []string{"foo", "bar"}

	message.SetTemplateID(templateID)
	message.SetSubstitutions(variables)
	message.SetPersonalization(personalization)
	message.SetTags(tags)

	assert.Equal(t, cc, message.CC)
	assert.Equal(t, bcc, message.Bcc)
	assert.Equal(t, from, message.From)
	assert.Equal(t, recipients[0], message.Recipients[0])
	assert.Equal(t, subject, message.Subject)
	assert.Equal(t, html, message.HTML)
	assert.Equal(t, text, message.Text)
	assert.Equal(t, variables, message.TemplateVariables)
	assert.Equal(t, templateID, message.TemplateID)
	assert.Equal(t, personalization, message.Personalization)
	assert.Equal(t, tags, message.Tags)
	assert.Len(t, message.Personalization, 1)

}

func TestFullMessageNew(t *testing.T) {
	message := basicEmailNew()

	message.SetCc(cc)
	message.SetBcc(bcc)

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

	personalization := []mailersend.Personalization{
		{
			Email: toEmail,
			Data: map[string]interface{}{
				"Var":   "foo",
				"Value": "bar",
			},
		},
	}

	tags := []string{"foo", "bar"}

	message.SetTemplateID(templateID)
	message.SetSubstitutions(variables)
	message.SetPersonalization(personalization)
	message.SetTags(tags)

	assert.Equal(t, cc, message.CC)
	assert.Equal(t, bcc, message.Bcc)
	assert.Equal(t, from, message.From)
	assert.Equal(t, recipients[0], message.Recipients[0])
	assert.Equal(t, subject, message.Subject)
	assert.Equal(t, html, message.HTML)
	assert.Equal(t, text, message.Text)
	assert.Equal(t, variables, message.TemplateVariables)
	assert.Equal(t, templateID, message.TemplateID)
	assert.Equal(t, personalization, message.Personalization)
	assert.Equal(t, tags, message.Tags)
	assert.Len(t, message.Personalization, 1)

}

func TestCanAddAttachments(t *testing.T) {
	message := basicEmail()

	f, _ := os.Open("./LICENCE")

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	encoded := base64.StdEncoding.EncodeToString(content)

	attachment := mailersend.Attachment{Filename: "test", Content: encoded}

	message.AddAttachment(attachment)

	assert.NotNil(t, message, message.Attachments)
	assert.Len(t, message.Attachments, 1)
}
