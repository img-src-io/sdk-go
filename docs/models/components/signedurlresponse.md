# SignedURLResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `SignedURL`                                                         | *string*                                                            | :heavy_check_mark:                                                  | Time-limited signed URL                                             | https://cdn.img-src.io/john/photo.webp?token=xxx&expires=1704153600 |
| `ExpiresAt`                                                         | *int64*                                                             | :heavy_check_mark:                                                  | Expiration timestamp (Unix epoch)                                   | 1704153600                                                          |
| `ExpiresInSeconds`                                                  | *int64*                                                             | :heavy_check_mark:                                                  | Seconds until expiration                                            | 3600                                                                |