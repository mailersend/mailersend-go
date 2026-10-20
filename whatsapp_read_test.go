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

const whatsAppSimplePagination = `
	"links": {
		"first": "https://api.mailersend.com/v1/whatsapp/messages?page=1",
		"last": null,
		"prev": null,
		"next": "https://api.mailersend.com/v1/whatsapp/messages?page=2"
	},
	"meta": {
		"current_page": 1,
		"current_page_url": "https://api.mailersend.com/v1/whatsapp/messages?page=1",
		"from": 1,
		"path": "https://api.mailersend.com/v1/whatsapp/messages",
		"per_page": 25,
		"to": 1
	}`

func TestWhatsAppMessageService_ListWithFilters(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t,
			"https://api.mailersend.com/v1/whatsapp/messages?date_from=1791936000&date_to=1792022400&from=15550001234&limit=25&page=2&template_id=23zxk54v6gjy6v7m&whatsapp_account_id=3enl6x27wmrxrl2v",
			req.URL.String(),
		)

		return jsonResponse(`{
			"data": [
				{
					"id": "67f91abd69f79df391e9d78d",
					"from": "15550001234",
					"to": ["48600000001", "48600000002"],
					"whatsapp_account_id": "3enl6x27wmrxrl2v",
					"template_id": "23zxk54v6gjy6v7m",
					"created_at": "2026-10-14T12:25:33.000000Z"
				},
				{
					"id": "67f91a0c69f79df391e9d6a1",
					"from": null,
					"to": ["US.13491208655302741918"],
					"whatsapp_account_id": null,
					"template_id": null,
					"created_at": "2026-10-14T12:21:16.000000Z"
				}
			],` + whatsAppSimplePagination + `
		}`)
	}))

	root, res, err := ms.WhatsAppMessage.List(context.TODO(), &mailersend.ListWhatsAppMessageOptions{
		WhatsAppAccountID: "3enl6x27wmrxrl2v",
		From:              "15550001234",
		TemplateID:        "23zxk54v6gjy6v7m",
		DateFrom:          1791936000,
		DateTo:            1792022400,
		Page:              2,
		Limit:             25,
	})

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Len(t, root.Data, 2)

	message := root.Data[0]
	assert.Equal(t, "67f91abd69f79df391e9d78d", message.ID)
	assert.Equal(t, "15550001234", *message.From)
	assert.Equal(t, []string{"48600000001", "48600000002"}, message.To)
	assert.Equal(t, "3enl6x27wmrxrl2v", *message.WhatsAppAccountID)
	assert.Equal(t, "23zxk54v6gjy6v7m", *message.TemplateID)
	assert.Equal(t, 2026, message.CreatedAt.Year())
	assert.Empty(t, message.Items)

	virtualNumber := root.Data[1]
	assert.Nil(t, virtualNumber.From)
	assert.Nil(t, virtualNumber.WhatsAppAccountID)
	assert.Nil(t, virtualNumber.TemplateID)

	assert.Equal(t, "", root.Links.Last)
	assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/messages?page=2", root.Links.Next)
	assert.Equal(t, json.Number("1"), root.Meta.CurrentPage)
	assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/messages?page=1", root.Meta.CurrentPageURL)
	assert.Equal(t, json.Number("25"), root.Meta.PerPage)
}

func TestWhatsAppMessageService_ListWithoutOptions(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/messages", req.URL.String())

		return jsonResponse(`{"data": [], "links": {"first": null, "last": null, "prev": null, "next": null}, "meta": {"from": null, "to": null}}`)
	}))

	root, _, err := ms.WhatsAppMessage.List(context.TODO(), nil)

	assert.NoError(t, err)
	assert.Empty(t, root.Data)
}

func TestWhatsAppMessageService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/messages/67f91abd69f79df391e9d78d", req.URL.String())

		return jsonResponse(`{
			"data": {
				"id": "67f91abd69f79df391e9d78d",
				"from": "15550001234",
				"to": ["48600000001", "48600000002"],
				"whatsapp_account_id": "3enl6x27wmrxrl2v",
				"template_id": "23zxk54v6gjy6v7m",
				"created_at": "2026-10-14T12:25:33.000000Z",
				"items": [
					{
						"id": "67f91abe2202f3705540234a",
						"whatsapp_message_id": "67f91abd69f79df391e9d78d",
						"whatsapp_account_id": "3enl6x27wmrxrl2v",
						"from": "15550001234",
						"to": "48600000001",
						"status": "read",
						"template_id": "23zxk54v6gjy6v7m",
						"template_name": "Order confirmation",
						"error_code": null,
						"error_message": null,
						"created_at": "2026-10-14T12:25:33.000000Z"
					},
					{
						"id": "67f91abe2202f3705540234b",
						"whatsapp_message_id": "67f91abd69f79df391e9d78d",
						"whatsapp_account_id": "3enl6x27wmrxrl2v",
						"from": "15550001234",
						"to": "48600000002",
						"status": "failed",
						"template_id": "23zxk54v6gjy6v7m",
						"template_name": "Order confirmation",
						"error_code": 131026,
						"error_message": "Message undeliverable",
						"created_at": "2026-10-14T12:25:33.000000Z"
					}
				]
			}
		}`)
	}))

	root, _, err := ms.WhatsAppMessage.Get(context.TODO(), "67f91abd69f79df391e9d78d")

	assert.NoError(t, err)
	assert.Equal(t, "23zxk54v6gjy6v7m", *root.Data.TemplateID)
	assert.Len(t, root.Data.Items, 2)

	read := root.Data.Items[0]
	assert.Equal(t, "67f91abe2202f3705540234a", read.ID)
	assert.Equal(t, "67f91abd69f79df391e9d78d", read.WhatsAppMessageID)
	assert.Equal(t, "48600000001", read.To)
	assert.Equal(t, "read", read.Status)
	assert.Equal(t, "Order confirmation", read.TemplateName)
	assert.Nil(t, read.ErrorCode)
	assert.Nil(t, read.ErrorMessage)
	assert.Nil(t, read.Recipient)
	assert.Empty(t, read.Activity)

	failed := root.Data.Items[1]
	assert.Equal(t, "failed", failed.Status)
	assert.Equal(t, 131026, *failed.ErrorCode)
	assert.Equal(t, "Message undeliverable", *failed.ErrorMessage)
}

func TestWhatsAppItemService_ListWithFilters(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "/v1/whatsapp/items", req.URL.Path)
		assert.Equal(t,
			"date_from=1791936000&date_to=1792022400&from=15550001234&limit=50&page=1&recipient=48600000002&status%5B%5D=failed&status%5B%5D=read&template_id=23zxk54v6gjy6v7m&whatsapp_account_id=3enl6x27wmrxrl2v&whatsapp_message_id=67f91abd69f79df391e9d78d",
			req.URL.RawQuery,
		)
		assert.Equal(t, []string{"failed", "read"}, req.URL.Query()["status[]"])

		return jsonResponse(`{
			"data": [
				{
					"id": "67f91abe2202f3705540234b",
					"whatsapp_message_id": "67f91abd69f79df391e9d78d",
					"whatsapp_account_id": "3enl6x27wmrxrl2v",
					"from": "15550001234",
					"to": "48600000002",
					"status": "failed",
					"template_id": "23zxk54v6gjy6v7m",
					"template_name": "Order confirmation",
					"error_code": 131026,
					"error_message": "Message undeliverable",
					"created_at": "2026-10-14T12:25:33.000000Z"
				},
				{
					"id": "67f91abe2202f3705540234c",
					"whatsapp_message_id": "67f91a0c69f79df391e9d6a1",
					"whatsapp_account_id": null,
					"from": null,
					"to": "US.13491208655302741918",
					"status": "read",
					"template_id": null,
					"template_name": "Order confirmation",
					"error_code": null,
					"error_message": null,
					"created_at": "2026-10-14T12:21:16.000000Z"
				}
			],` + whatsAppSimplePagination + `
		}`)
	}))

	root, _, err := ms.WhatsAppItem.List(context.TODO(), &mailersend.ListWhatsAppItemOptions{
		Recipient:         "48600000002",
		WhatsAppMessageID: "67f91abd69f79df391e9d78d",
		Status:            []string{"failed", "read"},
		WhatsAppAccountID: "3enl6x27wmrxrl2v",
		From:              "15550001234",
		TemplateID:        "23zxk54v6gjy6v7m",
		DateFrom:          1791936000,
		DateTo:            1792022400,
		Page:              1,
		Limit:             50,
	})

	assert.NoError(t, err)
	assert.Len(t, root.Data, 2)

	failed := root.Data[0]
	assert.Equal(t, "67f91abe2202f3705540234b", failed.ID)
	assert.Equal(t, "67f91abd69f79df391e9d78d", failed.WhatsAppMessageID)
	assert.Equal(t, "3enl6x27wmrxrl2v", *failed.WhatsAppAccountID)
	assert.Equal(t, "15550001234", *failed.From)
	assert.Equal(t, "48600000002", failed.To)
	assert.Equal(t, "failed", failed.Status)
	assert.Equal(t, "23zxk54v6gjy6v7m", *failed.TemplateID)
	assert.Equal(t, "Order confirmation", failed.TemplateName)
	assert.Equal(t, 131026, *failed.ErrorCode)
	assert.Equal(t, "Message undeliverable", *failed.ErrorMessage)
	assert.Equal(t, 2026, failed.CreatedAt.Year())

	virtualNumber := root.Data[1]
	assert.Nil(t, virtualNumber.WhatsAppAccountID)
	assert.Nil(t, virtualNumber.From)
	assert.Nil(t, virtualNumber.TemplateID)
	assert.Nil(t, virtualNumber.ErrorCode)

	assert.Equal(t, "", root.Links.Last)
	assert.Equal(t, json.Number("1"), root.Meta.To)
}

func TestWhatsAppItemService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/items/67f91abe2202f3705540234a", req.URL.String())

		return jsonResponse(`{
			"data": {
				"id": "67f91abe2202f3705540234a",
				"whatsapp_message_id": "67f91abd69f79df391e9d78d",
				"whatsapp_account_id": "3enl6x27wmrxrl2v",
				"from": "15550001234",
				"to": "48600000001",
				"status": "read",
				"template_id": "23zxk54v6gjy6v7m",
				"template_name": "Order confirmation",
				"error_code": null,
				"error_message": null,
				"created_at": "2026-10-14T12:25:33.000000Z",
				"recipient": {
					"id": "67f91abe2202f37055402301",
					"phone_number": "48600000001",
					"username": null,
					"status": "active",
					"created_at": "2026-09-02T08:14:51.000000Z"
				},
				"activity": [
					{"type": "queued", "created_at": "2026-10-14T12:25:33.000000Z"},
					{"type": "sent", "created_at": "2026-10-14T12:25:34.000000Z"},
					{"type": "delivered", "created_at": "2026-10-14T12:25:36.000000Z"},
					{"type": "read", "created_at": "2026-10-14T12:31:02.000000Z"}
				]
			}
		}`)
	}))

	root, _, err := ms.WhatsAppItem.Get(context.TODO(), "67f91abe2202f3705540234a")

	assert.NoError(t, err)
	assert.Equal(t, "read", root.Data.Status)

	recipient := root.Data.Recipient
	assert.NotNil(t, recipient)
	assert.Equal(t, "67f91abe2202f37055402301", recipient.ID)
	assert.Equal(t, "48600000001", *recipient.PhoneNumber)
	assert.Nil(t, recipient.Username)
	assert.Equal(t, "active", recipient.Status)
	assert.Equal(t, 9, int(recipient.CreatedAt.Month()))

	assert.Len(t, root.Data.Activity, 4)
	assert.Equal(t, "queued", root.Data.Activity[0].Type)
	assert.Equal(t, "read", root.Data.Activity[3].Type)
	assert.Equal(t, 31, root.Data.Activity[3].CreatedAt.Minute())
}

func TestWhatsAppItemService_GetWithoutRecipient(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		return jsonResponse(`{
			"data": {
				"id": "67f91abe2202f3705540234a",
				"whatsapp_message_id": "67f91abd69f79df391e9d78d",
				"to": "48600000001",
				"status": "queued",
				"recipient": null,
				"activity": []
			}
		}`)
	}))

	root, _, err := ms.WhatsAppItem.Get(context.TODO(), "67f91abe2202f3705540234a")

	assert.NoError(t, err)
	assert.Nil(t, root.Data.Recipient)
	assert.Empty(t, root.Data.Activity)
}

func TestWhatsAppActivityService_ListWithFilters(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "/v1/whatsapp/activities", req.URL.Path)
		assert.Equal(t,
			"date_from=1791936000&date_to=1792022400&limit=10&page=1&recipient=48600000001&type%5B%5D=delivered&type%5B%5D=read&whatsapp_item_id=67f91abe2202f3705540234a&whatsapp_message_id=67f91abd69f79df391e9d78d",
			req.URL.RawQuery,
		)
		assert.Equal(t, []string{"delivered", "read"}, req.URL.Query()["type[]"])

		return jsonResponse(`{
			"data": [
				{
					"id": "67f91d2e2202f37055402401",
					"whatsapp_item_id": "67f91abe2202f3705540234a",
					"whatsapp_message_id": "67f91abd69f79df391e9d78d",
					"whatsapp_account_id": "3enl6x27wmrxrl2v",
					"from": "15550001234",
					"to": "48600000001",
					"type": "read",
					"created_at": "2026-10-14T12:31:02.000000Z"
				},
				{
					"id": "67f91ac02202f37055402399",
					"whatsapp_item_id": "67f91abe2202f3705540234a",
					"whatsapp_message_id": "67f91abd69f79df391e9d78d",
					"whatsapp_account_id": null,
					"from": null,
					"to": "48600000001",
					"type": "delivered",
					"created_at": "2026-10-14T12:25:36.000000Z"
				}
			],` + whatsAppSimplePagination + `
		}`)
	}))

	root, _, err := ms.WhatsAppActivity.List(context.TODO(), &mailersend.ListWhatsAppActivityOptions{
		WhatsAppItemID:    "67f91abe2202f3705540234a",
		WhatsAppMessageID: "67f91abd69f79df391e9d78d",
		Recipient:         "48600000001",
		Type:              []string{"delivered", "read"},
		DateFrom:          1791936000,
		DateTo:            1792022400,
		Page:              1,
		Limit:             10,
	})

	assert.NoError(t, err)
	assert.Len(t, root.Data, 2)

	read := root.Data[0]
	assert.Equal(t, "67f91d2e2202f37055402401", read.ID)
	assert.Equal(t, "67f91abe2202f3705540234a", read.WhatsAppItemID)
	assert.Equal(t, "67f91abd69f79df391e9d78d", read.WhatsAppMessageID)
	assert.Equal(t, "3enl6x27wmrxrl2v", *read.WhatsAppAccountID)
	assert.Equal(t, "15550001234", *read.From)
	assert.Equal(t, "48600000001", read.To)
	assert.Equal(t, "read", read.Type)
	assert.Equal(t, 31, read.CreatedAt.Minute())

	assert.Nil(t, root.Data[1].WhatsAppAccountID)
	assert.Nil(t, root.Data[1].From)
}

func TestWhatsAppActivityService_ListWithoutOptions(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/activities", req.URL.String())

		return jsonResponse(`{"data": [], "links": {}, "meta": {}}`)
	}))

	_, _, err := ms.WhatsAppActivity.List(context.TODO(), nil)

	assert.NoError(t, err)
}

func TestWhatsAppInboundMessageService_ListWithFilters(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t,
			"https://api.mailersend.com/v1/whatsapp/inbound-messages?date_from=1791936000&date_to=1792022400&from=US.13491208655302741918&limit=10&page=1&type%5B%5D=text&type%5B%5D=image&whatsapp_account_id=3enl6x27wmrxrl2v",
			req.URL.String(),
		)

		return jsonResponse(`{
			"data": [
				{
					"id": "62f114f7165fe0d8db0288e2",
					"whatsapp_account_id": "3enl6x27wmrxrl2v",
					"from": "48600000001",
					"to": "15550001234",
					"type": "text",
					"received_at": "2026-10-14T12:41:19.000000Z",
					"meta": {"body": "Hello, I would like to check my order status."},
					"context": {"message_id": "wamid.HBgLMzcwNjMzMjI3OTcVAgARGBIwREE3N0EwN0EyNjhENjEwOEMA"}
				},
				{
					"id": "62f114c1165fe0d8db0288d9",
					"whatsapp_account_id": "3enl6x27wmrxrl2v",
					"from": "US.13491208655302741918",
					"to": "15550001234",
					"type": "image",
					"received_at": "2026-10-14T12:38:03.000000Z",
					"meta": {
						"url": "https://app.mailersend.com/whatsapp/inbound-message-attachment/62f114c1165fe0d8db0288d9?signature=abc",
						"status": "stored",
						"mime_type": "image/jpeg",
						"size": 102400,
						"caption": "Here is the damaged item"
					}
				}
			],` + whatsAppSimplePagination + `
		}`)
	}))

	root, _, err := ms.WhatsAppInboundMessage.List(context.TODO(), &mailersend.ListWhatsAppInboundMessageOptions{
		WhatsAppAccountID: "3enl6x27wmrxrl2v",
		From:              "US.13491208655302741918",
		Type:              []string{"text", "image"},
		DateFrom:          1791936000,
		DateTo:            1792022400,
		Page:              1,
		Limit:             10,
	})

	assert.NoError(t, err)
	assert.Len(t, root.Data, 2)

	text := root.Data[0]
	assert.Equal(t, "62f114f7165fe0d8db0288e2", text.ID)
	assert.Equal(t, "3enl6x27wmrxrl2v", text.WhatsAppAccountID)
	assert.Equal(t, "48600000001", text.From)
	assert.Equal(t, "15550001234", text.To)
	assert.Equal(t, "text", text.Type)
	assert.Equal(t, 41, text.ReceivedAt.Minute())
	assert.Equal(t, "Hello, I would like to check my order status.", text.Meta["body"])
	assert.Equal(t, "wamid.HBgLMzcwNjMzMjI3OTcVAgARGBIwREE3N0EwN0EyNjhENjEwOEMA", text.Context.MessageID)

	image := root.Data[1]
	assert.Nil(t, image.Context)
	assert.Equal(t, "stored", image.Meta["status"])
	assert.Equal(t, "image/jpeg", image.Meta["mime_type"])
	assert.Equal(t, float64(102400), image.Meta["size"])
	assert.Contains(t, image.Meta["url"], "signature=")
}

func TestWhatsAppInboundMessageService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/inbound-messages/62f114c1165fe0d8db0288f0", req.URL.String())

		return jsonResponse(`{
			"data": {
				"id": "62f114c1165fe0d8db0288f0",
				"whatsapp_account_id": "3enl6x27wmrxrl2v",
				"from": "48600000003",
				"to": "15550001234",
				"type": "unsupported",
				"received_at": "2026-10-14T12:50:00.000000Z",
				"meta": null
			}
		}`)
	}))

	root, _, err := ms.WhatsAppInboundMessage.Get(context.TODO(), "62f114c1165fe0d8db0288f0")

	assert.NoError(t, err)
	assert.Equal(t, "unsupported", root.Data.Type)
	assert.Nil(t, root.Data.Meta)
	assert.Nil(t, root.Data.Context)
}

