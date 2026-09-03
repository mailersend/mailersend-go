package mailersend_test

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
)

const language = "pt-BR"

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
	message.SetPersonalization(personalization)
	message.SetTags(tags)

	assert.Equal(t, templateID, message.TemplateID)
	assert.Equal(t, personalization, message.Personalization)
	assert.Equal(t, tags, message.Tags)
}

func TestTemplateLanguageMessage(t *testing.T) {
	message := basicEmail()

	message.SetTemplateID(templateID)
	message.SetLanguage(language)

	assert.Equal(t, language, message.Language)

	body, err := json.Marshal(message)
	assert.NoError(t, err)
	assert.Contains(t, string(body), `"language":"pt-BR"`)
}

func TestLanguageOmittedWhenEmpty(t *testing.T) {
	message := basicEmail()

	assert.Empty(t, message.Language)

	body, err := json.Marshal(message)
	assert.NoError(t, err)
	assert.NotContains(t, string(body), `"language"`)
}

func TestFullMessage(t *testing.T) {
	message := basicEmail()

	message.SetCc(cc)
	message.SetBcc(bcc)

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
	message.SetPersonalization(personalization)
	message.SetTags(tags)

	assert.Equal(t, cc, message.CC)
	assert.Equal(t, bcc, message.Bcc)
	assert.Equal(t, from, message.From)
	assert.Equal(t, recipients[0], message.Recipients[0])
	assert.Equal(t, subject, message.Subject)
	assert.Equal(t, html, message.HTML)
	assert.Equal(t, text, message.Text)
	assert.Equal(t, templateID, message.TemplateID)
	assert.Equal(t, personalization, message.Personalization)
	assert.Equal(t, tags, message.Tags)
	assert.Len(t, message.Personalization, 1)

}

func TestFullMessageNew(t *testing.T) {
	message := basicEmailNew()

	message.SetCc(cc)
	message.SetBcc(bcc)

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
	message.SetPersonalization(personalization)
	message.SetTags(tags)

	assert.Equal(t, cc, message.CC)
	assert.Equal(t, bcc, message.Bcc)
	assert.Equal(t, from, message.From)
	assert.Equal(t, recipients[0], message.Recipients[0])
	assert.Equal(t, subject, message.Subject)
	assert.Equal(t, html, message.HTML)
	assert.Equal(t, text, message.Text)
	assert.Equal(t, templateID, message.TemplateID)
	assert.Equal(t, personalization, message.Personalization)
	assert.Equal(t, tags, message.Tags)
	assert.Len(t, message.Personalization, 1)

}

func TestCanAddAttachments(t *testing.T) {
	message := basicEmail()

	f, _ := os.Open("./LICENCE")

	reader := bufio.NewReader(f)
	content, _ := io.ReadAll(reader)

	encoded := base64.StdEncoding.EncodeToString(content)

	attachment := mailersend.Attachment{Filename: "test", Content: encoded}

	message.AddAttachment(attachment)

	assert.NotNil(t, message, message.Attachments)
	assert.Len(t, message.Attachments, 1)
}

const (
	listEmailsDomainID = "7nxe3yjmeq28vp0k"
	listEmailsDateFrom = int64(1756252800)
	listEmailsDateTo   = int64(1756339200)

	emailID = "6a8fa9b1902fab56e0ce50dd"
)

// listEmailsResponse - measured GET v1/emails payload. The second row covers the
// shapes the first one does not: an empty interaction array, a null tags array,
// a null template_id and a set suppression_reason.
const listEmailsResponse = `{
  "data": [
    {
      "id": "6a8fa9b1902fab56e0ce50dd", "from": "sender@example.com", "to": "rcpt@example.org",
      "subject": "Welcome", "text": null, "html": null,
      "template_id": "7nxe3yjmeq28vp0k", "domain_id": "7nxe3yjmeq28vp0k",
      "message_id": "6a8fa9b1902fab56e0ce50aa", "status": "sent",
      "tags": ["newsletter"], "interaction": ["opened"], "suppression_reason": null,
      "created_at": "2026-08-27T16:48:42.000000Z", "updated_at": "2026-08-27T16:48:42.000000Z",
      "headers": [{"name": "X-Custom", "value": "foo"}]
    },
    {
      "id": "6a8fa9b1902fab56e0ce50de", "from": "sender@example.com", "to": "bounce@example.org",
      "subject": "Welcome", "text": null, "html": null,
      "template_id": null, "domain_id": "7nxe3yjmeq28vp0k",
      "message_id": "6a8fa9b1902fab56e0ce50aa", "status": "rejected",
      "tags": null, "interaction": [], "suppression_reason": "hard_bounced",
      "created_at": "2026-08-27T16:48:42.000000Z", "updated_at": "2026-08-27T16:48:42.000000Z",
      "headers": []
    }
  ],
  "links": {"first": "https://api.mailersend.com/v1/emails?page=1", "last": null, "prev": null, "next": null},
  "meta": {"current_page": 1, "current_page_url": "https://api.mailersend.com/v1/emails?page=1",
           "from": 1, "path": "https://api.mailersend.com/v1/emails", "per_page": 10, "to": 3}
}`

// listEmailsEmptyResponse - measured GET v1/emails payload for a page past the end
// of the result set.
const listEmailsEmptyResponse = `{
  "data": [],
  "links": {"first": "https://api.mailersend.com/v1/emails?page=1", "last": null,
            "prev": "https://api.mailersend.com/v1/emails?page=1", "next": null},
  "meta": {"current_page": 2, "current_page_url": "https://api.mailersend.com/v1/emails?page=2",
           "from": null, "path": "https://api.mailersend.com/v1/emails", "per_page": 10, "to": null}
}`

// singleEmailResponse - measured GET v1/email/{email_id} payload.
const singleEmailResponse = `{
  "data": {
    "id": "6a8fa9b1902fab56e0ce50dd", "from": "sender@example.com", "to": "rcpt@example.org",
    "subject": "Welcome", "text": "This is the text content", "html": "<p>This is the HTML content</p>",
    "template_id": null, "domain_id": "7nxe3yjmeq28vp0k",
    "message_id": "6a8fa9b1902fab56e0ce50aa", "status": "sent",
    "tags": ["newsletter"], "interaction": ["opened"], "suppression_reason": null,
    "created_at": "2026-08-27T16:48:42.000000Z", "updated_at": "2026-08-27T16:48:42.000000Z",
    "recipient": {
      "id": "6a8fa9b1902fab56e0ce50bb", "email": "rcpt@example.org",
      "created_at": "2026-08-27T16:48:42.000000Z", "updated_at": "2026-08-27T16:48:42.000000Z"
    },
    "headers": [{"name": "X-Custom", "value": "foo"}],
    "activity": [
      {"id": "6a8fa9b1902fab56e0ce50c1", "type": "opened", "created_at": "2026-08-27T16:49:11.000000Z"},
      {"id": "6a8fa9b1902fab56e0ce50c2", "type": "delivered", "created_at": "2026-08-27T16:48:48.000000Z"},
      {"id": "6a8fa9b1902fab56e0ce50c3", "type": "sent", "created_at": "2026-08-27T16:48:42.000000Z"}
    ]
  }
}`

func jsonResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}
}

func TestCanCreateListEmailOptions(t *testing.T) {
	options := mailersend.ListEmailOptions{
		DomainID: listEmailsDomainID,
		DateFrom: listEmailsDateFrom,
		DateTo:   listEmailsDateTo,
	}

	assert.Equal(t, listEmailsDomainID, options.DomainID)
	assert.Equal(t, listEmailsDateFrom, options.DateFrom)
	assert.Equal(t, listEmailsDateTo, options.DateTo)
	assert.Equal(t, 0, options.Page)
	assert.Equal(t, 0, options.Limit)
	assert.Empty(t, options.Status)
	assert.Empty(t, options.Interaction)
}

func TestCanMockEmailList(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters

		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, fmt.Sprintf(
			"https://api.mailersend.com/v1/emails?date_from=%v&date_to=%v&domain_id=%v&limit=50&page=2",
			listEmailsDateFrom, listEmailsDateTo, listEmailsDomainID,
		), req.URL.String())

		return jsonResponse(listEmailsResponse)
	})

	ctx := context.TODO()

	ms.SetClient(client)

	options := &mailersend.ListEmailOptions{
		DomainID: listEmailsDomainID,
		DateFrom: listEmailsDateFrom,
		DateTo:   listEmailsDateTo,
		Page:     2,
		Limit:    50,
	}

	_, res, err := ms.Email.List(ctx, options)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

// TestEmailListSendsStatusAndInteractionAsArrayParams - the API returns 422 for a
// scalar status or interaction, so both must go out as repeated status[] and
// interaction[] params and never as a single comma-joined value.
func TestEmailListSendsStatusAndInteractionAsArrayParams(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters

		assert.Equal(t, fmt.Sprintf(
			"date_from=%v&date_to=%v&domain_id=%v&interaction%%5B%%5D=opened&interaction%%5B%%5D=clicked&status%%5B%%5D=sent&status%%5B%%5D=delivered",
			listEmailsDateFrom, listEmailsDateTo, listEmailsDomainID,
		), req.URL.RawQuery)

		assert.NotContains(t, req.URL.RawQuery, "sent%2Cdelivered")
		assert.NotContains(t, req.URL.RawQuery, "opened%2Cclicked")

		q := req.URL.Query()
		assert.Equal(t, []string{"sent", "delivered"}, q["status[]"])
		assert.Equal(t, []string{"opened", "clicked"}, q["interaction[]"])
		assert.Empty(t, q["status"])
		assert.Empty(t, q["interaction"])

		return jsonResponse(listEmailsResponse)
	})

	ctx := context.TODO()

	ms.SetClient(client)

	options := &mailersend.ListEmailOptions{
		DomainID:    listEmailsDomainID,
		DateFrom:    listEmailsDateFrom,
		DateTo:      listEmailsDateTo,
		Status:      []string{"sent", "delivered"},
		Interaction: []string{"opened", "clicked"},
	}

	_, _, err := ms.Email.List(ctx, options)

	assert.NoError(t, err)
}

func TestEmailListSendsOptionalFiltersWhenSet(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters

		q := req.URL.Query()

		assert.Equal(t, "rcpt@example.org", q.Get("recipient_email"))
		assert.Equal(t, "6a8fa9b1902fab56e0ce50aa", q.Get("message_id"))
		assert.Equal(t, listEmailsDomainID, q.Get("template_id"))
		assert.Equal(t, "Welcome", q.Get("subject"))
		assert.Equal(t, "newsletter", q.Get("tag"))

		return jsonResponse(listEmailsResponse)
	})

	ctx := context.TODO()

	ms.SetClient(client)

	options := &mailersend.ListEmailOptions{
		DomainID:       listEmailsDomainID,
		DateFrom:       listEmailsDateFrom,
		DateTo:         listEmailsDateTo,
		RecipientEmail: "rcpt@example.org",
		MessageID:      "6a8fa9b1902fab56e0ce50aa",
		TemplateID:     listEmailsDomainID,
		Subject:        "Welcome",
		Tag:            "newsletter",
	}

	_, _, err := ms.Email.List(ctx, options)

	assert.NoError(t, err)
}

func TestEmailListOmitsUnsetOptionalFilters(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters

		assert.Equal(t, fmt.Sprintf(
			"date_from=%v&date_to=%v&domain_id=%v",
			listEmailsDateFrom, listEmailsDateTo, listEmailsDomainID,
		), req.URL.RawQuery)

		for _, param := range []string{
			"page", "limit", "status[]", "interaction[]", "recipient_email",
			"message_id", "template_id", "subject", "tag",
		} {
			assert.NotContains(t, req.URL.Query(), param)
		}

		return jsonResponse(listEmailsResponse)
	})

	ctx := context.TODO()

	ms.SetClient(client)

	options := &mailersend.ListEmailOptions{
		DomainID: listEmailsDomainID,
		DateFrom: listEmailsDateFrom,
		DateTo:   listEmailsDateTo,
	}

	_, _, err := ms.Email.List(ctx, options)

	assert.NoError(t, err)
}

func TestEmailListResponse(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		return jsonResponse(listEmailsResponse)
	})

	ctx := context.TODO()

	ms.SetClient(client)

	root, _, err := ms.Email.List(ctx, &mailersend.ListEmailOptions{
		DomainID: listEmailsDomainID,
		DateFrom: listEmailsDateFrom,
		DateTo:   listEmailsDateTo,
	})

	assert.NoError(t, err)
	assert.Len(t, root.Data, 2)

	first := root.Data[0]

	assert.Equal(t, emailID, first.ID)
	assert.Equal(t, "sender@example.com", first.From)
	assert.Equal(t, "rcpt@example.org", first.To)
	assert.Equal(t, "Welcome", first.Subject)
	assert.Equal(t, listEmailsDomainID, first.DomainID)
	assert.Equal(t, "6a8fa9b1902fab56e0ce50aa", first.MessageID)
	assert.Equal(t, "sent", first.Status)
	assert.Equal(t, []string{"newsletter"}, first.Tags)
	assert.Equal(t, []string{"opened"}, first.Interaction)
	assert.Equal(t, "2026-08-27T16:48:42.000000Z", first.CreatedAt)
	assert.Equal(t, "2026-08-27T16:48:42.000000Z", first.UpdatedAt)

	// List rows never carry content, and template_id / suppression_reason are only
	// set for templated / rejected emails, so all of these arrive as nil pointers.
	assert.Nil(t, first.Text)
	assert.Nil(t, first.HTML)
	assert.Nil(t, first.SuppressionReason)

	assert.NotNil(t, first.TemplateID)
	assert.Equal(t, listEmailsDomainID, *first.TemplateID)

	headers, ok := first.Headers.([]interface{})
	assert.True(t, ok)
	assert.Len(t, headers, 1)

	header, ok := headers[0].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "X-Custom", header["name"])
	assert.Equal(t, "foo", header["value"])

	second := root.Data[1]

	assert.Equal(t, "rejected", second.Status)
	assert.Nil(t, second.TemplateID)
	assert.Nil(t, second.Tags)

	assert.NotNil(t, second.SuppressionReason)
	assert.Equal(t, "hard_bounced", *second.SuppressionReason)

	// An empty JSON array decodes to an empty but non-nil slice, so ranging over
	// Interaction is safe without a nil check.
	assert.NotNil(t, second.Interaction)
	assert.Empty(t, second.Interaction)

	assert.Equal(t, "https://api.mailersend.com/v1/emails?page=1", root.Links.First)
	assert.Equal(t, "https://api.mailersend.com/v1/emails?page=1", root.Meta.CurrentPageURL)
	assert.Equal(t, "https://api.mailersend.com/v1/emails", root.Meta.Path)
	assert.Equal(t, "1", root.Meta.CurrentPage.String())
	assert.Equal(t, "10", root.Meta.PerPage.String())
	assert.Equal(t, "1", root.Meta.From.String())
	assert.Equal(t, "3", root.Meta.To.String())

	// The endpoint reports no last page, and a JSON null decodes to "" on the
	// shared Links struct rather than to a nil pointer, so callers have to page
	// until Links.Next is empty instead of comparing against Links.Last.
	assert.Equal(t, "", root.Links.Last)
	assert.Equal(t, "", root.Links.Next)
	assert.Equal(t, "", root.Links.Prev)
}

