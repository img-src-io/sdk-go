package operations

import (
	"testing"

	"github.com/img-src-io/sdk-go/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateVisibilityRequest_Getters(t *testing.T) {
	t.Parallel()

	t.Run("all fields", func(t *testing.T) {
		t.Parallel()
		r := UpdateVisibilityRequest{
			ID: "img_abc123",
			Body: components.UpdateVisibilityRequest{
				Visibility: components.VisibilityPrivate,
			},
		}
		assert.Equal(t, "img_abc123", r.GetID())
		assert.Equal(t, components.VisibilityPrivate, r.GetBody().Visibility)
	})

	t.Run("zero values", func(t *testing.T) {
		t.Parallel()
		r := UpdateVisibilityRequest{}
		assert.Equal(t, "", r.GetID())
		assert.Equal(t, components.Visibility(""), r.GetBody().Visibility)
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var r *UpdateVisibilityRequest
		assert.Equal(t, "", r.GetID())
		assert.Equal(t, components.UpdateVisibilityRequest{}, r.GetBody())
	})
}

func TestUpdateVisibilityRequest_MarshalJSON(t *testing.T) {
	t.Parallel()
	r := UpdateVisibilityRequest{
		ID: "img_test",
		Body: components.UpdateVisibilityRequest{
			Visibility: components.VisibilityPublic,
		},
	}
	data, err := r.MarshalJSON()
	require.NoError(t, err)
	// The JSON should contain the body fields (pathParam fields are not JSON-serialized)
	assert.Contains(t, string(data), `"visibility"`)
}

func TestUpdateVisibilityRequest_UnmarshalJSON(t *testing.T) {
	t.Parallel()
	jsonData := `{"id":"img_u","body":{"visibility":"private"}}`
	var r UpdateVisibilityRequest
	err := r.UnmarshalJSON([]byte(jsonData))
	require.NoError(t, err)
}

func TestUpdateVisibilityResponse_Getters(t *testing.T) {
	t.Parallel()

	t.Run("with response body", func(t *testing.T) {
		t.Parallel()
		r := UpdateVisibilityResponse{
			UpdateVisibilityResponse: &components.UpdateVisibilityResponse{
				ID:         "img_resp",
				Visibility: components.VisibilityPrivate,
				Message:    "Updated",
			},
		}
		resp := r.GetUpdateVisibilityResponse()
		require.NotNil(t, resp)
		assert.Equal(t, "img_resp", resp.ID)
		assert.Equal(t, components.VisibilityPrivate, resp.Visibility)
		assert.Equal(t, "Updated", resp.Message)
	})

	t.Run("nil response body", func(t *testing.T) {
		t.Parallel()
		r := UpdateVisibilityResponse{}
		assert.Nil(t, r.GetUpdateVisibilityResponse())
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var r *UpdateVisibilityResponse
		assert.Equal(t, components.HTTPMetadata{}, r.GetHTTPMeta())
		assert.Nil(t, r.GetUpdateVisibilityResponse())
	})
}

func TestUpdateVisibilityResponse_MarshalJSON(t *testing.T) {
	t.Parallel()
	r := UpdateVisibilityResponse{
		UpdateVisibilityResponse: &components.UpdateVisibilityResponse{
			ID:         "img_m",
			Visibility: components.VisibilityPublic,
			Message:    "done",
		},
	}
	data, err := r.MarshalJSON()
	require.NoError(t, err)
	assert.NotEmpty(t, data)
}
