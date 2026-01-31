package components

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- UserSettings.Plan ---

func TestUserSettings_PlanField(t *testing.T) {
	t.Parallel()

	t.Run("getter with value", func(t *testing.T) {
		t.Parallel()
		s := UserSettings{Plan: "pro"}
		assert.Equal(t, "pro", s.GetPlan())
	})

	t.Run("getter with empty", func(t *testing.T) {
		t.Parallel()
		s := UserSettings{}
		assert.Equal(t, "", s.GetPlan())
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var s *UserSettings
		assert.Equal(t, "", s.GetPlan())
	})

	t.Run("json marshal includes plan", func(t *testing.T) {
		t.Parallel()
		s := UserSettings{
			ID:       "user_1",
			Username: "alice",
			Plan:     "pro",
		}
		data, err := json.Marshal(s)
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)
		assert.Equal(t, "pro", result["plan"])
	})

	t.Run("json unmarshal reads plan", func(t *testing.T) {
		t.Parallel()
		jsonData := `{"id":"user_1","username":"bob","plan":"free","delivery_formats":[],"default_quality":85,"default_fit_mode":"cover","theme":"dark","language":"en","created_at":1700000000,"updated_at":1700000001,"total_uploads":10,"storage_used_bytes":5000}`
		var s UserSettings
		err := json.Unmarshal([]byte(jsonData), &s)
		require.NoError(t, err)
		assert.Equal(t, "free", s.Plan)
		assert.Equal(t, "bob", s.Username)
	})
}

// --- UsageResponse.Credits ---

func TestUsageResponse_CreditsField(t *testing.T) {
	t.Parallel()

	t.Run("getter with value", func(t *testing.T) {
		t.Parallel()
		u := UsageResponse{
			Credits: Credits{
				StorageBytes:    1024,
				APIRequests:     100,
				Transformations: 50,
			},
		}
		c := u.GetCredits()
		assert.Equal(t, int64(1024), c.StorageBytes)
		assert.Equal(t, int64(100), c.APIRequests)
		assert.Equal(t, int64(50), c.Transformations)
	})

	t.Run("getter with zero", func(t *testing.T) {
		t.Parallel()
		u := UsageResponse{}
		c := u.GetCredits()
		assert.Equal(t, Credits{}, c)
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var u *UsageResponse
		c := u.GetCredits()
		assert.Equal(t, Credits{}, c)
	})

	t.Run("json roundtrip with credits", func(t *testing.T) {
		t.Parallel()
		u := UsageResponse{
			Plan:             "pro",
			PlanName:         "Pro Plan",
			PlanStatus:       PlanStatusActive,
			TotalImages:      42,
			StorageUsedBytes: 5000000,
			StorageUsedMb:    4.77,
			StorageUsedGb:    0.005,
			Credits: Credits{
				StorageBytes:    5000000,
				APIRequests:     999,
				Transformations: 123,
			},
		}
		data, err := json.Marshal(u)
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		credits, ok := result["credits"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, float64(5000000), credits["storage_bytes"])
		assert.Equal(t, float64(999), credits["api_requests"])
		assert.Equal(t, float64(123), credits["transformations"])
	})
}

// --- UploadResponse.Visibility ---

func TestUploadResponse_VisibilityField(t *testing.T) {
	t.Parallel()

	t.Run("getter with value", func(t *testing.T) {
		t.Parallel()
		u := UploadResponse{Visibility: VisibilityPublic}
		assert.Equal(t, VisibilityPublic, u.GetVisibility())
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var u *UploadResponse
		assert.Equal(t, Visibility(""), u.GetVisibility())
	})

	t.Run("json includes visibility", func(t *testing.T) {
		t.Parallel()
		u := UploadResponse{
			ID:         "img_1",
			Hash:       "sha256hash",
			URL:        "https://cdn.img-src.io/alice/photo.webp",
			Paths:      []string{"/alice/photo.webp"},
			Size:       1024,
			Format:     "webp",
			UploadedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			Visibility: VisibilityPrivate,
		}
		data, err := u.MarshalJSON()
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)
		assert.Equal(t, "private", result["visibility"])
	})
}

// --- ImageListItem.Visibility and ActiveSignedUrl ---

func TestImageListItem_NewFields(t *testing.T) {
	t.Parallel()

	t.Run("visibility getter", func(t *testing.T) {
		t.Parallel()
		i := ImageListItem{Visibility: VisibilityPublic}
		assert.Equal(t, VisibilityPublic, i.GetVisibility())
	})

	t.Run("visibility nil receiver", func(t *testing.T) {
		t.Parallel()
		var i *ImageListItem
		assert.Equal(t, Visibility(""), i.GetVisibility())
	})

	t.Run("active_signed_url getter with value", func(t *testing.T) {
		t.Parallel()
		i := ImageListItem{
			ActiveSignedUrl: &ActiveSignedUrl{
				SignedURL: "https://signed.example.com",
				ExpiresAt: 1700000000,
			},
		}
		asu := i.GetActiveSignedUrl()
		require.NotNil(t, asu)
		assert.Equal(t, "https://signed.example.com", asu.SignedURL)
		assert.Equal(t, int64(1700000000), asu.ExpiresAt)
	})

	t.Run("active_signed_url getter nil", func(t *testing.T) {
		t.Parallel()
		i := ImageListItem{}
		assert.Nil(t, i.GetActiveSignedUrl())
	})

	t.Run("active_signed_url nil receiver", func(t *testing.T) {
		t.Parallel()
		var i *ImageListItem
		assert.Nil(t, i.GetActiveSignedUrl())
	})

	t.Run("json marshal with both fields", func(t *testing.T) {
		t.Parallel()
		i := ImageListItem{
			ID:               "img_abc",
			OriginalFilename: "photo.jpg",
			Size:             2048,
			UploadedAt:       time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC),
			URL:              "https://api.img-src.io/v1/images/img_abc",
			Paths:            []string{"/alice/photo.jpg"},
			Visibility:       VisibilityPrivate,
			ActiveSignedUrl: &ActiveSignedUrl{
				SignedURL: "https://cdn.img-src.io/s/token",
				ExpiresAt: 1700001000,
			},
		}
		data, err := i.MarshalJSON()
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		assert.Equal(t, "private", result["visibility"])
		asu, ok := result["active_signed_url"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "https://cdn.img-src.io/s/token", asu["signed_url"])
		assert.Equal(t, float64(1700001000), asu["expires_at"])
	})

	t.Run("json unmarshal with both fields", func(t *testing.T) {
		t.Parallel()
		jsonData := `{
			"id": "img_xyz",
			"original_filename": "test.png",
			"size": 4096,
			"uploaded_at": "2024-06-15T12:00:00Z",
			"url": "https://api.img-src.io/v1/images/img_xyz",
			"paths": ["/bob/test.png"],
			"visibility": "public",
			"active_signed_url": {
				"signed_url": "https://cdn.img-src.io/s/abc",
				"expires_at": 1700002000
			}
		}`
		var i ImageListItem
		err := i.UnmarshalJSON([]byte(jsonData))
		require.NoError(t, err)

		assert.Equal(t, VisibilityPublic, i.Visibility)
		require.NotNil(t, i.ActiveSignedUrl)
		assert.Equal(t, "https://cdn.img-src.io/s/abc", i.ActiveSignedUrl.SignedURL)
		assert.Equal(t, int64(1700002000), i.ActiveSignedUrl.ExpiresAt)
	})

	t.Run("json unmarshal without active_signed_url", func(t *testing.T) {
		t.Parallel()
		jsonData := `{
			"id": "img_no_signed",
			"original_filename": "no_signed.png",
			"size": 1024,
			"uploaded_at": "2024-06-15T12:00:00Z",
			"url": "https://api.img-src.io/v1/images/img_no_signed",
			"paths": ["/alice/no_signed.png"],
			"visibility": "public"
		}`
		var i ImageListItem
		err := i.UnmarshalJSON([]byte(jsonData))
		require.NoError(t, err)

		assert.Equal(t, VisibilityPublic, i.Visibility)
		assert.Nil(t, i.ActiveSignedUrl)
	})
}

// --- SearchResult.Visibility ---

func TestSearchResult_VisibilityField(t *testing.T) {
	t.Parallel()

	t.Run("getter with value", func(t *testing.T) {
		t.Parallel()
		s := SearchResult{Visibility: VisibilityPrivate}
		assert.Equal(t, VisibilityPrivate, s.GetVisibility())
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var s *SearchResult
		assert.Equal(t, Visibility(""), s.GetVisibility())
	})

	t.Run("json includes visibility", func(t *testing.T) {
		t.Parallel()
		s := SearchResult{
			ID:               "img_search",
			OriginalFilename: "found.jpg",
			Paths:            []string{"/user/found.jpg"},
			Size:             512,
			UploadedAt:       time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
			URL:              "https://api.img-src.io/v1/images/img_search",
			Visibility:       VisibilityPublic,
		}
		data, err := s.MarshalJSON()
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)
		assert.Equal(t, "public", result["visibility"])
	})

	t.Run("json unmarshal with visibility", func(t *testing.T) {
		t.Parallel()
		jsonData := `{
			"id": "img_s",
			"original_filename": "s.jpg",
			"paths": ["/u/s.jpg"],
			"size": 100,
			"uploaded_at": "2024-01-01T00:00:00Z",
			"url": "https://api.img-src.io/v1/images/img_s",
			"visibility": "private"
		}`
		var s SearchResult
		err := s.UnmarshalJSON([]byte(jsonData))
		require.NoError(t, err)
		assert.Equal(t, VisibilityPrivate, s.Visibility)
	})
}

// --- MetadataResponse.Visibility ---

func TestMetadataResponse_VisibilityField(t *testing.T) {
	t.Parallel()

	t.Run("getter with value", func(t *testing.T) {
		t.Parallel()
		m := MetadataResponse{Visibility: VisibilityPublic}
		assert.Equal(t, VisibilityPublic, m.GetVisibility())
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var m *MetadataResponse
		assert.Equal(t, Visibility(""), m.GetVisibility())
	})

	t.Run("json includes visibility", func(t *testing.T) {
		t.Parallel()
		m := MetadataResponse{
			ID:         "img_meta",
			Visibility: VisibilityPrivate,
		}
		data, err := json.Marshal(m)
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)
		assert.Equal(t, "private", result["visibility"])
	})

	t.Run("json unmarshal with visibility", func(t *testing.T) {
		t.Parallel()
		jsonData := `{
			"id": "img_m",
			"metadata": {},
			"urls": {},
			"visibility": "public",
			"_links": {}
		}`
		var m MetadataResponse
		err := json.Unmarshal([]byte(jsonData), &m)
		require.NoError(t, err)
		assert.Equal(t, VisibilityPublic, m.Visibility)
	})
}
