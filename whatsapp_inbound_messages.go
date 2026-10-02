package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const whatsappInboundMessagesPath = "/whatsapp/inbound-messages"

type WhatsAppInboundMessageService interface {
	List(ctx context.Context, options *ListWhatsAppInboundMessageOptions) (*WhatsAppListInboundMessagesRoot, *Response, error)
	Get(ctx context.Context, whatsAppInboundMessageID string) (*WhatsAppSingleInboundMessageRoot, *Response, error)
}

type whatsAppInboundMessageService struct {
	*service
}

// WhatsAppListInboundMessagesRoot - format of WhatsApp inbound messages list response
type WhatsAppListInboundMessagesRoot struct {
	Data  []WhatsAppInboundMessageData `json:"data"`
	Links Links                        `json:"links"`
	Meta  Meta                         `json:"meta"`
}

// WhatsAppSingleInboundMessageRoot - format of single WhatsApp inbound message response
type WhatsAppSingleInboundMessageRoot struct {
	Data WhatsAppInboundMessageData `json:"data"`
}

// WhatsAppInboundMessageData - a WhatsApp message received by one of your phone numbers.
// Exactly one content field is set, determined by Type.
type WhatsAppInboundMessageData struct {
	ID                string    `json:"id"`
	WhatsAppAccountID string    `json:"whatsapp_account_id"`
	From              string    `json:"from"`
	To                string    `json:"to"`
	Type              string    `json:"type"`
	ReceivedAt        time.Time `json:"received_at"`

	Text        *WhatsAppInboundText       `json:"text,omitempty"`
	Attachment  *WhatsAppInboundAttachment `json:"attachment,omitempty"`
	Location    *WhatsAppInboundLocation   `json:"location,omitempty"`
	Contacts    []map[string]interface{}   `json:"contacts,omitempty"`
	Reaction    *WhatsAppInboundReaction   `json:"reaction,omitempty"`
	Button      *WhatsAppInboundButton     `json:"button,omitempty"`
	ListReply   *WhatsAppInboundListReply  `json:"list_reply,omitempty"`
	Interactive map[string]interface{}     `json:"interactive,omitempty"`
	Order       map[string]interface{}     `json:"order,omitempty"`
	System      map[string]interface{}     `json:"system,omitempty"`
	Unknown     map[string]interface{}     `json:"unknown,omitempty"`
	Unsupported map[string]interface{}     `json:"unsupported,omitempty"`

	Context *WhatsAppInboundContext `json:"context,omitempty"`
}

type WhatsAppInboundText struct {
	Body string `json:"body"`
}

// WhatsAppInboundAttachment - media file of an image, audio, video, document or sticker message
type WhatsAppInboundAttachment struct {
	URL      *string `json:"url"`
	Status   string  `json:"status"`
	MimeType string  `json:"mime_type,omitempty"`
	Size     int64   `json:"size,omitempty"`
	Caption  string  `json:"caption,omitempty"`
	Filename string  `json:"filename,omitempty"`
	Animated *bool   `json:"animated,omitempty"`
}

type WhatsAppInboundLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name,omitempty"`
	Address   string  `json:"address,omitempty"`
}

type WhatsAppInboundReaction struct {
	MessageID string `json:"message_id"`
	Emoji     string `json:"emoji"`
}

type WhatsAppInboundButton struct {
	Text    string `json:"text"`
	Payload string `json:"payload"`
}

type WhatsAppInboundListReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// WhatsAppInboundContext - present when the inbound message is a reply
type WhatsAppInboundContext struct {
	MessageID string `json:"message_id"`
}

// ListWhatsAppInboundMessageOptions - modifies the behavior of WhatsAppInboundMessageService.List method
type ListWhatsAppInboundMessageOptions struct {
	WhatsAppAccountID string   `url:"whatsapp_account_id,omitempty"`
	Type              []string `url:"type[],omitempty"`
	DateFrom          int64    `url:"date_from,omitempty"`
	DateTo            int64    `url:"date_to,omitempty"`
	Page              int      `url:"page,omitempty"`
	Limit             int      `url:"limit,omitempty"`
}

func (s *whatsAppInboundMessageService) List(ctx context.Context, options *ListWhatsAppInboundMessageOptions) (*WhatsAppListInboundMessagesRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, whatsappInboundMessagesPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppListInboundMessagesRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *whatsAppInboundMessageService) Get(ctx context.Context, whatsAppInboundMessageID string) (*WhatsAppSingleInboundMessageRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", whatsappInboundMessagesPath, whatsAppInboundMessageID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(WhatsAppSingleInboundMessageRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
