# Projects
(*Projects*)

## Overview

### Available Operations

* [ListProjects](#listprojects) - List projects
* [GetProject](#getproject) - Get project

## ListProjects

Retrieve a paginated list of OWASP projects.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_projects" method="get" path="/api/v0/projects/" -->
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

    res, err := s.Projects.ListProjects(ctx, nil, operations.ListProjectsOrderingMinusCreatedAt.ToPointer(), nest.Pointer[int64](1), nest.Pointer[int64](100))
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedProject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |                                                                                     |
| `level`                                                                             | [*components.ProjectLevel](../../models/components/projectlevel.md)                 | :heavy_minus_sign:                                                                  | Level of the project                                                                |                                                                                     |
| `ordering`                                                                          | [*operations.ListProjectsOrdering](../../models/operations/listprojectsordering.md) | :heavy_minus_sign:                                                                  | Ordering field                                                                      | -created_at                                                                         |
| `page`                                                                              | **int64*                                                                            | :heavy_minus_sign:                                                                  | Page number                                                                         |                                                                                     |
| `pageSize`                                                                          | **int64*                                                                            | :heavy_minus_sign:                                                                  | Number of items per page                                                            |                                                                                     |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |                                                                                     |

### Response

**[*operations.ListProjectsResponse](../../models/operations/listprojectsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetProject

Retrieve project details.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_project" method="get" path="/api/v0/projects/{project_id}" -->
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

    res, err := s.Projects.GetProject(ctx, "Nest")
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | Nest                                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetProjectResponse](../../models/operations/getprojectresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ProjectError | 404                    | application/json       |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |