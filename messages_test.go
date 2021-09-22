package mailersend_test

import (
	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCreateMessagesOptions(t *testing.T) {
	options := mailersend.ListMessageOptions{
		Page:  1,
		Limit: 25,
	}

	assert.Equal(t, 1, options.Page)
	assert.Equal(t, 25, options.Limit)

}
