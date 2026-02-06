# UserSettings


## Fields

| Field                                   | Type                                    | Required                                | Description                             | Example                                 |
| --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| `ID`                                    | *string*                                | :heavy_check_mark:                      | Clerk user ID                           | user_abc123                             |
| `Username`                              | *string*                                | :heavy_check_mark:                      | Username                                | johndoe                                 |
| `Email`                                 | **string*                               | :heavy_minus_sign:                      | Email address                           | john@example.com                        |
| `DeliveryFormats`                       | []*string*                              | :heavy_check_mark:                      | Preferred delivery formats (ordered)    | [<br/>"webp",<br/>"avif",<br/>"jpeg",<br/>"jxl"<br/>] |
| `DefaultQuality`                        | *int64*                                 | :heavy_check_mark:                      | Default image quality (1-100)           | 80                                      |
| `DefaultFitMode`                        | *string*                                | :heavy_check_mark:                      | Default fit mode                        | contain                                 |
| `DefaultMaxWidth`                       | **int64*                                | :heavy_minus_sign:                      | Default maximum width                   | 1920                                    |
| `DefaultMaxHeight`                      | **int64*                                | :heavy_minus_sign:                      | Default maximum height                  | 1080                                    |
| `Theme`                                 | *string*                                | :heavy_check_mark:                      | UI theme                                | light                                   |
| `Language`                              | *string*                                | :heavy_check_mark:                      | UI language                             | en                                      |
| `CreatedAt`                             | *int64*                                 | :heavy_check_mark:                      | Account creation timestamp (Unix epoch) | 1704067200                              |
| `UpdatedAt`                             | *int64*                                 | :heavy_check_mark:                      | Last update timestamp (Unix epoch)      | 1704067200                              |
| `TotalUploads`                          | *int64*                                 | :heavy_check_mark:                      | Total number of uploads                 | 150                                     |
| `StorageUsedBytes`                      | *int64*                                 | :heavy_check_mark:                      | Total storage used in bytes             | 104857600                               |