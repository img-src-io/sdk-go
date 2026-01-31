package components

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVisibility_Constants(t *testing.T) {
	t.Parallel()
	assert.Equal(t, Visibility("public"), VisibilityPublic)
	assert.Equal(t, Visibility("private"), VisibilityPrivate)
}

func TestVisibility_ToPointer(t *testing.T) {
	t.Parallel()

	t.Run("public", func(t *testing.T) {
		t.Parallel()
		ptr := VisibilityPublic.ToPointer()
		require.NotNil(t, ptr)
		assert.Equal(t, VisibilityPublic, *ptr)
	})

	t.Run("private", func(t *testing.T) {
		t.Parallel()
		ptr := VisibilityPrivate.ToPointer()
		require.NotNil(t, ptr)
		assert.Equal(t, VisibilityPrivate, *ptr)
	})
}

func TestVisibility_IsExact(t *testing.T) {
	t.Parallel()

	t.Run("known values return true", func(t *testing.T) {
		t.Parallel()
		pub := VisibilityPublic
		assert.True(t, pub.IsExact())

		priv := VisibilityPrivate
		assert.True(t, priv.IsExact())
	})

	t.Run("unknown value returns false", func(t *testing.T) {
		t.Parallel()
		unknown := Visibility("unknown")
		assert.False(t, unknown.IsExact())

		empty := Visibility("")
		assert.False(t, empty.IsExact())
	})

	t.Run("nil receiver returns false", func(t *testing.T) {
		t.Parallel()
		var v *Visibility
		assert.False(t, v.IsExact())
	})
}

func TestVisibility_JSONMarshal(t *testing.T) {
	t.Parallel()

	t.Run("public", func(t *testing.T) {
		t.Parallel()
		data, err := json.Marshal(VisibilityPublic)
		require.NoError(t, err)
		assert.Equal(t, `"public"`, string(data))
	})

	t.Run("private", func(t *testing.T) {
		t.Parallel()
		data, err := json.Marshal(VisibilityPrivate)
		require.NoError(t, err)
		assert.Equal(t, `"private"`, string(data))
	})
}

func TestVisibility_JSONUnmarshal(t *testing.T) {
	t.Parallel()

	t.Run("public", func(t *testing.T) {
		t.Parallel()
		var v Visibility
		err := json.Unmarshal([]byte(`"public"`), &v)
		require.NoError(t, err)
		assert.Equal(t, VisibilityPublic, v)
	})

	t.Run("private", func(t *testing.T) {
		t.Parallel()
		var v Visibility
		err := json.Unmarshal([]byte(`"private"`), &v)
		require.NoError(t, err)
		assert.Equal(t, VisibilityPrivate, v)
	})

	t.Run("unknown value", func(t *testing.T) {
		t.Parallel()
		var v Visibility
		err := json.Unmarshal([]byte(`"other"`), &v)
		require.NoError(t, err)
		assert.Equal(t, Visibility("other"), v)
		assert.False(t, v.IsExact())
	})
}

func TestVisibility_JSONRoundTrip(t *testing.T) {
	t.Parallel()
	for _, v := range []Visibility{VisibilityPublic, VisibilityPrivate} {
		v := v
		t.Run(string(v), func(t *testing.T) {
			t.Parallel()
			data, err := json.Marshal(v)
			require.NoError(t, err)

			var got Visibility
			err = json.Unmarshal(data, &got)
			require.NoError(t, err)
			assert.Equal(t, v, got)
		})
	}
}
