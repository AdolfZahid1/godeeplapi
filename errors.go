package godeeplapi

import "fmt"

type APIError struct {
	StatusCode int
	Message    string
	Details    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("DeepL API error (%d): %s - %s", e.StatusCode, e.Message, e.Details)
}

// Common error types for more specific handling
var (
	ErrInvalidAuth   = &APIError{StatusCode: 403, Message: "Authentication failed"}
	ErrRateLimit     = &APIError{StatusCode: 429, Message: "Rate limit exceeded"}
	ErrQuotaExceeded = &APIError{StatusCode: 456, Message: "Quota exceeded"}
)
