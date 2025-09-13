# Committees
(*Committees*)

## Overview

### Available Operations

* [ListCommittees](#listcommittees) - List committees
* [GetCommittee](#getcommittee) - Get committee

## ListCommittees

Retrieve a paginated list of OWASP committees.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_committees" method="get" path="/api/v0/committees/" -->
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

    res, err := s.Committees.ListCommittees(ctx, nil, nest.Int64(1), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedCommitteeSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `ordering`                                                                              | [*operations.ListCommitteesOrdering](../../models/operations/listcommitteesordering.md) | :heavy_minus_sign:                                                                      | Ordering field                                                                          |
| `page`                                                                                  | **int64*                                                                                | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `pageSize`                                                                              | **int64*                                                                                | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |

### Response

**[*operations.ListCommitteesResponse](../../models/operations/listcommitteesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetCommittee

Retrieve committee details.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_committee" method="get" path="/api/v0/committees/{committee_id}" -->
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

    res, err := s.Committees.GetCommittee(ctx, "project")
    if err != nil {
        log.Fatal(err)
    }
    if res.CommitteeSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `committeeID`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | project                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCommitteeResponse](../../models/operations/getcommitteeresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.CommitteeErrorResponse | 404                              | application/json                 |
| apierrors.NestAPIError           | 4XX, 5XX                         | \*/\*                            |