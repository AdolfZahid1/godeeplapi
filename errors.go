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
	ErrBadRequest    = &APIError{StatusCode: 400, Message: "Bad Request"}
	ErrInvalidAuth   = &APIError{StatusCode: 401, Message: "Authentication failed"}
	ErrForbidden     = &APIError{StatusCode: 403, Message: "Forbidden. Insufficient access rights."}
	ErrNotFound      = &APIError{StatusCode: 404, Message: "The requested resource could not be found"}
	ErrRateLimit     = &APIError{StatusCode: 413, Message: "Rate limit exceeded"}
	ErrHeader        = &APIError{StatusCode: 415, Message: "The requested entries format specified in \"Accept\" is not supported."}
	Err429TooMany    = &APIError{StatusCode: 429, Message: "Too many requests. Please wait and resend request"}
	ErrQuotaExceeded = &APIError{StatusCode: 456, Message: "Quota exceeded"}
	ErrInternal      = &APIError{StatusCode: 500, Message: "Internal error"}
	ErrUnavailable   = &APIError{StatusCode: 503, Message: "Resource temporarily unavailable. Try again later."}
	Err529TooMany    = &APIError{StatusCode: 529, Message: "Too many requests. Please wait and resend request"}
)
