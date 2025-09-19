# Repositories
(*Repositories*)

## Overview

### Available Operations

* [ListRepositories](#listrepositories) - List repositories

## ListRepositories

Retrieve a paginated list of GitHub repositories.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_repositories" method="get" path="/api/v0/repositories/" -->
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

    res, err := s.Repositories.ListRepositories(ctx, nil, nest.Pointer[int64](1), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedRepositorySchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `ordering`                                                                                  | [*operations.ListRepositoriesOrdering](../../models/operations/listrepositoriesordering.md) | :heavy_minus_sign:                                                                          | Ordering field                                                                              |
| `page`                                                                                      | **int64*                                                                                    | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `pageSize`                                                                                  | **int64*                                                                                    | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.ListRepositoriesResponse](../../models/operations/listrepositoriesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |