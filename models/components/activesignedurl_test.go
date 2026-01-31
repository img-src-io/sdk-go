package components

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActiveSignedUrl_Getters(t *testing.T) {
	t.Parallel()

	t.Run("all fields populated", func(t *testing.T) {
		t.Parallel()
		a := ActiveSignedUrl{
			SignedURL: "https://cdn.img-src.io/signed/abc123?token=xyz",
			ExpiresAt: 1700000000,
		}
		assert.Equal(t, "https://cdn.img-src.io/signed/abc123?token=xyz", a.GetSignedURL())
		assert.Equal(t, int64(1700000000), a.GetExpiresAt())
	})

	t.Run("zero values", func(t *testing.T) {
		t.Parallel()
		a := ActiveSignedUrl{}
		assert.Equal(t, "", a.GetSignedURL())
		assert.Equal(t, int64(0), a.GetExpiresAt())
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var a *ActiveSignedUrl
		assert.Equal(t, "", a.GetSignedURL())
		assert.Equal(t, int64(0), a.GetExpiresAt())
	})
}

func TestActiveSignedUrl_JSONMarshal(t *testing.T) {
	t.Parallel()
	a := ActiveSignedUrl{
		SignedURL: "https://example.com/signed",
		ExpiresAt: 1700000000,
	}
	data, err := json.Marshal(a)
	require.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	require.NoError(t, err)

	assert.Equal(t, "https://example.com/signed", result["signed_url"])
	assert.Equal(t, float64(1700000000), result["expires_at"])
}

func TestActiveSignedUrl_JSONUnmarshal(t *testing.T) {
	t.Parallel()
	jsonData := `{"signed_url":"https://cdn.img-src.io/s/token123","expires_at":1700001000}`

	var a ActiveSignedUrl
	err := json.Unmarshal([]byte(jsonData), &a)
	require.NoError(t, err)

	assert.Equal(t, "https://cdn.img-src.io/s/token123", a.SignedURL)
	assert.Equal(t, int64(1700001000), a.ExpiresAt)
}

func TestActiveSignedUrl_JSONRoundTrip(t *testing.T) {
	t.Parallel()
	original := ActiveSignedUrl{
		SignedURL: "https://cdn.img-src.io/signed/roundtrip",
		ExpiresAt: 1700099999,
	}

	data, err := json.Marshal(original)
	require.NoError(t, err)

	var restored ActiveSignedUrl
	err = json.Unmarshal(data, &restored)
	require.NoError(t, err)

	assert.Equal(t, original, restored)
}
