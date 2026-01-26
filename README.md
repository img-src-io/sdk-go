# img-src Go SDK

Developer-friendly & type-safe Go SDK specifically catered to leverage *img-src* API.

[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)

<br /><br />

<!-- Start Summary [summary] -->
## Summary

img-src API: Image processing and delivery API.

A serverless image processing and delivery API built on Cloudflare Workers with parameter-driven image transformation and on-demand transcoding.

## Features

- **Image Upload**: Store original images in R2 with SHA256-based deduplication
- **On-Demand Transformation**: Resize, crop, and convert images via URL parameters
- **Format Conversion**: WebP, AVIF, JPEG, PNG output formats
- **Path Organization**: Organize images into folders with multiple paths per image
- **CDN Caching**: Automatic edge caching for transformed images

## Authentication

Authenticate using API Keys with `imgsrc_` prefix. Create your API key at https://img-src.io/settings

## Rate Limiting

- **Free Plan**: 100 requests/minute
- **Pro Plan**: 500 requests/minute

Rate limit headers are included in all responses.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [img-src Go SDK](#img-src-go-sdk)
  * [Features](#features)
  * [Authentication](#authentication)
  * [Rate Limiting](#rate-limiting)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication-1)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/img-src-io/sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Upload and Transform Images

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/img-src-io/sdk-go"
)

func main() {
    // Create API key at https://img-src.io/settings
    client := imgsrc.New(
        imgsrc.WithSecurity(os.Getenv("IMGSRC_API_KEY")),
    )

    ctx := context.Background()

    // Upload an image
    file, _ := os.Open("photo.jpg")
    defer file.Close()

    uploaded, _ := client.Images.UploadImage(ctx, file, "photos/2024")
    fmt.Printf("Uploaded: %s\n", uploaded.URL)

    // Access with transformations via CDN
    // https://img-src.io/i/{username}/photos/2024/photo.webp?w=800&h=600&fit=cover&q=85

    // List images
    images, _ := client.Images.ListImages(ctx, &imgsrc.ListImagesRequest{Limit: 20})
    fmt.Printf("Total: %d images\n", images.Total)
}
```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type | Scheme      |
| ------------ | ---- | ----------- |
| `BearerAuth` | http | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity("process.env["IMGSRC_API_KEY"]"),
	)

	res, err := s.Settings.GetSettings(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.SettingsResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Images](docs/sdks/images/README.md)

* [UploadImage](docs/sdks/images/README.md#uploadimage) - Upload image
* [ListImages](docs/sdks/images/README.md#listimages) - List images
* [SearchImages](docs/sdks/images/README.md#searchimages) - Search images
* [GetImage](docs/sdks/images/README.md#getimage) - Get image metadata
* [DeleteImage](docs/sdks/images/README.md#deleteimage) - Delete image
* [CreateSignedURL](docs/sdks/images/README.md#createsignedurl) - Create signed URL
* [DeleteImagePath](docs/sdks/images/README.md#deleteimagepath) - Delete image path

### [Presets](docs/sdks/presets/README.md)

* [ListPresets](docs/sdks/presets/README.md#listpresets) - List presets
* [CreatePreset](docs/sdks/presets/README.md#createpreset) - Create preset
* [GetPreset](docs/sdks/presets/README.md#getpreset) - Get preset
* [UpdatePreset](docs/sdks/presets/README.md#updatepreset) - Update preset
* [DeletePreset](docs/sdks/presets/README.md#deletepreset) - Delete preset

### [Settings](docs/sdks/settings/README.md)

* [GetSettings](docs/sdks/settings/README.md#getsettings) - Get user settings
* [UpdateSettings](docs/sdks/settings/README.md#updatesettings) - Update user settings

### [Usage](docs/sdks/usage/README.md)

* [GetUsage](docs/sdks/usage/README.md#getusage) - Get usage statistics

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"github.com/img-src-io/sdk-go/optionalnullable"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity("process.env["IMGSRC_API_KEY"]"),
	)

	res, err := s.Images.ListImages(ctx, sdkgo.Pointer[int64](50), optionalnullable.From(sdkgo.Pointer[int64](0)), sdkgo.Pointer("blog/2024"))
	if err != nil {
		log.Fatal(err)
	}
	if res.ImageListResponse != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"github.com/img-src-io/sdk-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity("process.env["IMGSRC_API_KEY"]"),
	)

	res, err := s.Settings.GetSettings(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.SettingsResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"github.com/img-src-io/sdk-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		sdkgo.WithSecurity("process.env["IMGSRC_API_KEY"]"),
	)

	res, err := s.Settings.GetSettings(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.SettingsResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetSettings` function may return the following errors:

| Error Type              | Status Code | Content Type     |
| ----------------------- | ----------- | ---------------- |
| apierrors.ErrorResponse | 401         | application/json |
| apierrors.ErrorResponse | 500         | application/json |
| apierrors.APIError      | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	sdkgo "github.com/img-src-io/sdk-go"
	"github.com/img-src-io/sdk-go/models/apierrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity("process.env["IMGSRC_API_KEY"]"),
	)

	res, err := s.Settings.GetSettings(ctx)
	if err != nil {

		var e *apierrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithServerURL("https://api.img-src.io"),
		sdkgo.WithSecurity("process.env["IMGSRC_API_KEY"]"),
	)

	res, err := s.Settings.GetSettings(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.SettingsResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/img-src-io/sdk-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdkgo.New(sdkgo.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release.
