# Releases
(*Releases*)

## Overview

### Available Operations

* [ListReleases](#listreleases) - List releases
* [GetRelease](#getrelease) - Get release

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
	"github.com/owasp/nest-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := nest.New(
        nest.WithSecurity(os.Getenv("NEST_API_KEY")),
    )

    res, err := s.Releases.ListReleases(ctx, operations.ListReleasesRequest{
        Organization: nest.Pointer("OWASP"),
        Repository: nest.Pointer("Nest"),
        TagName: nest.Pointer("0.2.10"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedRelease != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListReleasesRequest](../../models/operations/listreleasesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListReleasesResponse](../../models/operations/listreleasesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetRelease

Retrieve a specific GitHub release by organization, repository, and tag name.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_release" method="get" path="/api/v0/releases/{organization_id}/{repository_id}/{release_id}" -->
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

    res, err := s.Releases.GetRelease(ctx, "OWASP", "Nest", "0.2.10")
    if err != nil {
        log.Fatal(err)
    }
    if res.ReleaseDetail != nil {
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
| `releaseID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 0.2.10                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetReleaseResponse](../../models/operations/getreleaseresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ReleaseError | 404                    | application/json       |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |