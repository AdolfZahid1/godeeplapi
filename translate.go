package godeeplapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/AdolfZahid1/godeeplapi/models"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Translate text using the DeepL API
func (c *Client) Translate(ctx context.Context, request models.TranslationRequest) ([]string, error) {
	if c.authKey == "" {
		return nil, fmt.Errorf("DeepL API token is empty")
	}

	respBody, err := c.doRequest(ctx, "POST", "/translate", request, nil)
	if err != nil {
		return nil, err
	}

	var response models.TranslationResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	if len(response.Translations) == 0 {
		return nil, fmt.Errorf("no translations in response")
	}

	var translations []string
	for _, translation := range response.Translations {
		translations = append(translations, translation.Text)
	}

	c.logger.Info("Translated %d text(s)", len(translations))
	return translations, nil
}

// TranslateFile uploads a file for translation and monitors progress
func (c *Client) TranslateFile(ctx context.Context, req models.FileTranslationRequest, targetDir string) (string, error) {
	if c.authKey == "" {
		return "", fmt.Errorf("DeepL API token is empty")
	}

	// Validate inputs
	if req.File == nil {
		return "", fmt.Errorf("file is required")
	}
	if req.TargetLang == "" {
		return "", fmt.Errorf("target language is required")
	}

	// Upload file
	respBody, err := c.uploadFile(ctx, "/document", req)
	if err != nil {
		return "", err
	}

	var response models.DocumentResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Create a child context with timeout for the monitoring process
	monitorCtx, cancel := context.WithTimeout(ctx, 60*time.Minute)
	defer cancel()

	// Create result channel to get the final file path or error
	resultCh := make(chan struct {
		path string
		err  error
	}, 1)

	// Start monitoring in a goroutine
	go func() {
		path, err := c.monitorAndDownload(monitorCtx, response.DocumentId, response.DocumentKey, targetDir)
		resultCh <- struct {
			path string
			err  error
		}{path, err}
	}()

	// Wait for result or context cancellation
	select {
	case result := <-resultCh:
		return result.path, result.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// monitorAndDownload checks document status and downloads when ready
func (c *Client) monitorAndDownload(ctx context.Context, documentId, documentKey, targetDir string) (string, error) {
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			status, err := c.checkDocumentStatus(ctx, documentId, documentKey)
			if err != nil {
				c.logger.Error("Error checking document status: %v", err)
				return "", err
			}

			switch status.DocumentStatus {
			case "done":
				c.logger.Info("Document translation completed, downloading...")
				return c.downloadDocument(ctx, documentId, documentKey, targetDir)

			case "translating":
				waitTime := time.Duration(status.SecondsRemaining+1) * time.Second
				c.logger.Debug("Document is translating, seconds remaining: %d. Checking again in %v",
					status.SecondsRemaining, waitTime)

				select {
				case <-time.After(waitTime):
					// Continue checking
				case <-ctx.Done():
					return "", ctx.Err()
				}

			case "queued":
				c.logger.Debug("Document is queued for translation. Checking again in 10 seconds")
				select {
				case <-time.After(10 * time.Second):
					// Continue checking
				case <-ctx.Done():
					return "", ctx.Err()
				}

			case "error":
				errMsg := fmt.Sprintf("Document processing error: %s", status.ErrorMessage)
				c.logger.Error(errMsg)
				return "", fmt.Errorf(errMsg)

			default:
				errMsg := fmt.Sprintf("Unknown document status: %s", status.DocumentStatus)
				c.logger.Error(errMsg)
				return "", fmt.Errorf(errMsg)
			}
		}
	}
}

// checkDocumentStatus checks the status of a document translation
func (c *Client) checkDocumentStatus(ctx context.Context, documentId, documentKey string) (*models.DocumentStatusResponse, error) {
	requestBody := map[string]string{"document_key": documentKey}

	respBody, err := c.doRequest(ctx, "POST", "/document/"+documentId, requestBody, nil)
	if err != nil {
		return nil, err
	}

	var status models.DocumentStatusResponse
	if err := json.Unmarshal(respBody, &status); err != nil {
		return nil, fmt.Errorf("error unmarshaling status response: %w", err)
	}

	return &status, nil
}

// downloadDocument downloads a translated document
func (c *Client) downloadDocument(ctx context.Context, documentId, documentKey, targetDir string) (string, error) {
	endpoint := "/document/" + documentId + "/result"
	requestBody := map[string]string{"document_key": documentKey}

	// Create temp HTTP request to get the filename
	jsonData, _ := json.Marshal(requestBody)
	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+c.authKey)

	// Send HEAD request first to get the filename
	req.Method = "HEAD"
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error getting file info: %w", err)
	}
	resp.Body.Close()

	// Get filename from Content-Disposition header
	outputFilename := "translated_document"
	if contentDisposition := resp.Header.Get("Content-Disposition"); contentDisposition != "" {
		if _, params, err := mime.ParseMediaType(contentDisposition); err == nil {
			if filename, ok := params["filename"]; ok {
				outputFilename = filename
			}
		}
	}

	// Create target directory if needed
	if targetDir == "" {
		targetDir = "."
	}

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", fmt.Errorf("error creating target directory: %w", err)
	}

	outputPath := filepath.Join(targetDir, outputFilename)

	// Now download the actual file
	if err := c.downloadToFile(ctx, endpoint, requestBody, outputPath); err != nil {
		return "", err
	}

	return outputPath, nil
}
