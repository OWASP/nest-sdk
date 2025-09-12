# Projects
(*Projects*)

## Overview

### Available Operations

* [AppsAPIRestV0ProjectListProjects](#appsapirestv0projectlistprojects) - List projects
* [AppsAPIRestV0ProjectGetProject](#appsapirestv0projectgetproject) - Get project

## AppsAPIRestV0ProjectListProjects

Retrieve a paginated list of OWASP projects.

### Example Usage

<!-- UsageSnippet language="go" operationID="apps_api_rest_v0_project_list_projects" method="get" path="/api/v0/projects/" -->
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
        nest.WithSecurity(os.Getenv("NEST_API_KEY_HEADER")),
    )

    res, err := s.Projects.AppsAPIRestV0ProjectListProjects(ctx, nil, operations.AppsAPIRestV0ProjectListProjectsOrderingMinusCreatedAt.ToPointer(), nest.Int64(1), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedProjectSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                   | Type                                                                                                                        | Required                                                                                                                    | Description                                                                                                                 | Example                                                                                                                     |
| --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                       | :heavy_check_mark:                                                                                                          | The context to use for the request.                                                                                         |                                                                                                                             |
| `level`                                                                                                                     | [*components.ProjectLevel](../../models/components/projectlevel.md)                                                         | :heavy_minus_sign:                                                                                                          | Level of the project                                                                                                        |                                                                                                                             |
| `ordering`                                                                                                                  | [*operations.AppsAPIRestV0ProjectListProjectsOrdering](../../models/operations/appsapirestv0projectlistprojectsordering.md) | :heavy_minus_sign:                                                                                                          | Ordering field                                                                                                              | -created_at                                                                                                                 |
| `page`                                                                                                                      | **int64*                                                                                                                    | :heavy_minus_sign:                                                                                                          | N/A                                                                                                                         |                                                                                                                             |
| `pageSize`                                                                                                                  | **int64*                                                                                                                    | :heavy_minus_sign:                                                                                                          | N/A                                                                                                                         |                                                                                                                             |
| `opts`                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                    | :heavy_minus_sign:                                                                                                          | The options for this request.                                                                                               |                                                                                                                             |

### Response

**[*operations.AppsAPIRestV0ProjectListProjectsResponse](../../models/operations/appsapirestv0projectlistprojectsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## AppsAPIRestV0ProjectGetProject

Retrieve project details.

### Example Usage

<!-- UsageSnippet language="go" operationID="apps_api_rest_v0_project_get_project" method="get" path="/api/v0/projects/{project_id}" -->
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

    res, err := s.Projects.AppsAPIRestV0ProjectGetProject(ctx, "Nest")
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectSchema != nil {
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

**[*operations.AppsAPIRestV0ProjectGetProjectResponse](../../models/operations/appsapirestv0projectgetprojectresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ProjectErrorResponse | 404                            | application/json               |
| apierrors.NestAPIError         | 4XX, 5XX                       | \*/\*                          |