# Community
(*Community*)

## Overview

### Available Operations

* [ListMembers](#listmembers) - List members
* [GetMember](#getmember) - Get member by login
* [ListOrganizations](#listorganizations) - List organizations

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
        nest.WithSecurity(os.Getenv("NEST_API_KEY_HEADER")),
    )

    res, err := s.Community.ListMembers(ctx, operations.ListMembersRequest{
        Location: nest.String("India"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedMemberSchema != nil {
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

Retrieve a member by login.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_member" method="get" path="/api/v0/members/{login}" -->
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

    res, err := s.Community.GetMember(ctx, "Stanton97")
    if err != nil {
        log.Fatal(err)
    }
    if res.MemberSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `login`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMemberResponse](../../models/operations/getmemberresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.MemberErrorResponse | 404                           | application/json              |
| apierrors.NestAPIError        | 4XX, 5XX                      | \*/\*                         |

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
        nest.WithSecurity(os.Getenv("NEST_API_KEY_HEADER")),
    )

    res, err := s.Community.ListOrganizations(ctx, nest.String("United States of America"), nil, nest.Int64(1), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedOrganizationSchema != nil {
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
| `page`                                                                                        | **int64*                                                                                      | :heavy_minus_sign:                                                                            | N/A                                                                                           |                                                                                               |
| `pageSize`                                                                                    | **int64*                                                                                      | :heavy_minus_sign:                                                                            | N/A                                                                                           |                                                                                               |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |                                                                                               |

### Response

**[*operations.ListOrganizationsResponse](../../models/operations/listorganizationsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |