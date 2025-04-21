package godeeplapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AdolfZahid1/godeeplapi/models"
)

// ImproveText improves a text using the DeepL API
func (c *Client) ImproveText(ctx context.Context, req models.RephraseRequest) (string, error) {
	if c.authKey == "" {
		return "", fmt.Errorf("DeepL API token is empty")
	}

	endpoint := "/write/rephrase"

	respBody, err := c.doRequest(ctx, "POST", endpoint, req, nil)
	if err != nil {
		return "", err
	}

	var response models.RephraseResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %w", err)
	}

	if response.Text == "" {
		return "", fmt.Errorf("no improved text in response")
	}

	c.logger.Info("Successfully improved text")
	return response.Text, nil
}
