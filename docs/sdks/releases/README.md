# Releases
(*Releases*)

## Overview

### Available Operations

* [ListReleases](#listreleases) - List releases

## ListReleases

Retrieve a paginated list of GitHub releases.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_releases" method="get" path="/api/v0/releases/" -->
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
        nest.WithSecurity(os.Getenv("NEST_API_KEY_HEADER")),
    )

    res, err := s.Releases.ListReleases(ctx, nest.String("v1.0.0"), nil, nest.Int64(1), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedReleaseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |                                                                                     |
| `tagName`                                                                           | **string*                                                                           | :heavy_minus_sign:                                                                  | Tag name of the release                                                             | v1.0.0                                                                              |
| `ordering`                                                                          | [*operations.ListReleasesOrdering](../../models/operations/listreleasesordering.md) | :heavy_minus_sign:                                                                  | Ordering field                                                                      |                                                                                     |
| `page`                                                                              | **int64*                                                                            | :heavy_minus_sign:                                                                  | N/A                                                                                 |                                                                                     |
| `pageSize`                                                                          | **int64*                                                                            | :heavy_minus_sign:                                                                  | N/A                                                                                 |                                                                                     |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |                                                                                     |

### Response

**[*operations.ListReleasesResponse](../../models/operations/listreleasesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |