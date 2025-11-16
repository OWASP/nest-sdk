# SnapshotRelease

Schema for Snapshot Release (used in list endpoints).


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `CreatedAt`                                | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | N/A                                        |
| `Name`                                     | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `PublishedAt`                              | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `TagName`                                  | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `OrganizationLogin`                        | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `RepositoryName`                           | *string*                                   | :heavy_check_mark:                         | N/A                                        |