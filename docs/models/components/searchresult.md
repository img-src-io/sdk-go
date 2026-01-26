# SearchResult


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     | Example                                         |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `ID`                                            | *string*                                        | :heavy_check_mark:                              | Image ID                                        | abcdef1234567890                                |
| `OriginalFilename`                              | *string*                                        | :heavy_check_mark:                              | Original filename                               | vacation-photo.jpg                              |
| `SanitizedFilename`                             | **string*                                       | :heavy_minus_sign:                              | Sanitized filename                              | vacation-photo.jpg                              |
| `Paths`                                         | []*string*                                      | :heavy_check_mark:                              | All paths for this image                        | [<br/>"photos/vacation.jpg"<br/>]               |
| `Size`                                          | *int64*                                         | :heavy_check_mark:                              | File size in bytes                              | 1024000                                         |
| `UploadedAt`                                    | [time.Time](https://pkg.go.dev/time#Time)       | :heavy_check_mark:                              | Upload timestamp                                | 2025-01-21T12:00:00Z                            |
| `URL`                                           | *string*                                        | :heavy_check_mark:                              | API endpoint URL                                | /api/v1/images/abcdef1234567890                 |
| `CdnURL`                                        | **string*                                       | :heavy_minus_sign:                              | CDN URL                                         | https://cdn.img-src.io/john/photos/vacation.jpg |