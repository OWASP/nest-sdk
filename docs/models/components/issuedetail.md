# IssueDetail

Detail schema for Issue (used in single item endpoints).


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `CreatedAt`                                          | [time.Time](https://pkg.go.dev/time#Time)            | :heavy_check_mark:                                   | N/A                                                  |
| `State`                                              | [components.State](../../models/components/state.md) | :heavy_check_mark:                                   | N/A                                                  |
| `Title`                                              | *string*                                             | :heavy_check_mark:                                   | N/A                                                  |
| `UpdatedAt`                                          | [time.Time](https://pkg.go.dev/time#Time)            | :heavy_check_mark:                                   | N/A                                                  |
| `URL`                                                | *string*                                             | :heavy_check_mark:                                   | N/A                                                  |
| `Body`                                               | *string*                                             | :heavy_check_mark:                                   | N/A                                                  |