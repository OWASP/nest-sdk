# ReleaseDetail

Detail schema for Release (used in single item endpoints).


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `CreatedAt`                                | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | N/A                                        |
| `Name`                                     | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `PublishedAt`                              | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `TagName`                                  | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `Description`                              | *string*                                   | :heavy_check_mark:                         | N/A                                        |