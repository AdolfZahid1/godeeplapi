package translator

import (
	"bytes"
	"deeplapi"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Translator struct {
	Config go_deeplapi.Config
}

func (tr *Translator) Translate(request go_deeplapi.TranslationRequest) ([]string, error) {

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

	var response go_deeplapi.TranslationResponse
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
