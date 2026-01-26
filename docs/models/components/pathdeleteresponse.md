# PathDeleteResponse


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Success`                                                | *bool*                                                   | :heavy_check_mark:                                       | Operation success flag                                   | true                                                     |
| `Message`                                                | *string*                                                 | :heavy_check_mark:                                       | Human-readable message                                   | Path removed                                             |
| `RemainingPaths`                                         | []*string*                                               | :heavy_check_mark:                                       | Remaining paths for the image                            | [<br/>"blog/photo.webp"<br/>]                            |
| `ImageDeleted`                                           | *bool*                                                   | :heavy_check_mark:                                       | Whether the image itself was deleted (last path removed) | false                                                    |
| `DeletedAt`                                              | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | Deletion timestamp                                       | 2025-01-21T12:00:00Z                                     |