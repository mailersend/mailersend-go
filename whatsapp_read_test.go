package mailersend_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
)

func jsonResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}
}

func TestWhatsAppMessageService_List(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/messages?limit=25&page=2", req.URL.String())

		return jsonResponse(`{
			"data": [
				{
					"id": "6abf767aa82638483c42dba3",
					"from": "15550001111",
					"to": ["15559001001", "15559001002"],
					"whatsapp_account_id": "qozw59rko0dkxn2p",
					"template_id": null,
					"created_at": "2026-10-02T06:16:42.000000Z"
				}
			],
			"links": {},
			"meta": {"current_page": 2, "per_page": 25, "total": 26}
		}`)
	}))

	root, res, err := ms.WhatsAppMessage.List(context.TODO(), &mailersend.ListWhatsAppMessageOptions{Page: 2, Limit: 25})

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Len(t, root.Data, 1)
	assert.Equal(t, "6abf767aa82638483c42dba3", root.Data[0].ID)
	assert.Equal(t, []string{"15559001001", "15559001002"}, root.Data[0].To)
	assert.Equal(t, "qozw59rko0dkxn2p", *root.Data[0].WhatsAppAccountID)
	assert.Nil(t, root.Data[0].TemplateID)
	assert.Empty(t, root.Data[0].Recipients)
}

func TestWhatsAppMessageService_ListWithoutOptions(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/messages", req.URL.String())

		return jsonResponse(`{"data": [], "links": {}, "meta": {}}`)
	}))

	_, _, err := ms.WhatsAppMessage.List(context.TODO(), nil)

	assert.NoError(t, err)
}

func TestWhatsAppMessageService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/messages/6abf767aa82638483c42dba3", req.URL.String())

		return jsonResponse(`{
			"data": {
				"id": "6abf767aa82638483c42dba3",
				"from": "15550001111",
				"to": ["15559001003", "15559001004"],
				"whatsapp_account_id": "qozw59rko0dkxn2p",
				"template_id": "x9k2wd47mq3l5r8z",
				"created_at": "2026-10-02T06:16:42.000000Z",
				"recipients": [
					{
						"id": "6abf767a5ef12b2317cb8c0e",
						"to": "15559001003",
						"status": "delivered",
						"error_code": null,
						"error_message": null,
						"activity": [
							{"status": "queued", "created_at": "2026-10-02T06:16:42.000000Z"},
							{"status": "sent", "created_at": "2026-10-02T06:16:47.000000Z"},
							{"status": "delivered", "created_at": "2026-10-02T06:17:12.000000Z"}
						]
					},
					{
						"id": "6abf767aa113639674696bbd",
						"to": "15559001004",
						"status": "failed",
						"error_code": 131047,
						"error_message": "Re-engagement message",
						"activity": [
							{"status": "queued", "created_at": "2026-10-02T06:16:42.000000Z"},
							{"status": "failed", "created_at": "2026-10-02T06:16:47.000000Z"}
						]
					}
				]
			}
		}`)
	}))

	root, _, err := ms.WhatsAppMessage.Get(context.TODO(), "6abf767aa82638483c42dba3")

	assert.NoError(t, err)
	assert.Equal(t, "x9k2wd47mq3l5r8z", *root.Data.TemplateID)
	assert.Len(t, root.Data.Recipients, 2)

	delivered := root.Data.Recipients[0]
	assert.Equal(t, "delivered", delivered.Status)
	assert.Nil(t, delivered.ErrorCode)
	assert.Len(t, delivered.Activity, 3)
	assert.Equal(t, "queued", delivered.Activity[0].Status)
	assert.Equal(t, "delivered", delivered.Activity[2].Status)

	failed := root.Data.Recipients[1]
	assert.Equal(t, 131047, *failed.ErrorCode)
	assert.Equal(t, "Re-engagement message", *failed.ErrorMessage)
}

