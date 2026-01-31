package operations

import (
	"testing"

	"github.com/img-src-io/sdk-go/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUploadImageRequestBody_VisibilityField(t *testing.T) {
	t.Parallel()

	t.Run("getter with value", func(t *testing.T) {
		t.Parallel()
		vis := components.VisibilityPublic
		r := UploadImageRequestBody{
			Visibility: &vis,
		}
		got := r.GetVisibility()
		require.NotNil(t, got)
		assert.Equal(t, components.VisibilityPublic, *got)
	})

	t.Run("getter with private", func(t *testing.T) {
		t.Parallel()
		vis := components.VisibilityPrivate
		r := UploadImageRequestBody{
			Visibility: &vis,
		}
		got := r.GetVisibility()
		require.NotNil(t, got)
		assert.Equal(t, components.VisibilityPrivate, *got)
	})

	t.Run("getter nil (not set)", func(t *testing.T) {
		t.Parallel()
		r := UploadImageRequestBody{}
		assert.Nil(t, r.GetVisibility())
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var r *UploadImageRequestBody
		assert.Nil(t, r.GetVisibility())
	})

	t.Run("with all fields", func(t *testing.T) {
		t.Parallel()
		vis := components.VisibilityPrivate
		path := "/photos/vacation"
		r := UploadImageRequestBody{
			File: &File{
				FileName: "photo.jpg",
				Content:  []byte("fake-content"),
			},
			TargetPath: &path,
			Visibility: &vis,
		}

		assert.Equal(t, "photo.jpg", r.GetFile().FileName)
		assert.Equal(t, "/photos/vacation", *r.GetTargetPath())
		assert.Equal(t, components.VisibilityPrivate, *r.GetVisibility())
	})
}