func TestWhatsAppRecipientService_ListWithFilters(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/recipients?limit=10&page=1&phone_number=48600000001&status=active", req.URL.String())

		return jsonResponse(`{
			"data": [
				{
					"id": "67f91abe2202f37055402301",
					"phone_number": "48600000001",
					"username": null,
					"status": "active",
					"created_at": "2026-09-02T08:14:51.000000Z"
				},
				{
					"id": "67f91abe2202f37055402302",
					"phone_number": null,
					"username": "acme.support",
					"status": "active",
					"created_at": "2026-09-03T08:14:51.000000Z"
				}
			],` + whatsAppSimplePagination + `
		}`)
	}))

	root, _, err := ms.WhatsAppRecipient.List(context.TODO(), &mailersend.ListWhatsAppRecipientOptions{
		Status:      "active",
		PhoneNumber: "48600000001",
		Page:        1,
		Limit:       10,
	})

	assert.NoError(t, err)
	assert.Len(t, root.Data, 2)
	assert.Equal(t, "67f91abe2202f37055402301", root.Data[0].ID)
	assert.Equal(t, "48600000001", *root.Data[0].PhoneNumber)
	assert.Nil(t, root.Data[0].Username)
	assert.Equal(t, "active", root.Data[0].Status)
	assert.Empty(t, root.Data[0].Items)
	assert.Nil(t, root.Data[1].PhoneNumber)
	assert.Equal(t, "acme.support", *root.Data[1].Username)
}

func TestWhatsAppRecipientService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/recipients/67f91abe2202f37055402301", req.URL.String())

		return jsonResponse(`{
			"data": {
				"id": "67f91abe2202f37055402301",
				"phone_number": "48600000001",
				"username": null,
				"status": "active",
				"created_at": "2026-09-02T08:14:51.000000Z",
				"items": [
					{
						"id": "67f91abe2202f3705540234a",
						"whatsapp_message_id": "67f91abd69f79df391e9d78d",
						"whatsapp_account_id": "3enl6x27wmrxrl2v",
						"from": "15550001234",
						"to": "48600000001",
						"status": "failed",
						"template_id": "23zxk54v6gjy6v7m",
						"template_name": "Order confirmation",
						"error_code": 131047,
						"error_message": "Re-engagement message",
						"created_at": "2026-10-14T12:25:33.000000Z"
					}
				]
			}
		}`)
	}))

	root, _, err := ms.WhatsAppRecipient.Get(context.TODO(), "67f91abe2202f37055402301")

	assert.NoError(t, err)
	assert.Len(t, root.Data.Items, 1)

	item := root.Data.Items[0]
	assert.Equal(t, "67f91abe2202f3705540234a", item.ID)
	assert.Equal(t, "67f91abd69f79df391e9d78d", item.WhatsAppMessageID)
	assert.Equal(t, "48600000001", item.To)
	assert.Equal(t, "Order confirmation", item.TemplateName)
	assert.Equal(t, 131047, *item.ErrorCode)
	assert.Equal(t, "Re-engagement message", *item.ErrorMessage)
}

