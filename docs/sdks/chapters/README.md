# Chapters
(*Chapters*)

## Overview

### Available Operations

* [AppsAPIRestV0ChapterListChapters](#appsapirestv0chapterlistchapters) - List chapters
* [AppsAPIRestV0ChapterGetChapter](#appsapirestv0chaptergetchapter) - Get chapter

## AppsAPIRestV0ChapterListChapters

Retrieve a paginated list of OWASP chapters.

### Example Usage

<!-- UsageSnippet language="go" operationID="apps_api_rest_v0_chapter_list_chapters" method="get" path="/api/v0/chapters/" -->
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

    res, err := s.Chapters.AppsAPIRestV0ChapterListChapters(ctx, operations.AppsAPIRestV0ChapterListChaptersRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PagedChapterSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.AppsAPIRestV0ChapterListChaptersRequest](../../models/operations/appsapirestv0chapterlistchaptersrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.AppsAPIRestV0ChapterListChaptersResponse](../../models/operations/appsapirestv0chapterlistchaptersresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.NestAPIError | 4XX, 5XX               | \*/\*                  |

## AppsAPIRestV0ChapterGetChapter

Retrieve chapter details.

### Example Usage

<!-- UsageSnippet language="go" operationID="apps_api_rest_v0_chapter_get_chapter" method="get" path="/api/v0/chapters/{chapter_id}" -->
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

    res, err := s.Chapters.AppsAPIRestV0ChapterGetChapter(ctx, "London")
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

**[*operations.AppsAPIRestV0ChapterGetChapterResponse](../../models/operations/appsapirestv0chaptergetchapterresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ChapterErrorResponse | 404                            | application/json               |
| apierrors.NestAPIError         | 4XX, 5XX                       | \*/\*                          |