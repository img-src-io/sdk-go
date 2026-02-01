# Presets

## Overview

### Available Operations

* [ListPresets](#listpresets) - List presets
* [CreatePreset](#createpreset) - Create preset
* [GetPreset](#getpreset) - Get preset
* [UpdatePreset](#updatepreset) - Update preset
* [DeletePreset](#deletepreset) - Delete preset

## ListPresets

Returns all transformation presets for the authenticated user. Requires Pro plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPresets" method="get" path="/api/v1/settings/presets" -->
```go
package main

import(
	"context"
	imgsrc "github.com/img-src-io/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := imgsrc.New(
        imgsrc.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Presets.ListPresets(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPresetsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListPresetsResponse](../../models/operations/listpresetsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CreatePreset

Creates a new transformation preset. Requires Pro plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPreset" method="post" path="/api/v1/settings/presets" -->
```go
package main

import(
	"context"
	imgsrc "github.com/img-src-io/sdk-go"
	"github.com/img-src-io/sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := imgsrc.New(
        imgsrc.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Presets.CreatePreset(ctx, &components.CreatePresetRequest{
        Name: "thumbnail",
        Description: imgsrc.Pointer("200x200 thumbnail with cover fit"),
        Params: map[string]any{
            "w": 200,
            "h": 200,
            "fit": "cover",
            "format": "webp",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Preset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.CreatePresetRequest](../../models/components/createpresetrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreatePresetResponse](../../models/operations/createpresetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 409      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetPreset

Returns a specific preset by ID. Requires Pro plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPreset" method="get" path="/api/v1/settings/presets/{id}" -->
```go
package main

import(
	"context"
	imgsrc "github.com/img-src-io/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := imgsrc.New(
        imgsrc.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Presets.GetPreset(ctx, "preset_abc123")
    if err != nil {
        log.Fatal(err)
    }
    if res.Preset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | preset_abc123                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPresetResponse](../../models/operations/getpresetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdatePreset

Updates an existing preset. Requires Pro plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePreset" method="put" path="/api/v1/settings/presets/{id}" -->
```go
package main

import(
	"context"
	imgsrc "github.com/img-src-io/sdk-go"
	"github.com/img-src-io/sdk-go/optionalnullable"
	"github.com/img-src-io/sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := imgsrc.New(
        imgsrc.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Presets.UpdatePreset(ctx, "preset_abc123", &components.UpdatePresetRequest{
        Name: imgsrc.Pointer("card-image"),
        Description: optionalnullable.From(imgsrc.Pointer("Card thumbnail for product listings")),
        Params: map[string]any{
            "w": 400,
            "h": 300,
            "fit": "cover",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Preset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |                                                                                   |
| `id`                                                                              | *string*                                                                          | :heavy_check_mark:                                                                | N/A                                                                               | preset_abc123                                                                     |
| `body`                                                                            | [*components.UpdatePresetRequest](../../models/components/updatepresetrequest.md) | :heavy_minus_sign:                                                                | N/A                                                                               |                                                                                   |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |                                                                                   |

### Response

**[*operations.UpdatePresetResponse](../../models/operations/updatepresetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeletePreset

Deletes a preset. Requires Pro plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePreset" method="delete" path="/api/v1/settings/presets/{id}" -->
```go
package main

import(
	"context"
	imgsrc "github.com/img-src-io/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := imgsrc.New(
        imgsrc.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Presets.DeletePreset(ctx, "preset_abc123")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeletePresetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | preset_abc123                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePresetResponse](../../models/operations/deletepresetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |