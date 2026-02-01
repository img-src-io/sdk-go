# Settings

## Overview

### Available Operations

* [Get](#get) - Get user settings
* [Update](#update) - Update user settings

## Get

Returns the authenticated user's settings

### Example Usage

<!-- UsageSnippet language="go" operationID="getSettings" method="get" path="/api/v1/settings" -->
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

    res, err := s.Settings.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.SettingsResponse != nil {
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

**[*operations.GetSettingsResponse](../../models/operations/getsettingsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Updates the authenticated user's settings

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSettings" method="put" path="/api/v1/settings" -->
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

    res, err := s.Settings.Update(ctx, &components.UpdateSettingsRequest{
        DeliveryFormats: []string{
            "webp",
            "avif",
            "jpeg",
        },
        DefaultQuality: imgsrc.Pointer[int64](80),
        DefaultFitMode: imgsrc.Pointer("contain"),
        DefaultMaxWidth: optionalnullable.From(imgsrc.Pointer[int64](1920)),
        DefaultMaxHeight: optionalnullable.From(imgsrc.Pointer[int64](1080)),
        Theme: imgsrc.Pointer("dark"),
        Language: imgsrc.Pointer("ko"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SettingsUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.UpdateSettingsRequest](../../models/components/updatesettingsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateSettingsResponse](../../models/operations/updatesettingsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |