package godeeplapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) doRequest(ctx context.Context, method, endpoint string, body interface{}, headers map[string]string) ([]byte, error) {
	var bodyReader io.Reader

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

	// Add auth header
	req.Header.Set("Authorization", "DeepL-Auth-Key "+c.authKey)
	req.Header.Set("Content-Type", "application/json")

	// Add custom headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	c.logger.Debug("Sending request to %s", c.baseURL+endpoint)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
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
