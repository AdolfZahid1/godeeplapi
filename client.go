package godeeplapi

import (
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	authKey    string
	httpClient *http.Client
	logger     Logger
}

type ClientOption func(*Client)

// WithLogger sets a custom logger
func WithLogger(logger Logger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithTimeout sets a custom timeout for the HTTP client
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// NewClient creates a new DeepL API client
func NewClient(apiKey string, isPro bool, opts ...ClientOption) *Client {
	baseURL := "https://api-free.deepl.com/v2"
	if isPro {
		baseURL = "https://api.deepl.com/v2"
	}

	client := &Client{
		baseURL:    baseURL,
		authKey:    apiKey,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		logger:     &defaultLogger{},
	}

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	return client
}
