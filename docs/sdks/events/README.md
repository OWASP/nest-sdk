# Events
(*Events*)

## Overview

### Available Operations

* [ListEvents](#listevents) - List events

## ListEvents

Retrieve a paginated list of OWASP events.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_events" method="get" path="/api/v1/events/" -->
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

    res, err := s.Events.ListEvents(ctx, nil, nest.Int64(1), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedEventSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |
| `ordering`                                                                      | [*operations.ListEventsOrdering](../../models/operations/listeventsordering.md) | :heavy_minus_sign:                                                              | Ordering field                                                                  |
| `page`                                                                          | **int64*                                                                        | :heavy_minus_sign:                                                              | N/A                                                                             |
| `pageSize`                                                                      | **int64*                                                                        | :heavy_minus_sign:                                                              | N/A                                                                             |
| `opts`                                                                          | [][operations.Option](../../models/operations/option.md)                        | :heavy_minus_sign:                                                              | The options for this request.                                                   |

### Response

**[*operations.ListEventsResponse](../../models/operations/listeventsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |