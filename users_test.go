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

func TestCanCreateUserListOptions(t *testing.T) {
	options := mailersend.ListUserOptions{
		Page:  1,
		Limit: 25,
	}

	assert.Equal(t, 1, options.Page)
	assert.Equal(t, 25, options.Limit)
}

func TestCanCreateInviteUserOptions(t *testing.T) {
	options := mailersend.InviteUserOptions{
		Email: "user@example.com",
		Role:  "admin",
	}

	assert.Equal(t, "user@example.com", options.Email)
	assert.Equal(t, "admin", options.Role)
}

func TestCanCreateUpdateUserOptions(t *testing.T) {
	options := mailersend.UpdateUserOptions{
		Role: "member",
	}

	assert.Equal(t, "member", options.Role)
}

func TestUserService_List(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/users?limit=25&page=1", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": [
					{
						"id": "user-id",
						"email": "user@example.com",
						"name": "User Name",
						"role": "admin",
						"status": "active",
						"created_at": "2023-01-01T00:00:00.000000Z"
					}
				]
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	options := &mailersend.ListUserOptions{
		Page:  1,
		Limit: 25,
	}

	response, _, err := ms.User.List(ctx, options)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, "user-id", response.Data[0].ID)
}

func TestUserService_Get(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/users/user-id", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": {
					"id": "user-id",
					"email": "user@example.com",
					"name": "User Name",
					"role": "admin",
					"status": "active",
					"created_at": "2023-01-01T00:00:00.000000Z"
				}
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	response, _, err := ms.User.Get(ctx, "user-id")

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "user-id", response.Data.ID)
}

func TestUserService_Invite(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/users", req.URL.String())
		assert.Equal(t, http.MethodPost, req.Method)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": {
					"id": "new-user-id",
					"email": "newuser@example.com",
					"role": "member",
					"status": "invited",
					"created_at": "2023-01-01T00:00:00.000000Z"
				}
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	options := &mailersend.InviteUserOptions{
		Email: "newuser@example.com",
		Role:  "member",
	}

	response, _, err := ms.User.Invite(ctx, options)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "new-user-id", response.Data.ID)
}

func TestUserService_Update(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/users/user-id", req.URL.String())
		assert.Equal(t, http.MethodPut, req.Method)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": {
					"id": "user-id",
					"email": "user@example.com",
					"name": "User Name",
					"role": "admin",
					"status": "active",
					"created_at": "2023-01-01T00:00:00.000000Z"
				}
			}`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	options := &mailersend.UpdateUserOptions{
		Role: "admin",
	}

	response, _, err := ms.User.Update(ctx, "user-id", options)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "user-id", response.Data.ID)
	assert.Equal(t, "admin", response.Data.Role)
}

func TestUserService_Delete(t *testing.T) {
	ms := mailersend.NewMailersend(testKey)

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.mailersend.com/v1/users/user-id", req.URL.String())
		assert.Equal(t, http.MethodDelete, req.Method)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()
	ms.SetClient(client)

	_, err := ms.User.Delete(ctx, "user-id")

	assert.NoError(t, err)
}