func TestWhatsAppInboundMessageService_ListWithFilters(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t,
			"https://api.mailersend.com/v1/whatsapp/inbound-messages?date_from=1790000000&date_to=1790086400&limit=10&page=1&type%5B%5D=text&type%5B%5D=image&whatsapp_account_id=qozw59rko0dkxn2p",
			req.URL.String(),
		)

		return jsonResponse(`{
			"data": [
				{
					"id": "6abf767a7aea5e17fc8d570b",
					"whatsapp_account_id": "qozw59rko0dkxn2p",
					"from": "15558000001",
					"to": "15550001111",
					"type": "text",
					"received_at": "2026-10-01T09:16:42.000000Z",
					"text": {"body": "Thanks!"},
					"context": {"message_id": "wamid.original"}
				},
				{
					"id": "6abf767a2c345cddbe7814d5",
					"whatsapp_account_id": "qozw59rko0dkxn2p",
					"from": "15558000002",
					"to": "15550001111",
					"type": "image",
					"received_at": "2026-10-01T13:16:42.000000Z",
					"attachment": {
						"url": "https://app.mailersend.com/whatsapp/inbound-message-attachment/6abf767a2c345cddbe7814d5?signature=abc",
						"status": "stored",
						"mime_type": "image/jpeg",
						"size": 2048,
						"caption": "Damaged box"
					}
				}
			],
			"links": {},
			"meta": {"current_page": 1, "per_page": 10, "total": 2}
		}`)
	}))

	root, _, err := ms.WhatsAppInboundMessage.List(context.TODO(), &mailersend.ListWhatsAppInboundMessageOptions{
		WhatsAppAccountID: "qozw59rko0dkxn2p",
		Type:              []string{"text", "image"},
		DateFrom:          1790000000,
		DateTo:            1790086400,
		Page:              1,
		Limit:             10,
	})

	assert.NoError(t, err)
	assert.Len(t, root.Data, 2)

	text := root.Data[0]
	assert.Equal(t, "Thanks!", text.Text.Body)
	assert.Equal(t, "wamid.original", text.Context.MessageID)
	assert.Nil(t, text.Attachment)

	image := root.Data[1]
	assert.Nil(t, image.Text)
	assert.Equal(t, "stored", image.Attachment.Status)
	assert.Equal(t, int64(2048), image.Attachment.Size)
	assert.Equal(t, "Damaged box", image.Attachment.Caption)
	assert.Contains(t, *image.Attachment.URL, "signature=")
}

func TestWhatsAppInboundMessageService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/inbound-messages/6abf767a42f6ca317e9be52f", req.URL.String())

		return jsonResponse(`{
			"data": {
				"id": "6abf767a42f6ca317e9be52f",
				"whatsapp_account_id": "qozw59rko0dkxn2p",
				"from": "15558000003",
				"to": "15550001111",
				"type": "document",
				"received_at": "2026-10-02T07:16:42.000000Z",
				"attachment": {"url": null, "status": "pending", "filename": "invoice.pdf"}
			}
		}`)
	}))

	root, _, err := ms.WhatsAppInboundMessage.Get(context.TODO(), "6abf767a42f6ca317e9be52f")

	assert.NoError(t, err)
	assert.Equal(t, "document", root.Data.Type)
	assert.Equal(t, "pending", root.Data.Attachment.Status)
	assert.Nil(t, root.Data.Attachment.URL)
	assert.Equal(t, "invoice.pdf", root.Data.Attachment.Filename)
	assert.Nil(t, root.Data.Context)
}

func TestWhatsAppRecipientService_ListWithStatus(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/recipients?limit=10&status=blocked", req.URL.String())

		return jsonResponse(`{
			"data": [
				{
					"id": "6abf767a6b269d07b76df9dd",
					"phone_number": "15559990000",
					"username": null,
					"bsuid": null,
					"status": "blocked",
					"created_at": "2026-10-02T09:16:42.000000Z"
				}
			],
			"links": {},
			"meta": {"current_page": 1, "per_page": 10}
		}`)
	}))

	root, _, err := ms.WhatsAppRecipient.List(context.TODO(), &mailersend.ListWhatsAppRecipientOptions{Status: "blocked", Limit: 10})

	assert.NoError(t, err)
	assert.Len(t, root.Data, 1)
	assert.Equal(t, "15559990000", *root.Data[0].PhoneNumber)
	assert.Nil(t, root.Data[0].Username)
	assert.Nil(t, root.Data[0].BSUID)
	assert.Equal(t, "blocked", root.Data[0].Status)
	assert.Empty(t, root.Data[0].Messages)
}

func TestWhatsAppRecipientService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/recipients/6abf767a7265f6c14fd5675c", req.URL.String())

		return jsonResponse(`{
			"data": {
				"id": "6abf767a7265f6c14fd5675c",
				"phone_number": "15559000002",
				"username": null,
				"bsuid": null,
				"status": "active",
				"created_at": "2026-10-02T09:16:42.000000Z",
				"messages": [
					{
						"id": "6abf767aa4d7ee44c8cfd0f8",
						"whatsapp_message_id": "6abf7679a0b3bf7de38ee784",
						"status": "failed",
						"template_name": "order_update",
						"error_code": 131047,
						"error_message": "Re-engagement message",
						"created_at": "2026-10-02T03:16:41.000000Z"
					}
				]
			}
		}`)
	}))

	root, _, err := ms.WhatsAppRecipient.Get(context.TODO(), "6abf767a7265f6c14fd5675c")

	assert.NoError(t, err)
	assert.Len(t, root.Data.Messages, 1)

	message := root.Data.Messages[0]
	assert.Equal(t, "6abf7679a0b3bf7de38ee784", message.WhatsAppMessageID)
	assert.Equal(t, "order_update", message.TemplateName)
	assert.Equal(t, 131047, *message.ErrorCode)
	assert.Equal(t, "Re-engagement message", *message.ErrorMessage)
}

func TestWhatsAppReadServices_ReturnAPIErrors(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       io.NopCloser(bytes.NewBufferString(`{"message": "Resource not found."}`)),
		}
	}))

	_, res, err := ms.WhatsAppMessage.Get(context.TODO(), "missing")

	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
