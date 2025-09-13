# Issues
(*Issues*)

## Overview

### Available Operations

* [ListIssues](#listissues) - List issues

## ListIssues

Retrieve a paginated list of GitHub issues.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_issues" method="get" path="/api/v0/issues/" -->
```go
package main

import(
	"context"
	"os"
	nest "github.com/owasp/nest-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := nest.New(
        nest.WithSecurity(os.Getenv("NEST_API_KEY")),
    )

    res, err := s.Issues.ListIssues(ctx, nil, nil, nest.Int64(1), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedIssueSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |
| `state`                                                                         | [*components.State](../../models/components/state.md)                           | :heavy_minus_sign:                                                              | State of the issue                                                              |
| `ordering`                                                                      | [*operations.ListIssuesOrdering](../../models/operations/listissuesordering.md) | :heavy_minus_sign:                                                              | Ordering field                                                                  |
| `page`                                                                          | **int64*                                                                        | :heavy_minus_sign:                                                              | N/A                                                                             |
| `pageSize`                                                                      | **int64*                                                                        | :heavy_minus_sign:                                                              | N/A                                                                             |
| `opts`                                                                          | [][operations.Option](../../models/operations/option.md)                        | :heavy_minus_sign:                                                              | The options for this request.                                                   |

### Response

**[*operations.ListIssuesResponse](../../models/operations/listissuesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |