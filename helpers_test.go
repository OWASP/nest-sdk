// Helper functions tests for the OWASP Nest SDK
//
// Copyright (c) 2026 OWASP Nest SDK Contributors

package nest

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestValidateContext(t *testing.T) {
	t.Run("valid context", func(t *testing.T) {
		ctx := context.Background()
		err := ValidateContext(ctx)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("nil context", func(t *testing.T) {
		err := ValidateContext(context.TODO())
		if err != nil {
			t.Errorf("expected no error for context.TODO, got %v", err)
		}
	})

	t.Run("truly nil context", func(t *testing.T) {
		var ctx context.Context
		err := ValidateContext(ctx)
		if err == nil {
			t.Error("expected error for nil context, got nil")
		}
	})

	t.Run("cancelled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		err := ValidateContext(ctx)
		if err == nil {
			t.Error("expected error for cancelled context, got nil")
		}
	})
}

func TestBuildPaginationRequest(t *testing.T) {
	t.Run("valid pagination", func(t *testing.T) {
		req, err := BuildPaginationRequest(1, 50)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if req.Page == nil || *req.Page != 1 {
			t.Error("expected page to be 1")
		}
		if req.PageSize == nil || *req.PageSize != 50 {
			t.Error("expected pageSize to be 50")
		}
	})

	t.Run("invalid page number", func(t *testing.T) {
		_, err := BuildPaginationRequest(0, 50)
		if err == nil {
			t.Error("expected error for page 0, got nil")
		}
	})

	t.Run("invalid page size - too small", func(t *testing.T) {
		_, err := BuildPaginationRequest(1, 0)
		if err == nil {
			t.Error("expected error for page size 0, got nil")
		}
	})

	t.Run("invalid page size - too large", func(t *testing.T) {
		_, err := BuildPaginationRequest(1, 1001)
		if err == nil {
			t.Error("expected error for page size > 1000, got nil")
		}
	})

	t.Run("boundary page size", func(t *testing.T) {
		req, err := BuildPaginationRequest(1, 1000)
		if err != nil {
			t.Errorf("expected no error for page size 1000, got %v", err)
		}
		if req.PageSize == nil || *req.PageSize != 1000 {
			t.Error("expected pageSize to be 1000")
		}
	})
}

func TestStringSliceContains(t *testing.T) {
	t.Run("contains value", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		if !StringSliceContains(slice, "banana") {
			t.Error("expected to find 'banana' in slice")
		}
	})

	t.Run("does not contain value", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		if StringSliceContains(slice, "orange") {
			t.Error("expected not to find 'orange' in slice")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []string{}
		if StringSliceContains(slice, "test") {
			t.Error("expected not to find value in empty slice")
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var slice []string
		if StringSliceContains(slice, "test") {
			t.Error("expected not to find value in nil slice")
		}
	})
}

func TestSafeStringDeref(t *testing.T) {
	t.Run("valid string pointer", func(t *testing.T) {
		str := "test"
		result := SafeStringDeref(&str)
		if result != "test" {
			t.Errorf("expected 'test', got '%s'", result)
		}
	})

	t.Run("nil string pointer", func(t *testing.T) {
		result := SafeStringDeref(nil)
		if result != "" {
			t.Errorf("expected empty string, got '%s'", result)
		}
	})

	t.Run("empty string pointer", func(t *testing.T) {
		str := ""
		result := SafeStringDeref(&str)
		if result != "" {
			t.Errorf("expected empty string, got '%s'", result)
		}
	})
}

func TestSafeInt64Deref(t *testing.T) {
	t.Run("valid int64 pointer", func(t *testing.T) {
		num := int64(42)
		result := SafeInt64Deref(&num)
		if result != 42 {
			t.Errorf("expected 42, got %d", result)
		}
	})

	t.Run("nil int64 pointer", func(t *testing.T) {
		result := SafeInt64Deref(nil)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})

	t.Run("zero int64 pointer", func(t *testing.T) {
		num := int64(0)
		result := SafeInt64Deref(&num)
		if result != 0 {
			t.Errorf("expected 0, got %d", result)
		}
	})

	t.Run("negative int64 pointer", func(t *testing.T) {
		num := int64(-100)
		result := SafeInt64Deref(&num)
		if result != -100 {
			t.Errorf("expected -100, got %d", result)
		}
	})
}

func TestSafeBoolDeref(t *testing.T) {
	t.Run("valid true bool pointer", func(t *testing.T) {
		b := true
		result := SafeBoolDeref(&b)
		if !result {
			t.Error("expected true, got false")
		}
	})

	t.Run("valid false bool pointer", func(t *testing.T) {
		b := false
		result := SafeBoolDeref(&b)
		if result {
			t.Error("expected false, got true")
		}
	})

	t.Run("nil bool pointer", func(t *testing.T) {
		result := SafeBoolDeref(nil)
		if result {
			t.Error("expected false, got true")
		}
	})
}

func TestErrorHelper_IsAPIError(t *testing.T) {
	helper := NewErrorHelper()

	t.Run("regular error", func(t *testing.T) {
		err := errors.New("regular error")
		if helper.IsAPIError(err) {
			t.Error("expected false for regular error")
		}
	})

	t.Run("nil error", func(t *testing.T) {
		if helper.IsAPIError(nil) {
			t.Error("expected false for nil error")
		}
	})
}

func TestErrorHelper_GetStatusCode(t *testing.T) {
	helper := NewErrorHelper()

	t.Run("nil error", func(t *testing.T) {
		code := helper.GetStatusCode(nil)
		if code != 0 {
			t.Errorf("expected 0 for nil error, got %d", code)
		}
	})

	t.Run("regular error", func(t *testing.T) {
		err := errors.New("regular error")
		code := helper.GetStatusCode(err)
		if code != 0 {
			t.Errorf("expected 0 for regular error, got %d", code)
		}
	})
}

func TestRetryHelper_WithExponentialBackoff(t *testing.T) {
	helper := NewRetryHelper()

	t.Run("succeeds on first attempt", func(t *testing.T) {
		ctx := context.Background()
		attempts := 0

		fn := func() error {
			attempts++
			return nil
		}

		err := helper.WithExponentialBackoff(ctx, 3, 10*time.Millisecond, fn)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if attempts != 1 {
			t.Errorf("expected 1 attempt, got %d", attempts)
		}
	})

	t.Run("succeeds on retry", func(t *testing.T) {
		ctx := context.Background()
		attempts := 0

		fn := func() error {
			attempts++
			if attempts < 3 {
				return errors.New("temporary error")
			}
			return nil
		}

		err := helper.WithExponentialBackoff(ctx, 5, 10*time.Millisecond, fn)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if attempts != 3 {
			t.Errorf("expected 3 attempts, got %d", attempts)
		}
	})

	t.Run("fails after max retries", func(t *testing.T) {
		ctx := context.Background()
		attempts := 0

		fn := func() error {
			attempts++
			return errors.New("persistent error")
		}

		err := helper.WithExponentialBackoff(ctx, 3, 10*time.Millisecond, fn)
		if err == nil {
			t.Error("expected error after max retries, got nil")
		}
		if attempts != 4 { // maxRetries + 1
			t.Errorf("expected 4 attempts, got %d", attempts)
		}
	})

	t.Run("context cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		attempts := 0

		fn := func() error {
			attempts++
			if attempts == 2 {
				cancel() // Cancel context on second attempt
			}
			return errors.New("error")
		}

		err := helper.WithExponentialBackoff(ctx, 5, 10*time.Millisecond, fn)
		if err == nil {
			t.Error("expected error due to context cancellation, got nil")
		}
		// Should fail on the retry after cancellation
		if attempts < 2 {
			t.Errorf("expected at least 2 attempts, got %d", attempts)
		}
	})

	t.Run("exponential delay increases", func(t *testing.T) {
		ctx := context.Background()
		attempts := 0
		startTime := time.Now()

		fn := func() error {
			attempts++
			if attempts <= 3 {
				return errors.New("temporary error")
			}
			return nil
		}

		initialDelay := 50 * time.Millisecond
		err := helper.WithExponentialBackoff(ctx, 5, initialDelay, fn)
		elapsed := time.Since(startTime)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Expected delays: 50ms, 100ms, 200ms = 350ms minimum
		// Add some buffer for execution time
		minExpected := initialDelay + (initialDelay * 2) + (initialDelay * 4)
		if elapsed < minExpected {
			t.Errorf("expected at least %v delay, got %v", minExpected, elapsed)
		}
	})
}

func TestNewErrorHelper(t *testing.T) {
	helper := NewErrorHelper()
	if helper == nil {
		t.Error("expected non-nil ErrorHelper")
	}
}

func TestNewRetryHelper(t *testing.T) {
	helper := NewRetryHelper()
	if helper == nil {
		t.Error("expected non-nil RetryHelper")
	}
}
