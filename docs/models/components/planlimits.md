# PlanLimits


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `MaxUploadsPerMonth`                                       | *int64*                                                    | :heavy_check_mark:                                         | Maximum uploads per month (null = unlimited)               | 1000                                                       |
| `MaxStorageBytes`                                          | *int64*                                                    | :heavy_check_mark:                                         | Maximum storage in bytes (null = unlimited)                | 5368709120                                                 |
| `MaxBandwidthPerMonth`                                     | *int64*                                                    | :heavy_check_mark:                                         | Maximum bandwidth per month in bytes (null = unlimited)    | 10737418240                                                |
| `MaxAPIRequestsPerMonth`                                   | *int64*                                                    | :heavy_check_mark:                                         | Maximum API requests per month (null = unlimited)          | 100000                                                     |
| `MaxTransformationsPerMonth`                               | *int64*                                                    | :heavy_check_mark:                                         | Maximum image transformations per month (null = unlimited) | 10000                                                      |