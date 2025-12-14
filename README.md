# nest

Developer-friendly & type-safe Go SDK specifically catered to leverage *nest* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=nest&utm_campaign=go"><img src="https://www.speakeasy.com/assets/badges/built-by-speakeasy.svg" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

<!-- Start Summary [summary] -->
## Summary

OWASP Nest: Open Worldwide Application Security Project API
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [nest](#nest)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/owasp/nest-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type   | Scheme  | Environment Variable |
| -------- | ------ | ------- | -------------------- |
| `APIKey` | apiKey | API key | `NEST_API_KEY`       |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Chapters](docs/sdks/chapters/README.md)

* [ListChapters](docs/sdks/chapters/README.md#listchapters) - List chapters
* [GetChapter](docs/sdks/chapters/README.md#getchapter) - Get chapter

### [Committees](docs/sdks/committees/README.md)

* [ListCommittees](docs/sdks/committees/README.md#listcommittees) - List committees
* [GetCommittee](docs/sdks/committees/README.md#getcommittee) - Get committee

### [Community](docs/sdks/community/README.md)

* [ListMembers](docs/sdks/community/README.md#listmembers) - List members
* [GetMember](docs/sdks/community/README.md#getmember) - Get member
* [ListOrganizations](docs/sdks/community/README.md#listorganizations) - List organizations
* [GetOrganization](docs/sdks/community/README.md#getorganization) - Get organization
* [ListSnapshots](docs/sdks/community/README.md#listsnapshots) - List snapshots
* [GetSnapshot](docs/sdks/community/README.md#getsnapshot) - Get snapshot
* [ListSnapshotChapters](docs/sdks/community/README.md#listsnapshotchapters) - List new chapters in snapshot
* [ListSnapshotIssues](docs/sdks/community/README.md#listsnapshotissues) - List new issues in snapshot
* [ListSnapshotMembers](docs/sdks/community/README.md#listsnapshotmembers) - List new members in snapshot
* [ListSnapshotProjects](docs/sdks/community/README.md#listsnapshotprojects) - List new projects in snapshot
* [ListSnapshotReleases](docs/sdks/community/README.md#listsnapshotreleases) - List new releases in snapshot

### [Events](docs/sdks/events/README.md)

* [ListEvents](docs/sdks/events/README.md#listevents) - List events
* [GetEvent](docs/sdks/events/README.md#getevent) - Get event

### [Issues](docs/sdks/issues/README.md)

* [ListIssues](docs/sdks/issues/README.md#listissues) - List issues
* [GetIssue](docs/sdks/issues/README.md#getissue) - Get issue

### [Milestones](docs/sdks/milestones/README.md)

* [ListMilestones](docs/sdks/milestones/README.md#listmilestones) - List milestones
* [GetMilestone](docs/sdks/milestones/README.md#getmilestone) - Get milestone

### [Projects](docs/sdks/projects/README.md)

* [ListProjects](docs/sdks/projects/README.md#listprojects) - List projects
* [GetProject](docs/sdks/projects/README.md#getproject) - Get project

### [Releases](docs/sdks/releases/README.md)

* [ListReleases](docs/sdks/releases/README.md#listreleases) - List releases
* [GetRelease](docs/sdks/releases/README.md#getrelease) - Get release

### [Repositories](docs/sdks/repositories/README.md)

* [ListRepositories](docs/sdks/repositories/README.md#listrepositories) - List repositories
* [GetRepository](docs/sdks/repositories/README.md#getrepository) - Get repository

### [Sponsors](docs/sdks/sponsors/README.md)

* [ListSponsors](docs/sdks/sponsors/README.md#listsponsors) - List sponsors
* [GetSponsor](docs/sdks/sponsors/README.md#getsponsor) - Get sponsor

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	nest "github.com/owasp/nest-sdk"
	"github.com/owasp/nest-sdk/models/operations"
	"github.com/owasp/nest-sdk/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := nest.New(
		nest.WithSecurity(os.Getenv("NEST_API_KEY")),
	)

	res, err := s.Chapters.ListChapters(ctx, operations.ListChaptersRequest{
		Country: nest.Pointer("India"),
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.PagedChapter != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	nest "github.com/owasp/nest-sdk"
	"github.com/owasp/nest-sdk/models/operations"
	"github.com/owasp/nest-sdk/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := nest.New(
		nest.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.NestAPIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetChapter` function may return the following errors:

| Error Type             | Status Code | Content Type     |
| ---------------------- | ----------- | ---------------- |
| apierrors.ChapterError | 404         | application/json |
| apierrors.NestAPIError | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	nest "github.com/owasp/nest-sdk"
	"github.com/owasp/nest-sdk/models/apierrors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := nest.New(
		nest.WithSecurity(os.Getenv("NEST_API_KEY")),
	)

	res, err := s.Chapters.GetChapter(ctx, "London")
	if err != nil {

		var e *apierrors.ChapterError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.NestAPIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
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
		nest.WithServerURL("https://nest.owasp.org"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/owasp/nest-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = nest.New(nest.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation.
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release.

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=nest&utm_campaign=go)
