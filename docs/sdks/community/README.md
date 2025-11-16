# Community
(*Community*)

## Overview

### Available Operations

* [ListMembers](#listmembers) - List members
* [GetMember](#getmember) - Get member
* [ListOrganizations](#listorganizations) - List organizations
* [GetOrganization](#getorganization) - Get organization
* [ListSnapshots](#listsnapshots) - List snapshots
* [GetSnapshot](#getsnapshot) - Get snapshot
* [ListSnapshotChapters](#listsnapshotchapters) - List new chapters in snapshot
* [ListSnapshotIssues](#listsnapshotissues) - List new issues in snapshot
* [ListSnapshotMembers](#listsnapshotmembers) - List new members in snapshot
* [ListSnapshotProjects](#listsnapshotprojects) - List new projects in snapshot
* [ListSnapshotReleases](#listsnapshotreleases) - List new releases in snapshot

## ListMembers

Retrieve a paginated list of OWASP community members.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_members" method="get" path="/api/v0/members/" -->
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

    res, err := s.Community.ListMembers(ctx, operations.ListMembersRequest{
        Location: nest.Pointer("India"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListMembersRequest](../../models/operations/listmembersrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListMembersResponse](../../models/operations/listmembersresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetMember

Retrieve member details.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_member" method="get" path="/api/v0/members/{member_id}" -->
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

    res, err := s.Community.GetMember(ctx, "OWASP")
    if err != nil {
        log.Fatal(err)
    }
    if res.MemberDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `memberID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | OWASP                                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMemberResponse](../../models/operations/getmemberresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.MemberError  | 404                    | application/json       |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## ListOrganizations

Retrieve a paginated list of GitHub organizations.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_organizations" method="get" path="/api/v0/organizations/" -->
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

    res, err := s.Community.ListOrganizations(ctx, nest.Pointer("United States of America"), nil, nest.Pointer[int64](1), nest.Pointer[int64](100))
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |                                                                                               |
| `location`                                                                                    | **string*                                                                                     | :heavy_minus_sign:                                                                            | Location of the organization                                                                  | United States of America                                                                      |
| `ordering`                                                                                    | [*operations.ListOrganizationsOrdering](../../models/operations/listorganizationsordering.md) | :heavy_minus_sign:                                                                            | Ordering field                                                                                |                                                                                               |
| `page`                                                                                        | **int64*                                                                                      | :heavy_minus_sign:                                                                            | Page number                                                                                   |                                                                                               |
| `pageSize`                                                                                    | **int64*                                                                                      | :heavy_minus_sign:                                                                            | Number of items per page                                                                      |                                                                                               |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |                                                                                               |

### Response

**[*operations.ListOrganizationsResponse](../../models/operations/listorganizationsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetOrganization

Retrieve project details.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_organization" method="get" path="/api/v0/organizations/{organization_id}" -->
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

    res, err := s.Community.GetOrganization(ctx, "OWASP")
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `organizationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | OWASP                                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetOrganizationResponse](../../models/operations/getorganizationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| apierrors.OrganizationError | 404                         | application/json            |
| apierrors.NestAPIError      | 4XX, 5XX                    | \*/\*                       |

## ListSnapshots

Retrieve a paginated list of OWASP snapshots.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_snapshots" method="get" path="/api/v0/snapshots/" -->
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

    res, err := s.Community.ListSnapshots(ctx, nil, nest.Pointer[int64](1), nest.Pointer[int64](100))
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedSnapshot != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `ordering`                                                                            | [*operations.ListSnapshotsOrdering](../../models/operations/listsnapshotsordering.md) | :heavy_minus_sign:                                                                    | Ordering field                                                                        |
| `page`                                                                                | **int64*                                                                              | :heavy_minus_sign:                                                                    | Page number                                                                           |
| `pageSize`                                                                            | **int64*                                                                              | :heavy_minus_sign:                                                                    | Number of items per page                                                              |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

### Response

**[*operations.ListSnapshotsResponse](../../models/operations/listsnapshotsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetSnapshot

Retrieve snapshot details.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_snapshot" method="get" path="/api/v0/snapshots/{snapshot_id}" -->
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

    res, err := s.Community.GetSnapshot(ctx, "2025-02")
    if err != nil {
        log.Fatal(err)
    }
    if res.SnapshotDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `snapshotID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2025-02                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetSnapshotResponse](../../models/operations/getsnapshotresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.SnapshotError | 404                     | application/json        |
| apierrors.NestAPIError  | 4XX, 5XX                | \*/\*                   |

## ListSnapshotChapters

Retrieve a paginated list of new chapters in a snapshot.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_snapshot_chapters" method="get" path="/api/v0/snapshots/{snapshot_id}/chapters/" -->
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

    res, err := s.Community.ListSnapshotChapters(ctx, "2025-02", nil, nest.Pointer[int64](1), nest.Pointer[int64](100))
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedChapter != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `snapshotID`                                                                                        | *string*                                                                                            | :heavy_check_mark:                                                                                  | N/A                                                                                                 | 2025-02                                                                                             |
| `ordering`                                                                                          | [*operations.ListSnapshotChaptersOrdering](../../models/operations/listsnapshotchaptersordering.md) | :heavy_minus_sign:                                                                                  | Ordering field                                                                                      |                                                                                                     |
| `page`                                                                                              | **int64*                                                                                            | :heavy_minus_sign:                                                                                  | Page number                                                                                         |                                                                                                     |
| `pageSize`                                                                                          | **int64*                                                                                            | :heavy_minus_sign:                                                                                  | Number of items per page                                                                            |                                                                                                     |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.ListSnapshotChaptersResponse](../../models/operations/listsnapshotchaptersresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## ListSnapshotIssues

Retrieve a paginated list of new issues in a snapshot.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_snapshot_issues" method="get" path="/api/v0/snapshots/{snapshot_id}/issues/" -->
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

    res, err := s.Community.ListSnapshotIssues(ctx, "2025-02", nil, nest.Pointer[int64](1), nest.Pointer[int64](100))
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedSnapshotIssue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     | Example                                                                                         |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |                                                                                                 |
| `snapshotID`                                                                                    | *string*                                                                                        | :heavy_check_mark:                                                                              | N/A                                                                                             | 2025-02                                                                                         |
| `ordering`                                                                                      | [*operations.ListSnapshotIssuesOrdering](../../models/operations/listsnapshotissuesordering.md) | :heavy_minus_sign:                                                                              | Ordering field                                                                                  |                                                                                                 |
| `page`                                                                                          | **int64*                                                                                        | :heavy_minus_sign:                                                                              | Page number                                                                                     |                                                                                                 |
| `pageSize`                                                                                      | **int64*                                                                                        | :heavy_minus_sign:                                                                              | Number of items per page                                                                        |                                                                                                 |
| `opts`                                                                                          | [][operations.Option](../../models/operations/option.md)                                        | :heavy_minus_sign:                                                                              | The options for this request.                                                                   |                                                                                                 |

### Response

**[*operations.ListSnapshotIssuesResponse](../../models/operations/listsnapshotissuesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## ListSnapshotMembers

Retrieve a paginated list of new members in a snapshot.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_snapshot_members" method="get" path="/api/v0/snapshots/{snapshot_id}/members/" -->
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

    res, err := s.Community.ListSnapshotMembers(ctx, "2025-02", nil, nest.Pointer[int64](1), nest.Pointer[int64](100))
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |                                                                                                   |
| `snapshotID`                                                                                      | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               | 2025-02                                                                                           |
| `ordering`                                                                                        | [*operations.ListSnapshotMembersOrdering](../../models/operations/listsnapshotmembersordering.md) | :heavy_minus_sign:                                                                                | Ordering field                                                                                    |                                                                                                   |
| `page`                                                                                            | **int64*                                                                                          | :heavy_minus_sign:                                                                                | Page number                                                                                       |                                                                                                   |
| `pageSize`                                                                                        | **int64*                                                                                          | :heavy_minus_sign:                                                                                | Number of items per page                                                                          |                                                                                                   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |                                                                                                   |

### Response

**[*operations.ListSnapshotMembersResponse](../../models/operations/listsnapshotmembersresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## ListSnapshotProjects

Retrieve a paginated list of new projects in a snapshot.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_snapshot_projects" method="get" path="/api/v0/snapshots/{snapshot_id}/projects/" -->
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

    res, err := s.Community.ListSnapshotProjects(ctx, "2025-02", nil, nest.Pointer[int64](1), nest.Pointer[int64](100))
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedProject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `snapshotID`                                                                                        | *string*                                                                                            | :heavy_check_mark:                                                                                  | N/A                                                                                                 | 2025-02                                                                                             |
| `ordering`                                                                                          | [*operations.ListSnapshotProjectsOrdering](../../models/operations/listsnapshotprojectsordering.md) | :heavy_minus_sign:                                                                                  | Ordering field                                                                                      |                                                                                                     |
| `page`                                                                                              | **int64*                                                                                            | :heavy_minus_sign:                                                                                  | Page number                                                                                         |                                                                                                     |
| `pageSize`                                                                                          | **int64*                                                                                            | :heavy_minus_sign:                                                                                  | Number of items per page                                                                            |                                                                                                     |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.ListSnapshotProjectsResponse](../../models/operations/listsnapshotprojectsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## ListSnapshotReleases

Retrieve a paginated list of new releases in a snapshot.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_snapshot_releases" method="get" path="/api/v0/snapshots/{snapshot_id}/releases/" -->
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

    res, err := s.Community.ListSnapshotReleases(ctx, "2025-02", nil, nest.Pointer[int64](1), nest.Pointer[int64](100))
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedSnapshotRelease != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `snapshotID`                                                                                        | *string*                                                                                            | :heavy_check_mark:                                                                                  | N/A                                                                                                 | 2025-02                                                                                             |
| `ordering`                                                                                          | [*operations.ListSnapshotReleasesOrdering](../../models/operations/listsnapshotreleasesordering.md) | :heavy_minus_sign:                                                                                  | Ordering field                                                                                      |                                                                                                     |
| `page`                                                                                              | **int64*                                                                                            | :heavy_minus_sign:                                                                                  | Page number                                                                                         |                                                                                                     |
| `pageSize`                                                                                          | **int64*                                                                                            | :heavy_minus_sign:                                                                                  | Number of items per page                                                                            |                                                                                                     |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.ListSnapshotReleasesResponse](../../models/operations/listsnapshotreleasesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |