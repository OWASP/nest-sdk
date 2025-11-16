# RepositoryDetail

Detail schema for Repository (used in single item endpoints).


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `CreatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |
| `Name`                                    | *string*                                  | :heavy_check_mark:                        | N/A                                       |
| `UpdatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |
| `CommitsCount`                            | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `ContributorsCount`                       | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `Description`                             | **string*                                 | :heavy_minus_sign:                        | N/A                                       |
| `ForksCount`                              | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `OpenIssuesCount`                         | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `StarsCount`                              | *int64*                                   | :heavy_check_mark:                        | N/A                                       |