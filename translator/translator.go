package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"godeeplapi"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
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
func (tr *Translator) TranslateFile(req FileTranslationRequest, source, target string) error {
	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add text fields
	if req.SourceLang != "" {
		err := writer.WriteField("source_lang", req.SourceLang)
		if err != nil {
			return err
		}
	}
	err := writer.WriteField("target_lang", req.TargetLang)
	if err != nil {
		return err
	}
	if req.FileName != "" {
		err := writer.WriteField("filename", req.FileName)
		if err != nil {
			return err
		}
	}
	if req.OutputFormat != "" {
		err := writer.WriteField("output_format", req.OutputFormat)
		if err != nil {
			return err
		}
	}
	if req.Formality != "" {
		err := writer.WriteField("formality", req.Formality)
		if err != nil {
			return err
		}
	}
	if req.GlossaryId != "" {
		err := writer.WriteField("glossary_id", req.GlossaryId)
		if err != nil {
			return err
		}
	}

	// Add the file
	fileWriter, err := writer.CreateFormFile("file", source)
	if err != nil {
		return fmt.Errorf("error creating form file: %w", err)
	}

	_, err = io.Copy(fileWriter, req.File)
	if err != nil {
		return fmt.Errorf("error copying file data: %w", err)
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		return fmt.Errorf("error closing writer: %w", err)
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
		return fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Authorization", "DeepL-Auth-Key "+tr.Config.DeeplApiToken)

	// Send the request
	client := &http.Client{}

	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
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
		return fmt.Errorf("error reading response: %w", err)
	}

	// Check for error status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: %s, %s", resp.Status, string(respBody))
	}
	var response documentResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Start a goroutine to periodically check document status and download when done
	go func() {
		var checkDocumentStatusRecursive func() error

		checkDocumentStatusRecursive = func() error {
			documentStatus, err := tr.checkDocumentStatus(response.DocumentId, response.DocumentKey)
			if err != nil {
				log.Printf("Error checking document status: %v", err)
				return err
			}

			var unmarshalledDocumentStatus documentStatusResponse
			err = json.Unmarshal(documentStatus, &unmarshalledDocumentStatus)
			if err != nil {
				log.Printf("Error unmarshaling document status: %v", err)
				return err
			}

			switch unmarshalledDocumentStatus.DocumentStatus {
			case "done":
				log.Printf("Document translation completed, downloading...")
				err := tr.downloadDocument(response.DocumentId, response.DocumentKey, target)
				if err != nil {
					log.Printf("Error downloading document: %v", err)
					return err
				}
				log.Printf("Document successfully downloaded to: %s", target)
				return nil

			case "translating":
				sleepTime := time.Duration(unmarshalledDocumentStatus.SecondsRemaining+1) * time.Second
				log.Printf("Document is translating, seconds remaining: %d. Checking again in %v",
					unmarshalledDocumentStatus.SecondsRemaining, sleepTime)
				time.Sleep(sleepTime)
				return checkDocumentStatusRecursive() // Recursive call after sleeping

			case "queued":
				log.Printf("Document is queued for translation. Checking again in 10 seconds")
				time.Sleep(10 * time.Second)
				return checkDocumentStatusRecursive() // Recursive call after sleeping

			case "error":
				errMsg := "Document processing error reported by DeepL API"
				log.Printf("%s", errMsg)
				return fmt.Errorf(errMsg)

			default:
				errMsg := fmt.Sprintf("Unknown document status: %s", unmarshalledDocumentStatus.DocumentStatus)
				log.Printf("%s", errMsg)
				return fmt.Errorf(errMsg)
			}
		}

		// Start the recursive status checking
		err := checkDocumentStatusRecursive()
		if err != nil {
			log.Printf("Translation process failed: %v", err)
		}
	}()

	return nil
}

// helper function to download a file from endpoint
func (tr *Translator) downloadDocument(documentId, documentKey, target string) error {
	var apiUrl string
	if tr.Config.IsPro {
		apiUrl = "https://api.deepl.com/v2/document/" + documentId + "/result"
	} else {
		apiUrl = "https://api-free.deepl.com/v2/document/" + documentId + "/result"
	}

	// Create the correct JSON request body
	requestBody := map[string]string{"document_key": documentKey}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "DeepL-Auth-Key "+tr.Config.DeeplApiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s, %s", resp.Status, string(body))
	}

	// Get filename from Content-Disposition header if available
	outputFilename := "translated_document"
	if contentDisposition := resp.Header.Get("Content-Disposition"); contentDisposition != "" {
		if _, params, err := mime.ParseMediaType(contentDisposition); err == nil {
			if filename, ok := params["filename"]; ok {
				outputFilename = filename
			}
		}
	}

	// If target is empty, use current directory
	if target == "" {
		target = "."
	}

	// Create full path to save the file
	outputPath := filepath.Join(target, outputFilename)

	// Create the output file
	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer out.Close()

	// Write the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file: %w", err)
	}

	return nil
}
func (tr *Translator) checkDocumentStatus(documentId, documentKey string) ([]byte, error) {
	var apiUrl string
	if tr.Config.IsPro {
		apiUrl = "https://api.deepl.com/v2/document/" + documentId
	} else {
		apiUrl = "https://api-free.deepl.com/v2/document/" + documentId
	}
	jsonStr, err := json.Marshal(fmt.Sprintf("document_key:%s", documentKey))
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "DeepL-Auth-Key "+tr.Config.DeeplApiToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s, %s", resp.Status, string(body))
	}
	return body, nil
}
