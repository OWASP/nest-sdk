# Projects
(*Projects*)

## Overview

### Available Operations

* [ListProjects](#listprojects) - List projects

## ListProjects

Retrieve a paginated list of OWASP projects.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_projects" method="get" path="/api/v1/projects/" -->
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
        nest.WithSecurity(os.Getenv("NEST_API_KEY_AUTH")),
    )

    res, err := s.Projects.ListProjects(ctx, nil, nil, nest.Int64(1), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedProjectSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `level`                                                                             | [*components.ProjectLevel](../../models/components/projectlevel.md)                 | :heavy_minus_sign:                                                                  | Level of the project                                                                |
| `ordering`                                                                          | [*operations.ListProjectsOrdering](../../models/operations/listprojectsordering.md) | :heavy_minus_sign:                                                                  | Ordering field                                                                      |
| `page`                                                                              | **int64*                                                                            | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `pageSize`                                                                          | **int64*                                                                            | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |

### Response

**[*operations.ListProjectsResponse](../../models/operations/listprojectsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |