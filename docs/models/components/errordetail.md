# ErrorDetail


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Code`                                                       | *string*                                                     | :heavy_check_mark:                                           | Error code (e.g., NOT_FOUND, UNAUTHORIZED, VALIDATION_ERROR) | NOT_FOUND                                                    |
| `Message`                                                    | *string*                                                     | :heavy_check_mark:                                           | Human-readable error message                                 | The requested resource was not found                         |
| `Status`                                                     | *int64*                                                      | :heavy_check_mark:                                           | HTTP status code                                             | 404                                                          |
| `Path`                                                       | **string*                                                    | :heavy_minus_sign:                                           | Request path (optional)                                      | /api/v1/images/nonexistent                                   |