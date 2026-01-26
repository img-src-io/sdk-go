# UpdateSettingsRequest


## Fields

| Field                                  | Type                                   | Required                               | Description                            | Example                                |
| -------------------------------------- | -------------------------------------- | -------------------------------------- | -------------------------------------- | -------------------------------------- |
| `DeliveryFormats`                      | []*string*                             | :heavy_minus_sign:                     | Preferred delivery formats (ordered)   | [<br/>"webp",<br/>"avif",<br/>"jpeg"<br/>] |
| `DefaultQuality`                       | **int64*                               | :heavy_minus_sign:                     | Default image quality (1-100)          | 80                                     |
| `DefaultFitMode`                       | **string*                              | :heavy_minus_sign:                     | Default fit mode                       | contain                                |
| `DefaultMaxWidth`                      | **int64*                               | :heavy_minus_sign:                     | Default maximum width (null to clear)  | 1920                                   |
| `DefaultMaxHeight`                     | **int64*                               | :heavy_minus_sign:                     | Default maximum height (null to clear) | 1080                                   |
| `Theme`                                | **string*                              | :heavy_minus_sign:                     | UI theme                               | dark                                   |
| `Language`                             | **string*                              | :heavy_minus_sign:                     | UI language                            | ko                                     |