# CreateSignedURLRequest


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ExpiresInSeconds`                                                      | **int64*                                                                | :heavy_minus_sign:                                                      | Expiration time in seconds (60-86400, default 3600)                     | 3600                                                                    |
| `Transformation`                                                        | [*components.Transformation](../../models/components/transformation.md) | :heavy_minus_sign:                                                      | Optional image transformation parameters                                |                                                                         |