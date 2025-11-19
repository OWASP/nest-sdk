<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	nest "github.com/owasp/nest-sdk"
	"github.com/owasp/nest-sdk/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := nest.New(
		nest.WithSecurity(os.Getenv("NEST_API_KEY")),
	)

	res, err := s.Chapters.ListChapters(ctx, operations.ListChaptersRequest{
		Country: nest.Pointer("India"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.PagedChapter != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->