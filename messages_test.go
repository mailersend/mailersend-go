package mailersend

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCreateMessagesOptions(t *testing.T) {
	options := ListMessageOptions{
		Page:  1,
		Limit: 25,
	}

	assert.Equal(t, 1, options.Page)
	assert.Equal(t, 25, options.Limit)

}

func TestCodeMessageGetResponse(t *testing.T) {
	response := `{
		"data": {
			"id": "654be07896ceecb1b0548728",
			"created_at": "2023-11-08T19:24:40.000000Z",
			"updated_at": "2023-11-08T19:24:40.000000Z",
			"emails": [
			{
				"id": "654be078bd7e66234bc9288a",
				"from": "johnsmith@domain.com",
				"subject": "subject",
				"text": null,
				"html": null,
				"status": "delivered",
				"tags": null,
				"created_at": "2023-11-08T19:24:40.000000Z",
				"updated_at": "2023-11-08T19:24:48.000000Z"
			}
			],
			"domain": {
				"id": "0p7kx4xxr9749yjr",
				"name": "domain.com",
				"dkim": true,
				"spf": true,
				"tracking": false,
				"is_verified": true,
				"is_cname_verified": false,
				"is_dns_active": true,
				"is_cname_active": false,
				"is_tracking_allowed": false,
				"has_not_queued_messages": false,
				"not_queued_messages_count": 0,
				"domain_settings": {
					"send_paused": false,
					"track_clicks": true,
					"track_opens": true,
					"track_unsubscribe": false,
					"track_unsubscribe_html": "<p>Click here to <a href=\"{$unsubscribe}\">unsubscribe</a></p>",
					"track_unsubscribe_html_enabled": false,
					"track_unsubscribe_plain": "Click here to unsubscribe: {$unsubscribe}",
					"track_unsubscribe_plain_enabled": false,
					"track_content": false,
					"custom_tracking_enabled": false,
					"custom_tracking_subdomain": "email",
					"return_path_subdomain": "mta",
					"inbound_routing_enabled": false,
					"inbound_routing_subdomain": "inbound",
					"precedence_bulk": false,
					"ignore_duplicated_recipients": false
				},
				"created_at": "2023-10-23T00:25:20.000000Z",
				"updated_at": "2023-10-23T00:43:28.000000Z",
				"totals": {
					"sent": 0,
					"delivered": 27,
					"hard_bounced": 0,
					"soft_bounced": 1
				}
			}
		}
	}`

	var root singleMessageRoot
	err := json.NewDecoder(strings.NewReader(response)).Decode(&root)
	if err != nil {
		t.Error(err)
	}
	if root.Data.ID != "654be07896ceecb1b0548728" {
		t.Errorf("root.Data.ID = %s; want 654be07896ceecb1b0548728", root.Data.ID)
	}
	if len(root.Data.Emails) != 1 {
		t.Errorf("len(res.Data.Emails) = %d; want 1", len(root.Data.Emails))
	}
	if root.Data.Emails[0].From != "johnsmith@domain.com" {
		t.Errorf("res.Data.Emails[0].From = %s; want johnsmith@domain.com", root.Data.Emails[0].From)
	}
}
