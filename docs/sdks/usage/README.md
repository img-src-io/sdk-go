# Usage

## Overview

### Available Operations

* [Get](#get) - Get usage statistics

## Get

Returns usage statistics for the authenticated user

### Example Usage

<!-- UsageSnippet language="go" operationID="getUsage" method="get" path="/api/v1/usage" -->
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

    res, err := s.Usage.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.UsageResponse != nil {
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

**[*operations.GetUsageResponse](../../models/operations/getusageresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |