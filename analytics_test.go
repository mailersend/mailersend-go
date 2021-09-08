package mailersend_test

import (
	"testing"
	"time"

	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateAnalyticsOptions(t *testing.T) {
	from := time.Now().Unix()
	to := time.Now().Add(-24 * time.Hour).Unix()

	domainID := "domain-id"

	options := mailersend.AnalyticsOptions{DomainID: domainID, DateFrom: from, DateTo: to}

	assert.Equal(t, domainID, options.DomainID)
	assert.Equal(t, from, options.DateFrom)
	assert.Equal(t, to, options.DateTo)

}
