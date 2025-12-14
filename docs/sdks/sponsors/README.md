# Sponsors

## Overview

### Available Operations

* [ListSponsors](#listsponsors) - List sponsors
* [GetSponsor](#getsponsor) - Get sponsor

## ListSponsors

Retrieve a paginated list of OWASP sponsors.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_sponsors" method="get" path="/api/v0/sponsors/" -->
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

    res, err := s.Sponsors.ListSponsors(ctx, operations.ListSponsorsRequest{
        SponsorType: nest.Pointer("Silver"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedSponsor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListSponsorsRequest](../../models/operations/listsponsorsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListSponsorsResponse](../../models/operations/listsponsorsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetSponsor

Retrieve a sponsor details.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_sponsor" method="get" path="/api/v0/sponsors/{sponsor_id}" -->
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

    res, err := s.Sponsors.GetSponsor(ctx, "adobe")
    if err != nil {
        log.Fatal(err)
    }
    if res.SponsorDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `sponsorID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | adobe                                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetSponsorResponse](../../models/operations/getsponsorresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.SponsorError | 404                    | application/json       |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |