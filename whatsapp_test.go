package mailersend_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
)

const (
	whatsAppFrom       = "15550001234"
	whatsAppTo         = "48600000001"
	whatsAppTemplateID = "23zxk54v6gjy6v7m"
)

func basicWhatsAppMessage() *mailersend.WhatsAppMessage {
	ms := mailersend.NewMailersend(testKey)

	message := ms.WhatsApp.NewMessage()
	message.SetFrom(whatsAppFrom)
	message.SetTo([]string{whatsAppTo})
	message.SetTemplateID(whatsAppTemplateID)

	return message
}

func TestNewWhatsAppMessage(t *testing.T) {
	message := basicWhatsAppMessage()

	assert.Equal(t, whatsAppFrom, message.From)
	assert.Equal(t, []string{whatsAppTo}, message.To)
	assert.Equal(t, whatsAppTemplateID, message.TemplateID)
	assert.Empty(t, message.Personalization)
}

func TestWhatsAppPersonalization(t *testing.T) {
	message := basicWhatsAppMessage()

	personalization := []mailersend.WhatsAppPersonalization{
		{
			To: whatsAppTo,
			Data: mailersend.WhatsAppPersonalizationData{
				Header:  []string{"Order #12345"},
				Body:    []string{"John", "December 31, 2026"},
				Buttons: []string{"orders/12345"},
			},
		},
	}

	message.SetPersonalization(personalization)

	assert.Equal(t, personalization, message.Personalization)
}

func TestWhatsAppMessageMarshal(t *testing.T) {
	message := basicWhatsAppMessage()

	message.SetPersonalization([]mailersend.WhatsAppPersonalization{
		{
			To: whatsAppTo,
			Data: mailersend.WhatsAppPersonalizationData{
				Body: []string{"John"},
			},
		},
	})

	body, err := json.Marshal(message)
	assert.NoError(t, err)
	assert.JSONEq(t, `{
		"from": "15550001234",
		"to": ["48600000001"],
		"template_id": "23zxk54v6gjy6v7m",
		"personalization": [
			{
				"to": "48600000001",
				"data": {
					"body": ["John"]
				}
			}
		]
	}`, string(body))
}

func TestWhatsAppPersonalizationOmittedWhenEmpty(t *testing.T) {
	message := basicWhatsAppMessage()

	body, err := json.Marshal(message)
	assert.NoError(t, err)
	assert.NotContains(t, string(body), "personalization")
}

func TestWhatsAppService_Send(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/send", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusAccepted,
			Header: http.Header{
				"X-Message-Id": []string{"67f91abd69f79df391e9d78d"},
			},
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": {
					"id": "67f91abd69f79df391e9d78d",
					"created_at": "2026-08-31 07:24:20"
				}
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	res, err := ms.WhatsApp.Send(ctx, basicWhatsAppMessage())

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusAccepted, res.StatusCode)
	assert.Equal(t, "67f91abd69f79df391e9d78d", res.Header.Get("X-Message-Id"))
}
