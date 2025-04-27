package godeeplapi

import (
	"context"
	"encoding/json"
	"github.com/AdolfZahid1/godeeplapi/models"
)

// GetUsageAndLimits returns characters translated in the current billing period and current maximum number of characters that can be translated per billing period.
func (c *Client) GetUsageAndLimits(ctx context.Context) (*models.UsageAndLimitResponse, error) {
	err := c.checkAuth()
	if err != nil {
		return nil, err
	}
	respBody, err := c.doRequestWithQuery(ctx, "GET", "/usage", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	var response models.UsageAndLimitResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetLanguages returns struct with all supported languages
func (c *Client) GetLanguages(ctx context.Context) ([]models.SupportedLanguage, error) {
	err := c.checkAuth()
	if err != nil {
		return nil, err
	}
	respBody, err := c.doRequestWithQuery(ctx, "GET", "/languages", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	var response []models.SupportedLanguage
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
