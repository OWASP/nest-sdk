<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	nest "github.com/owasp/nest-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := nest.New(
		nest.WithSecurity(os.Getenv("NEST_API_KEY")),
	)

	res, err := s.Chapters.ListChapters(ctx, nest.Pointer("India"), nil, nest.Pointer[int64](1), nest.Pointer[int64](100))
	if err != nil {
		log.Fatal(err)
	}
	if res.PagedChapter != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->