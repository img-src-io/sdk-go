# UpdatePresetRequest


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         | Example                                             |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Name`                                              | **string*                                           | :heavy_minus_sign:                                  | New preset name (1-50 characters)                   | card-image                                          |
| `Description`                                       | **string*                                           | :heavy_minus_sign:                                  | New description (max 200 characters, null to clear) | Card thumbnail for product listings                 |
| `Params`                                            | map[string]*any*                                    | :heavy_minus_sign:                                  | New transformation parameters                       | {<br/>"w": 400,<br/>"h": 300,<br/>"fit": "cover"<br/>} |