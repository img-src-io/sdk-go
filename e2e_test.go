package sdkgo_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"testing"
	"time"

	sdkgo "github.com/img-src-io/sdk-go"
	"github.com/img-src-io/sdk-go/models/apierrors"
	"github.com/img-src-io/sdk-go/models/components"
	"github.com/img-src-io/sdk-go/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2EIntegration(t *testing.T) {
	apiKey := os.Getenv("IMGSRC_API_KEY")
	serverURL := os.Getenv("IMGSRC_SERVER_URL")
	if apiKey == "" || serverURL == "" {
		t.Skip("IMGSRC_API_KEY or IMGSRC_SERVER_URL not set, skipping E2E test")
	}

	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(apiKey),
		sdkgo.WithServerURL(serverURL),
	)

	// Track uploaded image ID for cleanup
	var uploadedImageID string
	defer func() {
		if uploadedImageID != "" {
			t.Logf("Cleanup: deleting image %s", uploadedImageID)
			_, err := s.Images.Delete(ctx, uploadedImageID)
			if err != nil {
				t.Logf("Cleanup warning: failed to delete image: %v", err)
			}
		}
	}()

	// Generate unique 8x8 PNG
	img := image.NewRGBA(image.Rect(0, 0, 8, 8))
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(rng.Intn(256)),
				G: uint8(rng.Intn(256)),
				B: uint8(rng.Intn(256)),
				A: 255,
			})
		}
	}
	var pngBuf bytes.Buffer
	require.NoError(t, png.Encode(&pngBuf, img))
	pngData := pngBuf.Bytes()

	tmpFile, err := os.CreateTemp("", "e2e-test-*.png")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write(pngData)
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	// ---- 1. GetSettings ----
	t.Run("GetSettings", func(t *testing.T) {
		res, err := s.Settings.Get(ctx)
		require.NoError(t, err)
		require.NotNil(t, res.SettingsResponse)

		settings := res.SettingsResponse.Settings
		t.Logf("Username: %s, Plan: %s", settings.Username, settings.Plan)
		assert.NotEmpty(t, settings.Username)
		assert.NotEmpty(t, settings.Plan)
	})

	// ---- 2. GetUsage ----
	t.Run("GetUsage", func(t *testing.T) {
		res, err := s.Usage.Get(ctx)
		require.NoError(t, err)
		require.NotNil(t, res.UsageResponse)

		usage := res.UsageResponse
		t.Logf("Plan: %s, TotalImages: %d", usage.Plan, usage.TotalImages)
		assert.NotEmpty(t, usage.Plan)
		assert.GreaterOrEqual(t, usage.TotalImages, int64(0))
	})

	// ---- 3. UploadImage ----
	t.Run("UploadImage", func(t *testing.T) {
		targetPath := sdkgo.String("__sdk_e2e_test/go/test-image.png")
		visibility := components.VisibilityPublic

		res, err := s.Images.Upload(ctx, &operations.UploadImageRequestBody{
			File: &operations.File{
				FileName: "test-image.png",
				Content:  pngData,
			},
			TargetPath: targetPath,
			Visibility: &visibility,
		})
		require.NoError(t, err)
		require.NotNil(t, res.UploadResponse)

		uploadedImageID = res.UploadResponse.ID
		t.Logf("Uploaded image ID: %s, URL: %s", res.UploadResponse.ID, res.UploadResponse.URL)
		assert.NotEmpty(t, res.UploadResponse.ID)
		assert.NotEmpty(t, res.UploadResponse.URL)
		assert.Greater(t, res.UploadResponse.Size, int64(0))
		assert.Equal(t, components.VisibilityPublic, res.UploadResponse.Visibility)
	})

	require.NotEmpty(t, uploadedImageID, "Upload must succeed before continuing")

	// ---- 4. ListImages (with retry) ----
	t.Run("ListImages", func(t *testing.T) {
		var found bool
		path := sdkgo.String("__sdk_e2e_test/go/test-image.png")
		for attempt := 0; attempt < 5; attempt++ {
			res, err := s.Images.List(ctx, sdkgo.Int64(50), nil, path)
			require.NoError(t, err)
			require.NotNil(t, res.ImageListResponse)

			for _, img := range res.ImageListResponse.Images {
				if img.ID == uploadedImageID {
					found = true
					t.Logf("Found image in list: %s", img.ID)
					break
				}
			}
			if found {
				break
			}
			// Also try without path filter
			res2, err := s.Images.List(ctx, sdkgo.Int64(50), nil, nil)
			require.NoError(t, err)
			if res2.ImageListResponse != nil {
				for _, img := range res2.ImageListResponse.Images {
					if img.ID == uploadedImageID {
						found = true
						t.Logf("Found image in unfiltered list: %s", img.ID)
						break
					}
				}
			}
			if found {
				break
			}
			t.Logf("Attempt %d: image not yet in list, retrying...", attempt+1)
			time.Sleep(2 * time.Second)
		}
		assert.True(t, found, "Uploaded image should appear in list")
	})

	// ---- 5. SearchImages (with retry) ----
	t.Run("SearchImages", func(t *testing.T) {
		var found bool
		for attempt := 0; attempt < 5; attempt++ {
			res, err := s.Images.Search(ctx, "test-image", sdkgo.Int64(10))
			require.NoError(t, err)
			require.NotNil(t, res.SearchResponse)

			for _, result := range res.SearchResponse.Results {
				if result.ID == uploadedImageID {
					found = true
					t.Logf("Found image in search: %s", result.ID)
					break
				}
			}
			if found {
				break
			}
			t.Logf("Attempt %d: image not yet in search results, retrying...", attempt+1)
			time.Sleep(2 * time.Second)
		}
		assert.True(t, found, "Uploaded image should appear in search results")
	})

	// ---- 6. GetImage (metadata) ----
	t.Run("GetImage", func(t *testing.T) {
		res, err := s.Images.GetMetadata(ctx, uploadedImageID)
		require.NoError(t, err)
		require.NotNil(t, res.MetadataResponse)

		meta := res.MetadataResponse
		t.Logf("Image metadata: ID=%s, Visibility=%s", meta.ID, meta.Visibility)
		assert.Equal(t, uploadedImageID, meta.ID)
		assert.Equal(t, components.VisibilityPublic, meta.Visibility)
	})

	// ---- 7. UpdateVisibility ----
	t.Run("UpdateVisibility", func(t *testing.T) {
		// Set to private
		res, err := s.Images.UpdateVisibility(ctx, uploadedImageID, components.UpdateVisibilityRequest{
			Visibility: components.VisibilityPrivate,
		})
		if err != nil {
			if isHTTPStatus(err, 403) {
				t.Skip("Private visibility requires Pro plan, skipping")
			}
			require.NoError(t, err)
		}
		require.NotNil(t, res.UpdateVisibilityResponse)
		assert.Equal(t, components.VisibilityPrivate, res.UpdateVisibilityResponse.Visibility)
		t.Logf("Visibility changed to private")

		// Set back to public
		res, err = s.Images.UpdateVisibility(ctx, uploadedImageID, components.UpdateVisibilityRequest{
			Visibility: components.VisibilityPublic,
		})
		require.NoError(t, err)
		require.NotNil(t, res.UpdateVisibilityResponse)
		assert.Equal(t, components.VisibilityPublic, res.UpdateVisibilityResponse.Visibility)
		t.Logf("Visibility changed back to public")
	})

	// ---- 8. CreateSignedURL (Pro only) ----
	t.Run("CreateSignedURL", func(t *testing.T) {
		res, err := s.Images.CreateSignedURL(ctx, uploadedImageID, &components.CreateSignedURLRequest{
			ExpiresInSeconds: sdkgo.Int64(3600),
		})
		if err != nil {
			if isHTTPStatus(err, 403) {
				t.Skip("CreateSignedURL requires Pro plan, skipping")
			}
			require.NoError(t, err)
		}
		require.NotNil(t, res.SignedURLResponse)
		t.Logf("Signed URL created: %s", res.SignedURLResponse.SignedURL)
		assert.NotEmpty(t, res.SignedURLResponse.SignedURL)
	})

	// ---- 9. Presets CRUD (Pro only) ----
	t.Run("Presets", func(t *testing.T) {
		// List presets
		listRes, err := s.Presets.ListPresets(ctx)
		if err != nil {
			if isHTTPStatus(err, 403) {
				t.Skip("Presets require Pro plan, skipping")
			}
			require.NoError(t, err)
		}
		require.NotNil(t, listRes.ListPresetsResponse)
		initialCount := listRes.ListPresetsResponse.Total
		t.Logf("Initial preset count: %d", initialCount)

		// Create preset
		presetName := fmt.Sprintf("e2e-test-%d", time.Now().UnixMilli())
		createRes, err := s.Presets.CreatePreset(ctx, &components.CreatePresetRequest{
			Name:        presetName,
			Description: sdkgo.String("E2E test preset"),
			Params: map[string]any{
				"w":   800,
				"h":   600,
				"fit": "cover",
			},
		})
		require.NoError(t, err)
		require.NotNil(t, createRes.Preset)
		presetID := createRes.Preset.ID
		t.Logf("Created preset: ID=%s, Name=%s", presetID, createRes.Preset.Name)
		assert.Equal(t, presetName, createRes.Preset.Name)

		defer func() {
			_, delErr := s.Presets.DeletePreset(ctx, presetID)
			if delErr != nil {
				t.Logf("Cleanup warning: failed to delete preset: %v", delErr)
			}
		}()

		// Get preset
		getRes, err := s.Presets.GetPreset(ctx, presetID)
		require.NoError(t, err)
		require.NotNil(t, getRes.Preset)
		assert.Equal(t, presetName, getRes.Preset.Name)
		t.Logf("Got preset: %s", getRes.Preset.Name)

		// Update preset
		updatedName := presetName + "-updated"
		updateRes, err := s.Presets.UpdatePreset(ctx, presetID, &components.UpdatePresetRequest{
			Name: sdkgo.String(updatedName),
			Params: map[string]any{
				"w":   1024,
				"h":   768,
				"fit": "contain",
			},
		})
		require.NoError(t, err)
		require.NotNil(t, updateRes.Preset)
		assert.Equal(t, updatedName, updateRes.Preset.Name)
		t.Logf("Updated preset name to: %s", updateRes.Preset.Name)

		// Delete preset
		delRes, err := s.Presets.DeletePreset(ctx, presetID)
		require.NoError(t, err)
		require.NotNil(t, delRes.DeletePresetResponse)
		t.Logf("Deleted preset: %s", presetID)
	})

	// ---- 10. DeleteImage ----
	t.Run("DeleteImage", func(t *testing.T) {
		res, err := s.Images.Delete(ctx, uploadedImageID)
		require.NoError(t, err)
		require.NotNil(t, res.DeleteResponse)
		t.Logf("Deleted image: %s", uploadedImageID)

		// Clear so defer cleanup doesn't try again
		uploadedImageID = ""
	})

	// ---- 11. UpdateSettings ----
	t.Run("UpdateSettings", func(t *testing.T) {
		// Get current settings
		getRes, err := s.Settings.Get(ctx)
		require.NoError(t, err)
		originalQuality := getRes.SettingsResponse.Settings.DefaultQuality
		t.Logf("Original quality: %d", originalQuality)

		// Change quality
		newQuality := int64(75)
		if originalQuality == 75 {
			newQuality = 85
		}
		updateRes, err := s.Settings.Update(ctx, &components.UpdateSettingsRequest{
			DefaultQuality: &newQuality,
		})
		require.NoError(t, err)
		require.NotNil(t, updateRes.SettingsUpdateResponse)
		assert.Equal(t, newQuality, updateRes.SettingsUpdateResponse.Settings.DefaultQuality)
		t.Logf("Updated quality to: %d", newQuality)

		// Revert
		revertRes, err := s.Settings.Update(ctx, &components.UpdateSettingsRequest{
			DefaultQuality: &originalQuality,
		})
		require.NoError(t, err)
		require.NotNil(t, revertRes.SettingsUpdateResponse)
		assert.Equal(t, originalQuality, revertRes.SettingsUpdateResponse.Settings.DefaultQuality)
		t.Logf("Reverted quality to: %d", originalQuality)
	})
}

func isHTTPStatus(err error, statusCode int) bool {
	var apiErr *apierrors.APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == statusCode
	}
	var errResp *apierrors.ErrorResponse
	if errors.As(err, &errResp) {
		return errResp.Error_.Status == int64(statusCode)
	}
	return false
}
