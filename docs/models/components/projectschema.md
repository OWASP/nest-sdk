# ProjectSchema

Schema for Project.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `CreatedAt`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | N/A                                                                |
| `Description`                                                      | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `Level`                                                            | [components.ProjectLevel](../../models/components/projectlevel.md) | :heavy_check_mark:                                                 | Enum for OWASP project levels.                                     |
| `Name`                                                             | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `UpdatedAt`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | N/A                                                                |