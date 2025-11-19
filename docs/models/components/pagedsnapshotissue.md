# PagedSnapshotIssue


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `CurrentPage`                                                          | *int64*                                                                | :heavy_check_mark:                                                     | Current page number                                                    |
| `HasNext`                                                              | *bool*                                                                 | :heavy_check_mark:                                                     | Whether there is a next page                                           |
| `HasPrevious`                                                          | *bool*                                                                 | :heavy_check_mark:                                                     | Whether there is a previous page                                       |
| `Items`                                                                | [][components.SnapshotIssue](../../models/components/snapshotissue.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `TotalCount`                                                           | *int64*                                                                | :heavy_check_mark:                                                     | Total number of items                                                  |
| `TotalPages`                                                           | *int64*                                                                | :heavy_check_mark:                                                     | Total number of pages                                                  |