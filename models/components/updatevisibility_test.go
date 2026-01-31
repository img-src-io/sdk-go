package components

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateVisibilityRequest_Getter(t *testing.T) {
	t.Parallel()

	t.Run("with value", func(t *testing.T) {
		t.Parallel()
		r := UpdateVisibilityRequest{Visibility: VisibilityPublic}
		assert.Equal(t, VisibilityPublic, r.GetVisibility())

		r2 := UpdateVisibilityRequest{Visibility: VisibilityPrivate}
		assert.Equal(t, VisibilityPrivate, r2.GetVisibility())
	})

	t.Run("zero value", func(t *testing.T) {
		t.Parallel()
		r := UpdateVisibilityRequest{}
		assert.Equal(t, Visibility(""), r.GetVisibility())
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var r *UpdateVisibilityRequest
		assert.Equal(t, Visibility(""), r.GetVisibility())
	})
}

func TestUpdateVisibilityRequest_JSON(t *testing.T) {
	t.Parallel()

	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		r := UpdateVisibilityRequest{Visibility: VisibilityPrivate}
		data, err := json.Marshal(r)
		require.NoError(t, err)
		assert.Equal(t, `{"visibility":"private"}`, string(data))
	})

	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		var r UpdateVisibilityRequest
		err := json.Unmarshal([]byte(`{"visibility":"public"}`), &r)
		require.NoError(t, err)
		assert.Equal(t, VisibilityPublic, r.Visibility)
	})

	t.Run("roundtrip", func(t *testing.T) {
		t.Parallel()
		original := UpdateVisibilityRequest{Visibility: VisibilityPrivate}
		data, err := json.Marshal(original)
		require.NoError(t, err)

		var restored UpdateVisibilityRequest
		err = json.Unmarshal(data, &restored)
		require.NoError(t, err)
		assert.Equal(t, original, restored)
	})
}

func TestUpdateVisibilityResponse_Getters(t *testing.T) {
	t.Parallel()

	t.Run("all fields populated", func(t *testing.T) {
		t.Parallel()
		r := UpdateVisibilityResponse{
			ID:         "abc123",
			Visibility: VisibilityPrivate,
			Message:    "Visibility updated",
		}
		assert.Equal(t, "abc123", r.GetID())
		assert.Equal(t, VisibilityPrivate, r.GetVisibility())
		assert.Equal(t, "Visibility updated", r.GetMessage())
	})

	t.Run("zero values", func(t *testing.T) {
		t.Parallel()
		r := UpdateVisibilityResponse{}
		assert.Equal(t, "", r.GetID())
		assert.Equal(t, Visibility(""), r.GetVisibility())
		assert.Equal(t, "", r.GetMessage())
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var r *UpdateVisibilityResponse
		assert.Equal(t, "", r.GetID())
		assert.Equal(t, Visibility(""), r.GetVisibility())
		assert.Equal(t, "", r.GetMessage())
	})
}

func TestUpdateVisibilityResponse_JSON(t *testing.T) {
	t.Parallel()

	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		r := UpdateVisibilityResponse{
			ID:         "img_abc",
			Visibility: VisibilityPublic,
			Message:    "Image is now public",
		}
		data, err := json.Marshal(r)
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)
		assert.Equal(t, "img_abc", result["id"])
		assert.Equal(t, "public", result["visibility"])
		assert.Equal(t, "Image is now public", result["message"])
	})

	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		jsonData := `{"id":"img_xyz","visibility":"private","message":"Image is now private"}`
		var r UpdateVisibilityResponse
		err := json.Unmarshal([]byte(jsonData), &r)
		require.NoError(t, err)
		assert.Equal(t, "img_xyz", r.ID)
		assert.Equal(t, VisibilityPrivate, r.Visibility)
		assert.Equal(t, "Image is now private", r.Message)
	})

	t.Run("roundtrip", func(t *testing.T) {
		t.Parallel()
		original := UpdateVisibilityResponse{
			ID:         "roundtrip_id",
			Visibility: VisibilityPublic,
			Message:    "done",
		}
		data, err := json.Marshal(original)
		require.NoError(t, err)

		var restored UpdateVisibilityResponse
		err = json.Unmarshal(data, &restored)
		require.NoError(t, err)
		assert.Equal(t, original, restored)
	})
}
