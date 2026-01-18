// Helper functions for the OWASP Nest SDK
// These helpers provide additional utility functions for common operations
//
// Copyright (c) 2026 OWASP Nest SDK Contributors

package nest

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/owasp/nest-sdk/models/apierrors"
	"github.com/owasp/nest-sdk/models/operations"
)

// PaginationHelper provides utilities for handling paginated responses
type PaginationHelper struct {
	client *Nest
}

// NewPaginationHelper creates a new pagination helper
func (n *Nest) NewPaginationHelper() *PaginationHelper {
	return &PaginationHelper{client: n}
}

// FetchAllChapters fetches all chapters across all pages
func (p *PaginationHelper) FetchAllChapters(ctx context.Context, req operations.ListChaptersRequest) ([]interface{}, error) {
	var allChapters []interface{}
	page := int64(1)
	pageSize := int64(100) // Use a reasonable page size

	for {
		req.Page = &page
		req.PageSize = &pageSize

		res, err := p.client.Chapters.ListChapters(ctx, req)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch chapters page %d: %w", page, err)
		}

		if res.PagedChapter == nil {
			break
		}

		for _, chapter := range res.PagedChapter.Items {
			allChapters = append(allChapters, chapter)
		}

		// Check if there are more pages
		if !res.PagedChapter.HasNext {
			break
		}

		page++

		// Safety limit to prevent infinite loops
		if page > 1000 {
			return allChapters, fmt.Errorf("reached maximum page limit of 1000")
		}
	}

	return allChapters, nil
}

// ErrorHelper provides utilities for handling SDK errors
type ErrorHelper struct{}

// NewErrorHelper creates a new error helper
func NewErrorHelper() *ErrorHelper {
	return &ErrorHelper{}
}

// IsNotFoundError checks if the error is a 404 not found error
func (e *ErrorHelper) IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	var chapterErr *apierrors.ChapterError
	var committeeErr *apierrors.CommitteeError
	var eventErr *apierrors.EventError
	var issueErr *apierrors.IssueError
	var memberErr *apierrors.MemberError
	var milestoneErr *apierrors.MilestoneError
	var organizationErr *apierrors.OrganizationError
	var projectErr *apierrors.ProjectError
	var releaseErr *apierrors.ReleaseError
	var repositoryErr *apierrors.RepositoryError
	var snapshotErr *apierrors.SnapshotError
	var sponsorErr *apierrors.SponsorError

	return errors.As(err, &chapterErr) ||
		errors.As(err, &committeeErr) ||
		errors.As(err, &eventErr) ||
		errors.As(err, &issueErr) ||
		errors.As(err, &memberErr) ||
		errors.As(err, &milestoneErr) ||
		errors.As(err, &organizationErr) ||
		errors.As(err, &projectErr) ||
		errors.As(err, &releaseErr) ||
		errors.As(err, &repositoryErr) ||
		errors.As(err, &snapshotErr) ||
		errors.As(err, &sponsorErr)
}

// IsAPIError checks if the error is a general API error
func (e *ErrorHelper) IsAPIError(err error) bool {
	if err == nil {
		return false
	}

	var apiErr *apierrors.NestAPIError
	return errors.As(err, &apiErr)
}

// GetStatusCode extracts the HTTP status code from an error
func (e *ErrorHelper) GetStatusCode(err error) int {
	if err == nil {
		return 0
	}

	var apiErr *apierrors.NestAPIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode
	}

	// Check specific error types
	var chapterErr *apierrors.ChapterError
	if errors.As(err, &chapterErr) {
		return chapterErr.HTTPMeta.Response.StatusCode
	}

	var committeeErr *apierrors.CommitteeError
	if errors.As(err, &committeeErr) {
		return committeeErr.HTTPMeta.Response.StatusCode
	}

	var eventErr *apierrors.EventError
	if errors.As(err, &eventErr) {
		return eventErr.HTTPMeta.Response.StatusCode
	}

	var issueErr *apierrors.IssueError
	if errors.As(err, &issueErr) {
		return issueErr.HTTPMeta.Response.StatusCode
	}

	var memberErr *apierrors.MemberError
	if errors.As(err, &memberErr) {
		return memberErr.HTTPMeta.Response.StatusCode
	}

	var milestoneErr *apierrors.MilestoneError
	if errors.As(err, &milestoneErr) {
		return milestoneErr.HTTPMeta.Response.StatusCode
	}

	var organizationErr *apierrors.OrganizationError
	if errors.As(err, &organizationErr) {
		return organizationErr.HTTPMeta.Response.StatusCode
	}

	var projectErr *apierrors.ProjectError
	if errors.As(err, &projectErr) {
		return projectErr.HTTPMeta.Response.StatusCode
	}

	var releaseErr *apierrors.ReleaseError
	if errors.As(err, &releaseErr) {
		return releaseErr.HTTPMeta.Response.StatusCode
	}

	var repositoryErr *apierrors.RepositoryError
	if errors.As(err, &repositoryErr) {
		return repositoryErr.HTTPMeta.Response.StatusCode
	}

	var snapshotErr *apierrors.SnapshotError
	if errors.As(err, &snapshotErr) {
		return snapshotErr.HTTPMeta.Response.StatusCode
	}

	var sponsorErr *apierrors.SponsorError
	if errors.As(err, &sponsorErr) {
		return sponsorErr.HTTPMeta.Response.StatusCode
	}

	return 0
}

// RetryHelper provides utilities for retry operations
type RetryHelper struct{}

// NewRetryHelper creates a new retry helper
func NewRetryHelper() *RetryHelper {
	return &RetryHelper{}
}

// WithExponentialBackoff executes a function with exponential backoff retry logic
func (r *RetryHelper) WithExponentialBackoff(ctx context.Context, maxRetries int, initialDelay time.Duration, fn func() error) error {
	var lastErr error
	delay := initialDelay

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return fmt.Errorf("context cancelled: %w", ctx.Err())
			case <-time.After(delay):
				// Continue with retry
			}
		}

		lastErr = fn()
		if lastErr == nil {
			return nil
		}

		if attempt < maxRetries {
			// Exponential backoff: double the delay each time
			delay *= 2
		}
	}

	return fmt.Errorf("failed after %d attempts: %w", maxRetries+1, lastErr)
}

// ValidateContext checks if the context is valid and not cancelled
func ValidateContext(ctx context.Context) error {
	if ctx == nil {
		return fmt.Errorf("context cannot be nil")
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("context is already cancelled: %w", ctx.Err())
	default:
		return nil
	}
}

// BuildPaginationRequest creates a pagination request with common defaults
func BuildPaginationRequest(page, pageSize int64) (operations.ListChaptersRequest, error) {
	if page < 1 {
		return operations.ListChaptersRequest{}, fmt.Errorf("page must be greater than 0")
	}

	if pageSize < 1 || pageSize > 1000 {
		return operations.ListChaptersRequest{}, fmt.Errorf("pageSize must be between 1 and 1000")
	}

	return operations.ListChaptersRequest{
		Page:     &page,
		PageSize: &pageSize,
	}, nil
}

// StringSliceContains checks if a string slice contains a specific value
func StringSliceContains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// SafeStringDeref safely dereferences a string pointer, returning empty string if nil
func SafeStringDeref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// SafeInt64Deref safely dereferences an int64 pointer, returning 0 if nil
func SafeInt64Deref(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// SafeBoolDeref safely dereferences a bool pointer, returning false if nil
func SafeBoolDeref(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
