package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"godeeplapi"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

type Translator struct {
	Config godeeplapi.Config
}

func (tr *Translator) Translate(request TranslationRequest) ([]string, error) {

	// Check again if the token is still empty after init
	if tr.Config.DeeplApiToken == "" {
		return nil, fmt.Errorf("Error: DeepL API token is empty")
	}
	var url string
	if tr.Config.IsPro {
		url = "https://api.deepl.com/v2/translate"
	} else {
		url = "https://api-free.deepl.com/v2/translate"
	}
	jsonStr, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	// Fix: Set the Authorization header correctly
	req.Header.Set("Authorization", "DeepL-Auth-Key "+tr.Config.DeeplApiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing body: %v", err)
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// If status code is not 200, log the error
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var response translationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Translations) == 0 {
		return nil, fmt.Errorf("No translations in response")
	}

	var translations []string
	for _, translation := range response.Translations {
		translations = append(translations, translation.Text)
	}
	log.Printf("Translations: %v", translations)
	return translations, nil
}
func (tr *Translator) TranslateFile(req FileTranslationRequest, path string) ([]byte, error) {

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add text fields
	if req.SourceLang != "" {
		err := writer.WriteField("source_lang", req.SourceLang)
		if err != nil {
			return nil, err
		}
	}
	err := writer.WriteField("target_lang", req.TargetLang)
	if err != nil {
		return nil, err
	}
	if req.FileName != "" {
		err := writer.WriteField("filename", req.FileName)
		if err != nil {
			return nil, err
		}
	}
	if req.OutputFormat != "" {
		err := writer.WriteField("output_format", req.OutputFormat)
		if err != nil {
			return nil, err
		}
	}
	if req.Formality != "" {
		err := writer.WriteField("formality", req.Formality)
		if err != nil {
			return nil, err
		}
	}
	if req.GlossaryId != "" {
		err := writer.WriteField("glossary_id", req.GlossaryId)
		if err != nil {
			return nil, err
		}
	}

	// Add the file
	fileWriter, err := writer.CreateFormFile("file", path)
	if err != nil {
		return nil, fmt.Errorf("error creating form file: %w", err)
	}

	_, err = io.Copy(fileWriter, req.File)
	if err != nil {
		return nil, fmt.Errorf("error copying file data: %w", err)
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing writer: %w", err)
	}

	// Create the HTTP request
	var apiUrl string
	if tr.Config.IsPro {
		apiUrl = "https://api.deepl.com/v2/document"
	} else {
		apiUrl = "https://api-free.deepl.com/v2/document"
	}
	httpReq, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Authorization", "DeepL-Auth-Key "+tr.Config.DeeplApiToken)

	// Send the request
	client := &http.Client{}

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}(resp.Body)

	// Read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	// Check for error status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s, %s", resp.Status, string(respBody))
	}

	return respBody, nil
}
