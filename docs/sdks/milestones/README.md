# Milestones

## Overview

### Available Operations

* [ListMilestones](#listmilestones) - List milestones
* [GetMilestone](#getmilestone) - Get milestone

## ListMilestones

Retrieve a paginated list of GitHub milestones.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_milestones" method="get" path="/api/v0/milestones/" -->
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

    res, err := s.Milestones.ListMilestones(ctx, operations.ListMilestonesRequest{
        Organization: nest.Pointer("OWASP"),
        Repository: nest.Pointer("Nest"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedMilestone != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListMilestonesRequest](../../models/operations/listmilestonesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListMilestonesResponse](../../models/operations/listmilestonesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetMilestone

Retrieve a specific GitHub milestone by organization, repository, and milestone number.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_milestone" method="get" path="/api/v0/milestones/{organization_id}/{repository_id}/{milestone_id}" -->
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

    res, err := s.Milestones.GetMilestone(ctx, "OWASP", "Nest", 1)
    if err != nil {
        log.Fatal(err)
    }
    if res.MilestoneDetail != nil {
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
| `milestoneID`                                            | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      | 1                                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMilestoneResponse](../../models/operations/getmilestoneresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| apierrors.MilestoneError | 404                      | application/json         |
| apierrors.NestAPIError   | 4XX, 5XX                 | \*/\*                    |