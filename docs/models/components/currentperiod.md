# CurrentPeriod


## Fields

| Field                               | Type                                | Required                            | Description                         | Example                             |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `Period`                            | *string*                            | :heavy_check_mark:                  | Period identifier (YYYY-MM format)  | 2025-01                             |
| `PeriodStart`                       | *int64*                             | :heavy_check_mark:                  | Unix timestamp of period start      | 1735689600                          |
| `PeriodEnd`                         | *int64*                             | :heavy_check_mark:                  | Unix timestamp of period end        | 1738368000                          |
| `Uploads`                           | *int64*                             | :heavy_check_mark:                  | Uploads this period                 | 42                                  |
| `BandwidthBytes`                    | *int64*                             | :heavy_check_mark:                  | Bandwidth used this period in bytes | 1073741824                          |
| `APIRequests`                       | *int64*                             | :heavy_check_mark:                  | API requests this period            | 5000                                |
| `Transformations`                   | *int64*                             | :heavy_check_mark:                  | Image transformations this period   | 500                                 |