# Issues
(*Issues*)

## Overview

### Available Operations

* [ListIssues](#listissues) - List issues
* [GetIssue](#getissue) - Get issue

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
	"github.com/owasp/nest-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := nest.New(
        nest.WithSecurity(os.Getenv("NEST_API_KEY")),
    )

    res, err := s.Issues.ListIssues(ctx, operations.ListIssuesRequest{
        Organization: nest.Pointer("OWASP"),
        Repository: nest.Pointer("Nest"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedIssue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListIssuesRequest](../../models/operations/listissuesrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListIssuesResponse](../../models/operations/listissuesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetIssue

Retrieve a specific GitHub issue by organization, repository, and issue number.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_issue" method="get" path="/api/v0/issues/{organization_id}/{repository_id}/{issue_id}" -->
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

    res, err := s.Issues.GetIssue(ctx, "OWASP", "Nest", 1234)
    if err != nil {
        log.Fatal(err)
    }
    if res.IssueDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `organizationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | OWASP                                                    |
| `repositoryID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | Nest                                                     |
| `issueID`                                                | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      | 1234                                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetIssueResponse](../../models/operations/getissueresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.IssueError   | 404                    | application/json       |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |