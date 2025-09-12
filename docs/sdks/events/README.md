# Events
(*Events*)

## Overview

### Available Operations

* [AppsAPIRestV0EventListEvents](#appsapirestv0eventlistevents) - List events

## AppsAPIRestV0EventListEvents

Retrieve a paginated list of OWASP events.

### Example Usage

<!-- UsageSnippet language="go" operationID="apps_api_rest_v0_event_list_events" method="get" path="/api/v0/events/" -->
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

    res, err := s.Events.AppsAPIRestV0EventListEvents(ctx, nil, nest.Int64(1), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedEventSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                           | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                               | :heavy_check_mark:                                                                                                  | The context to use for the request.                                                                                 |
| `ordering`                                                                                                          | [*operations.AppsAPIRestV0EventListEventsOrdering](../../models/operations/appsapirestv0eventlisteventsordering.md) | :heavy_minus_sign:                                                                                                  | Ordering field                                                                                                      |
| `page`                                                                                                              | **int64*                                                                                                            | :heavy_minus_sign:                                                                                                  | N/A                                                                                                                 |
| `pageSize`                                                                                                          | **int64*                                                                                                            | :heavy_minus_sign:                                                                                                  | N/A                                                                                                                 |
| `opts`                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                            | :heavy_minus_sign:                                                                                                  | The options for this request.                                                                                       |

### Response

**[*operations.AppsAPIRestV0EventListEventsResponse](../../models/operations/appsapirestv0eventlisteventsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |