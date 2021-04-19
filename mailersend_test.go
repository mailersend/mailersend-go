package mailersend_test

import (
	"net/http"
	"testing"

	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
)

const (
	testKey = "valid-mailersend-api-key"
)

func TestNewMailersend(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	assert.Equal(t, ms.APIKey(), testKey)
	assert.Equal(t, ms.Client(), http.DefaultClient)

	client := new(http.Client)
	ms.SetClient(client)
	assert.Equal(t, client, ms.Client())
}
