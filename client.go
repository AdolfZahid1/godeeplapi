package godeeplapi

import (
	"fmt"
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

// NewClient creates a new DeepL API client for v2 API
func NewClient(apiKey string, isPro bool, opts ...ClientOption) *Client {
	return newClientWithVersion(apiKey, isPro, "v2", opts...)
}

// NewClientV3 creates a new DeepL API client for v3 API
func NewClientV3(apiKey string, isPro bool, opts ...ClientOption) *Client {
	return newClientWithVersion(apiKey, isPro, "v3", opts...)
}

// newClientWithVersion is a helper function to create a client with a specific API version
func newClientWithVersion(apiKey string, isPro bool, version string, opts ...ClientOption) *Client {
	baseURLPrefix := "https://api-free.deepl.com"
	if isPro {
		baseURLPrefix = "https://api.deepl.com"
	}

	client := &Client{
		baseURL:    fmt.Sprintf("%s/%s", baseURLPrefix, version),
		authKey:    apiKey,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		logger:     NewDefaultLogger(), // Initialize with our new function
	}

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	return client
}
