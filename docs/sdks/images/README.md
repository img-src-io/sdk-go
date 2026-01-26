# Images

## Overview

### Available Operations

* [Upload](#upload) - Upload image
* [List](#list) - List images
* [Search](#search) - Search images
* [GetMetadata](#getmetadata) - Get image metadata
* [Delete](#delete) - Delete image
* [CreateSignedURL](#createsignedurl) - Create signed URL
* [DeletePath](#deletepath) - Delete image path

## Upload

Upload a new image. Supports multipart/form-data with 'file' field.

### Example Usage

<!-- UsageSnippet language="go" operationID="uploadImage" method="post" path="/api/v1/images" -->
```go
package main

import(
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"github.com/img-src-io/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Images.Upload(ctx, &operations.UploadImageRequestBody{
        TargetPath: sdkgo.Pointer("blog/2024"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UploadImageRequestBody](../../models/operations/uploadimagerequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UploadImageResponse](../../models/operations/uploadimageresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 409, 413      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

List user's images with pagination and optional path filtering

### Example Usage

<!-- UsageSnippet language="go" operationID="listImages" method="get" path="/api/v1/images" -->
```go
package main

import(
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Images.List(ctx, sdkgo.Pointer[int64](50), sdkgo.Pointer[int64](0), sdkgo.Pointer("blog/2024"))
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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `offset`                                                 | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `path`                                                   | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListImagesResponse](../../models/operations/listimagesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Search

Search images by filename

### Example Usage

<!-- UsageSnippet language="go" operationID="searchImages" method="get" path="/api/v1/images/search" -->
```go
package main

import(
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Images.Search(ctx, "vacation", sdkgo.Pointer[int64](20))
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `q`                                                      | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SearchImagesResponse](../../models/operations/searchimagesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetMetadata

Get metadata for a specific image

### Example Usage

<!-- UsageSnippet language="go" operationID="getImage" method="get" path="/api/v1/images/{id}" -->
```go
package main

import(
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Images.GetMetadata(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.MetadataResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetImageResponse](../../models/operations/getimageresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 404                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Delete

Delete an image and all its paths

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteImage" method="delete" path="/api/v1/images/{id}" -->
```go
package main

import(
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Images.Delete(ctx, "abcdef1234567890")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteImageResponse](../../models/operations/deleteimageresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 404                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CreateSignedURL

Create a time-limited signed URL for an image (Pro plan only)

### Example Usage

<!-- UsageSnippet language="go" operationID="createSignedUrl" method="post" path="/api/v1/images/{id}/signed-url" -->
```go
package main

import(
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"github.com/img-src-io/sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Images.CreateSignedURL(ctx, "abcdef1234567890", &components.CreateSignedURLRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SignedURLResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `id`                                                                                    | *string*                                                                                | :heavy_check_mark:                                                                      | N/A                                                                                     |
| `body`                                                                                  | [*components.CreateSignedURLRequest](../../models/components/createsignedurlrequest.md) | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |

### Response

**[*operations.CreateSignedURLResponse](../../models/operations/createsignedurlresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeletePath

Delete a specific path from an image. If this is the last path, the image is deleted.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteImagePath" method="delete" path="/api/v1/images/path/{username}/{filepath}" -->
```go
package main

import(
	"context"
	sdkgo "github.com/img-src-io/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity("imgsrc_your_api_key_here"),
    )

    res, err := s.Images.DeletePath(ctx, "johndoe", "blog/2024/photo.webp")
    if err != nil {
        log.Fatal(err)
    }
    if res.PathDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `username`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `filepath`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteImagePathResponse](../../models/operations/deleteimagepathresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 403, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |