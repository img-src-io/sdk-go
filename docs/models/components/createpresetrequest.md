# CreatePresetRequest


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Name`                                                   | *string*                                                 | :heavy_check_mark:                                       | Preset name (1-50 characters)                            | thumbnail                                                |
| `Description`                                            | **string*                                                | :heavy_minus_sign:                                       | Optional description (max 200 characters)                | 200x200 thumbnail with cover fit                         |
| `Params`                                                 | map[string]*any*                                         | :heavy_check_mark:                                       | Transformation parameters                                | {<br/>"w": 200,<br/>"h": 200,<br/>"fit": "cover",<br/>"format": "webp"<br/>} |