func TestEmailListEmptyPageResponse(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		return jsonResponse(listEmailsEmptyResponse)
	})

	ctx := context.TODO()

	ms.SetClient(client)

	root, _, err := ms.Email.List(ctx, &mailersend.ListEmailOptions{
		DomainID: listEmailsDomainID,
		DateFrom: listEmailsDateFrom,
		DateTo:   listEmailsDateTo,
		Page:     2,
	})

	assert.NoError(t, err)
	assert.Empty(t, root.Data)

	assert.Equal(t, "https://api.mailersend.com/v1/emails?page=1", root.Links.Prev)
	assert.Equal(t, "", root.Links.Next)

	assert.Equal(t, "2", root.Meta.CurrentPage.String())
	assert.Equal(t, "https://api.mailersend.com/v1/emails?page=2", root.Meta.CurrentPageURL)

	// meta.from and meta.to are null on an empty page and Meta types them as
	// json.Number, so they decode to an empty string that does not parse as a
	// number. Callers must not call Int64() on them unconditionally.
	assert.Equal(t, "", root.Meta.From.String())
	assert.Equal(t, "", root.Meta.To.String())

	_, err = root.Meta.From.Int64()
	assert.Error(t, err)
}

func TestCanMockEmailGet(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters

		assert.Equal(t, http.MethodGet, req.Method)
		// The single email endpoint is singular, v1/email/{email_id}, unlike the
		// plural v1/emails list endpoint.
		assert.Equal(t, fmt.Sprintf("https://api.mailersend.com/v1/email/%v", emailID), req.URL.String())
		assert.Equal(t, "", req.URL.RawQuery)

		return jsonResponse(singleEmailResponse)
	})

	ctx := context.TODO()

	ms.SetClient(client)

	_, res, err := ms.Email.Get(ctx, emailID)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestEmailGetResponse(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		return jsonResponse(singleEmailResponse)
	})

	ctx := context.TODO()

	ms.SetClient(client)

	root, _, err := ms.Email.Get(ctx, emailID)

	assert.NoError(t, err)

	email := root.Data

	assert.Equal(t, emailID, email.ID)
	assert.Equal(t, "sender@example.com", email.From)
	assert.Equal(t, "rcpt@example.org", email.To)
	assert.Equal(t, "Welcome", email.Subject)
	assert.Equal(t, "sent", email.Status)
	assert.Equal(t, listEmailsDomainID, email.DomainID)
	assert.Equal(t, []string{"newsletter"}, email.Tags)
	assert.Equal(t, []string{"opened"}, email.Interaction)

	assert.NotNil(t, email.Text)
	assert.Equal(t, text, *email.Text)
	assert.NotNil(t, email.HTML)
	assert.Equal(t, html, *email.HTML)

	assert.Nil(t, email.TemplateID)
	assert.Nil(t, email.SuppressionReason)

	assert.Equal(t, "6a8fa9b1902fab56e0ce50bb", email.Recipient.ID)
	assert.Equal(t, "rcpt@example.org", email.Recipient.Email)

	headers, ok := email.Headers.([]interface{})
	assert.True(t, ok)
	assert.Len(t, headers, 1)

	assert.Len(t, email.Activity, 3)

	// Activity is newest first.
	assert.Equal(t, "6a8fa9b1902fab56e0ce50c1", email.Activity[0].ID)
	assert.Equal(t, "opened", email.Activity[0].Type)
	assert.Equal(t, "2026-08-27T16:49:11.000000Z", email.Activity[0].CreatedAt)
	assert.Nil(t, email.Activity[0].SuppressionReason)

	assert.Equal(t, "delivered", email.Activity[1].Type)
	assert.Equal(t, "sent", email.Activity[2].Type)
}
