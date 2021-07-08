package mailersend_test

import (
	"github.com/mailersend/mailersend-go"
	"testing"
	"time"

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
