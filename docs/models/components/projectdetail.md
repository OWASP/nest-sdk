# ProjectDetail

Detail schema for Project (used in single item endpoints).


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `CreatedAt`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | N/A                                                                |
| `Key`                                                              | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `Level`                                                            | [components.ProjectLevel](../../models/components/projectlevel.md) | :heavy_check_mark:                                                 | Enum for OWASP project levels.                                     |
| `Name`                                                             | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `UpdatedAt`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | N/A                                                                |
| `Description`                                                      | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `Leaders`                                                          | [][components.Leader](../../models/components/leader.md)           | :heavy_check_mark:                                                 | N/A                                                                |