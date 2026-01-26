# DeleteResponse


## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Success`                                 | *bool*                                    | :heavy_check_mark:                        | Operation success flag                    | true                                      |
| `Message`                                 | *string*                                  | :heavy_check_mark:                        | Human-readable message                    | Image deleted                             |
| `DeletedPaths`                            | []*string*                                | :heavy_minus_sign:                        | List of deleted paths                     | [<br/>"photo.webp",<br/>"blog/photo.webp"<br/>] |
| `DeletedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | Deletion timestamp                        | 2025-01-21T12:00:00Z                      |