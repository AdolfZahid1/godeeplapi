package godeeplapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
)

// checkAuth verifies if the API key is set
func (c *Client) checkAuth() error {
	if c == nil {
		return errors.New("client is nil")
	}
	if c.authKey == "" {
		return fmt.Errorf("DeepL API token is empty")
	}
	return nil
}

// unmarshalResponse handles JSON unmarshaling with error handling
func unmarshalResponse(respBody []byte, target interface{}) error {
	if err := json.Unmarshal(respBody, target); err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}
	return nil
}

func (c *Client) doRequest(ctx context.Context, method, endpoint string, body interface{}, headers map[string]string) ([]byte, error) {
	return c.doRequestWithQuery(ctx, method, endpoint, body, headers, nil)
}

func (c *Client) doRequestWithQuery(ctx context.Context, method, endpoint string, body interface{}, headers map[string]string, queryParams interface{}) ([]byte, error) {
	var bodyReader io.Reader
	if c == nil {
		return nil, errors.New("client is nil")
	}
	// Ensure logger is initialized
	if c.logger == nil {
		c.logger = NewDefaultLogger()
	}

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+endpoint, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add query parameters
	if queryParams != nil {
		q := req.URL.Query()

		// Use reflection to extract struct fields and their values
		v := reflect.ValueOf(queryParams)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if v.Kind() == reflect.Struct {
			t := v.Type()
			for i := 0; i < v.NumField(); i++ {
				field := t.Field(i)
				value := v.Field(i)

				// Get the query parameter name from the json tag, or use field name
				paramName := field.Name
				if jsonTag := field.Tag.Get("json"); jsonTag != "" {
					parts := strings.Split(jsonTag, ",")
					if parts[0] != "-" {
						paramName = parts[0]
					}
				}

				// Only add non-empty values to query
				if !value.IsZero() {
					q.Add(paramName, fmt.Sprintf("%v", value.Interface()))
				}
			}
		}

		req.URL.RawQuery = q.Encode()
	}

	// Add auth header
	req.Header.Set("Authorization", "DeepL-Auth-Key "+c.authKey)
	req.Header.Set("Content-Type", "application/json")

	// Add custom headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	c.logger.Debug("Sending request to %s", req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if !isSuccessStatus(resp.StatusCode) {
		var errMsg string
		if len(respBody) > 0 {
			errMsg = string(respBody)
		}

		apiErr := &APIError{
			StatusCode: resp.StatusCode,
			Message:    http.StatusText(resp.StatusCode),
			Details:    errMsg,
		}

		// Log error
		c.logger.Error("API error: %s", apiErr.Error())

		return nil, apiErr
	}

	return respBody, nil
}

// Helper function to check if status code indicates success
func isSuccessStatus(statusCode int) bool {
	switch statusCode {
	case http.StatusOK, http.StatusAccepted, http.StatusCreated, http.StatusNoContent:
		return true
	default:
		return false
	}
}
