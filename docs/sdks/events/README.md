# Events
(*Events*)

## Overview

### Available Operations

* [ListEvents](#listevents) - List events
* [GetEvent](#getevent) - Get event

## ListEvents

Retrieve a paginated list of OWASP events.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_events" method="get" path="/api/v0/events/" -->
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

    res, err := s.Events.ListEvents(ctx, operations.ListEventsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListEventsRequest](../../models/operations/listeventsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListEventsResponse](../../models/operations/listeventsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetEvent

Retrieve an event details.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_event" method="get" path="/api/v0/events/{event_id}" -->
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

    res, err := s.Events.GetEvent(ctx, "owasp-global-appsec-usa-2025-washington-dc")
    if err != nil {
        log.Fatal(err)
    }
    if res.EventDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `eventID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | owasp-global-appsec-usa-2025-washington-dc               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetEventResponse](../../models/operations/geteventresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.EventError   | 404                    | application/json       |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |