package godeeplapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/AdolfZahid1/godeeplapi/models"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (c *Client) uploadFile(ctx context.Context, endpoint string, req models.FileTranslationRequest) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add text fields with error handling
	fields := map[string]string{
		"target_lang": req.TargetLang,
	}

	if req.SourceLang != "" {
		fields["source_lang"] = req.SourceLang
	}
	if req.FileName != "" {
		fields["filename"] = req.FileName
	}
	if req.OutputFormat != "" {
		fields["output_format"] = req.OutputFormat
	}
	if req.Formality != "" {
		fields["formality"] = req.Formality
	}
	if req.GlossaryId != "" {
		fields["glossary_id"] = req.GlossaryId
	}

	for key, value := range fields {
		if err := writer.WriteField(key, value); err != nil {
			return nil, fmt.Errorf("error adding field %s: %w", key, err)
		}
	}

	// Create the file part
	fileWriter, err := writer.CreateFormFile("file", filepath.Base(req.FileName))
	if err != nil {
		return nil, fmt.Errorf("error creating form file: %w", err)
	}

	// Use buffered copy for better performance
	buf := make([]byte, 32*1024) // 32KB buffer
	if _, err := io.CopyBuffer(fileWriter, req.File, buf); err != nil {
		return nil, fmt.Errorf("error copying file data: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("error closing multipart writer: %w", err)
	}

	// Create request
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Authorization", "DeepL-Auth-Key "+c.authKey)

	c.logger.Debug("Uploading file to %s", c.baseURL+endpoint)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    http.StatusText(resp.StatusCode),
			Details:    string(respBody),
		}
	}

	return respBody, nil
}

func (c *Client) downloadToFile(ctx context.Context, endpoint string, requestBody interface{}, outputPath string) error {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+c.authKey)

	c.logger.Debug("Downloading file from %s", c.baseURL+endpoint)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    http.StatusText(resp.StatusCode),
			Details:    string(body),
		}
	}

	// Create output directory if it doesn't exist
	if dir := filepath.Dir(outputPath); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating output directory: %w", err)
		}
	}

	// Create output file
	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer out.Close()

	// Use buffered download for performance
	buf := make([]byte, 64*1024) // 64KB buffer
	if _, err := io.CopyBuffer(out, resp.Body, buf); err != nil {
		return fmt.Errorf("error saving file: %w", err)
	}

	c.logger.Info("File successfully downloaded to: %s", outputPath)
	return nil
}
