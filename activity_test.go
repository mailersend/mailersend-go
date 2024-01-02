package mailersend_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateActivityOptions(t *testing.T) {
	from := time.Now().Unix()
	to := time.Now().Add(-24 * time.Hour).Unix()

	options := mailersend.ActivityOptions{DateFrom: from, DateTo: to}

	assert.Equal(t, 0, options.Limit)
	assert.Equal(t, from, options.DateFrom)
	assert.Equal(t, to, options.DateTo)

}

func TestCanMockActivity(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	from := time.Now().Unix()
	to := time.Now().Add(-24 * time.Hour).Unix()

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters

		assert.Equal(t, req.URL.String(), fmt.Sprintf("https://api.mailersend.com/v1/activity/domain-id?date_from=%v&date_to=%v", from, to))
		return &http.Response{
			StatusCode: http.StatusAccepted,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	ms.SetClient(client)

	options := &mailersend.ActivityOptions{DomainID: "domain-id", DateFrom: from, DateTo: to}

	_, _, _ = ms.Activity.List(ctx, options)

	assert.Equal(t, 0, options.Limit)
	assert.Equal(t, from, options.DateFrom)
	assert.Equal(t, to, options.DateTo)

}
