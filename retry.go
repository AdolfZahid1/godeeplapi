package godeeplapi

import (
	"context"
	"fmt"
	"net"
	"time"
)

// RetryableError checks if an error is retryable
func isRetryableError(err error) bool {
	if err == nil {
		return false
	}

	if apiErr, ok := err.(*APIError); ok {
		// Retry on rate limiting and server errors
		return apiErr.StatusCode == 429 || apiErr.StatusCode >= 500
	}

	// Check for network errors
	if netErr, ok := err.(net.Error); ok {
		return netErr.Temporary() || netErr.Timeout()
	}

	return false
}

// RetryableClient wraps an HTTP client with retry capabilities
func (c *Client) retryRequest(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	var lastErr error

	// Configure retry parameters
	maxRetries := 3
	initialBackoff := 1 * time.Second
	maxBackoff := 30 * time.Second

	for i := 0; i < maxRetries; i++ {
		result, err := fn()
		if err == nil {
			return result, nil
		}

		lastErr = err

		// Check if error is retryable
		if !isRetryableError(err) {
			return nil, err
		}

		// Check for context cancellation
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			// Calculate backoff with exponential increase
			backoff := initialBackoff * (1 << uint(i))
			if backoff > maxBackoff {
				backoff = maxBackoff
			}

			c.logger.Info("Retrying request after error: %v (attempt %d/%d, waiting %v)",
				err, i+1, maxRetries, backoff)

			// Wait with context awareness
			select {
			case <-time.After(backoff):
				// Continue with retry
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
	}

	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}
