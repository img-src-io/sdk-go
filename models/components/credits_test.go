package components

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCredits_Getters(t *testing.T) {
	t.Parallel()

	t.Run("all fields populated", func(t *testing.T) {
		t.Parallel()
		c := Credits{
			StorageBytes:    1024000,
			APIRequests:     500,
			Transformations: 200,
		}
		assert.Equal(t, int64(1024000), c.GetStorageBytes())
		assert.Equal(t, int64(500), c.GetAPIRequests())
		assert.Equal(t, int64(200), c.GetTransformations())
	})

	t.Run("zero values", func(t *testing.T) {
		t.Parallel()
		c := Credits{}
		assert.Equal(t, int64(0), c.GetStorageBytes())
		assert.Equal(t, int64(0), c.GetAPIRequests())
		assert.Equal(t, int64(0), c.GetTransformations())
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var c *Credits
		assert.Equal(t, int64(0), c.GetStorageBytes())
		assert.Equal(t, int64(0), c.GetAPIRequests())
		assert.Equal(t, int64(0), c.GetTransformations())
	})
}

func TestCredits_JSONMarshal(t *testing.T) {
	t.Parallel()
	c := Credits{
		StorageBytes:    5242880,
		APIRequests:     1000,
		Transformations: 350,
	}
	data, err := json.Marshal(c)
	require.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	require.NoError(t, err)

	assert.Equal(t, float64(5242880), result["storage_bytes"])
	assert.Equal(t, float64(1000), result["api_requests"])
	assert.Equal(t, float64(350), result["transformations"])
}

func TestCredits_JSONUnmarshal(t *testing.T) {
	t.Parallel()
	jsonData := `{"storage_bytes":10485760,"api_requests":2000,"transformations":750}`

	var c Credits
	err := json.Unmarshal([]byte(jsonData), &c)
	require.NoError(t, err)

	assert.Equal(t, int64(10485760), c.StorageBytes)
	assert.Equal(t, int64(2000), c.APIRequests)
	assert.Equal(t, int64(750), c.Transformations)
}

func TestCredits_JSONRoundTrip(t *testing.T) {
	t.Parallel()
	original := Credits{
		StorageBytes:    999999,
		APIRequests:     42,
		Transformations: 7,
	}

	data, err := json.Marshal(original)
	require.NoError(t, err)

	var restored Credits
	err = json.Unmarshal(data, &restored)
	require.NoError(t, err)

	assert.Equal(t, original, restored)
}