func TestWhatsAppTemplateService_ListWithFilters(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t,
			"https://api.mailersend.com/v1/whatsapp/templates?category=UTILITY&language=en_US&limit=25&page=1&status=approved&whatsapp_account_id=3enl6x27wmrxrl2v",
			req.URL.String(),
		)

		return jsonResponse(`{
			"data": [
				{
					"id": "23zxk54v6gjy6v7m",
					"whatsapp_account_id": "3enl6x27wmrxrl2v",
					"name": "order_confirmation",
					"readable_name": "Order confirmation",
					"category": "UTILITY",
					"language": "en_US",
					"status": "approved",
					"components": {
						"header": {"text": "{{1}}"},
						"body": {"text": "Hi {{1}}, your order is confirmed and will arrive on {{2}}."},
						"buttons": [
							{"type": "URL", "text": "View order", "url": "https://shop.example.com/{{1}}"}
						]
					},
					"created_at": "2026-10-14T12:25:33.000000Z",
					"updated_at": "2026-10-14T12:41:07.000000Z"
				},
				{
					"id": "p6nq4jyzkyv5x3zl",
					"whatsapp_account_id": "3enl6x27wmrxrl2v",
					"name": "delivery_update",
					"readable_name": "Delivery update",
					"category": "UTILITY",
					"language": "en_US",
					"status": "approved",
					"components": {
						"header": {"text": null},
						"body": null,
						"buttons": [
							{"type": "QUICK_REPLY", "text": "Change delivery time", "url": null}
						]
					},
					"created_at": "2026-10-13T09:02:18.000000Z",
					"updated_at": "2026-10-13T09:02:18.000000Z"
				}
			],` + whatsAppSimplePagination + `
		}`)
	}))

	root, _, err := ms.WhatsAppTemplate.List(context.TODO(), &mailersend.ListWhatsAppTemplateOptions{
		WhatsAppAccountID: "3enl6x27wmrxrl2v",
		Status:            "approved",
		Category:          "UTILITY",
		Language:          "en_US",
		Page:              1,
		Limit:             25,
	})

	assert.NoError(t, err)
	assert.Len(t, root.Data, 2)

	template := root.Data[0]
	assert.Equal(t, "23zxk54v6gjy6v7m", template.ID)
	assert.Equal(t, "3enl6x27wmrxrl2v", template.WhatsAppAccountID)
	assert.Equal(t, "order_confirmation", template.Name)
	assert.Equal(t, "Order confirmation", template.ReadableName)
	assert.Equal(t, "UTILITY", template.Category)
	assert.Equal(t, "en_US", template.Language)
	assert.Equal(t, "approved", template.Status)
	assert.Equal(t, "{{1}}", *template.Components.Header.Text)
	assert.Equal(t, "Hi {{1}}, your order is confirmed and will arrive on {{2}}.", *template.Components.Body.Text)
	assert.Len(t, template.Components.Buttons, 1)
	assert.Equal(t, "URL", template.Components.Buttons[0].Type)
	assert.Equal(t, "View order", template.Components.Buttons[0].Text)
	assert.Equal(t, "https://shop.example.com/{{1}}", *template.Components.Buttons[0].URL)
	assert.Equal(t, 25, template.CreatedAt.Minute())
	assert.Equal(t, 41, template.UpdatedAt.Minute())

	mediaHeader := root.Data[1].Components
	assert.NotNil(t, mediaHeader.Header)
	assert.Nil(t, mediaHeader.Header.Text)
	assert.Nil(t, mediaHeader.Body)
	assert.Equal(t, "QUICK_REPLY", mediaHeader.Buttons[0].Type)
	assert.Nil(t, mediaHeader.Buttons[0].URL)
}

func TestWhatsAppTemplateService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	ms.SetClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://api.mailersend.com/v1/whatsapp/templates/23zxk54v6gjy6v7m", req.URL.String())

		return jsonResponse(`{
			"data": {
				"id": "23zxk54v6gjy6v7m",
				"whatsapp_account_id": "3enl6x27wmrxrl2v",
				"name": "order_confirmation",
				"readable_name": "Order confirmation",
				"category": "MARKETING",
				"language": "en_US",
				"status": "paused",
				"components": {"header": null, "body": {"text": "Hi {{1}}"}, "buttons": []},
				"created_at": "2026-10-14T12:25:33.000000Z",
				"updated_at": "2026-10-14T12:41:07.000000Z"
			}
		}`)
	}))

	root, _, err := ms.WhatsAppTemplate.Get(context.TODO(), "23zxk54v6gjy6v7m")

	assert.NoError(t, err)
	assert.Equal(t, "MARKETING", root.Data.Category)
	assert.Equal(t, "paused", root.Data.Status)
	assert.Nil(t, root.Data.Components.Header)
	assert.Equal(t, "Hi {{1}}", *root.Data.Components.Body.Text)
	assert.Empty(t, root.Data.Components.Buttons)
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

	_, res, err = ms.WhatsAppItem.Get(context.TODO(), "missing")
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	_, res, err = ms.WhatsAppTemplate.Get(context.TODO(), "missing")
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
