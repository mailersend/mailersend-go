package mailersend_test

import (
	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCreateDomainListOptions(t *testing.T) {
	options := mailersend.ListDomainOptions{Page: 0, Limit: 25, Verified: mailersend.Bool(true)}

	assert.Equal(t, 0, options.Page)
	assert.Equal(t, 25, options.Limit)
	assert.Equal(t, mailersend.Bool(true), options.Verified)

}

func TestCanCreateDomainOptions(t *testing.T) {
	options := mailersend.DomainSettingOptions{
		DomainID:                "domain-id",
		SendPaused:              mailersend.Bool(false),
		TrackClicks:             mailersend.Bool(true),
		TrackOpens:              mailersend.Bool(true),
		TrackUnsubscribe:        mailersend.Bool(true),
		TrackUnsubscribeHTML:    "",
		TrackUnsubscribePlain:   "",
		TrackContent:            mailersend.Bool(true),
		CustomTrackingEnabled:   mailersend.Bool(true),
		CustomTrackingSubdomain: "email.mailersend.com",
	}

	options.TrackOpens = mailersend.Bool(false)

	assert.Equal(t, "domain-id", options.DomainID)
	assert.Equal(t, mailersend.Bool(false), options.SendPaused)
	assert.Equal(t, mailersend.Bool(true), options.TrackClicks)
	assert.Equal(t, mailersend.Bool(false), options.TrackOpens)

}