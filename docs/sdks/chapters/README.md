# Chapters
(*Chapters*)

## Overview

### Available Operations

* [ListChapters](#listchapters) - List chapters
* [GetChapter](#getchapter) - Get chapter

## ListChapters

Retrieve a paginated list of OWASP chapters.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_chapters" method="get" path="/api/v0/chapters/" -->
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

    res, err := s.Chapters.ListChapters(ctx, operations.ListChaptersRequest{
        Country: nest.Pointer("India"),
        Region: nest.Pointer("Asia"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedChapterSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListChaptersRequest](../../models/operations/listchaptersrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListChaptersResponse](../../models/operations/listchaptersresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## GetChapter

Retrieve chapter details.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_chapter" method="get" path="/api/v0/chapters/{chapter_id}" -->
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

    res, err := s.Chapters.GetChapter(ctx, "London")
    if err != nil {
        log.Fatal(err)
    }
    if res.ChapterSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `chapterID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | London                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetChapterResponse](../../models/operations/getchapterresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ChapterErrorResponse | 404                            | application/json               |
| apierrors.NestAPIError         | 4XX, 5XX                       | \*/\*                          |