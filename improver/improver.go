package improver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"godeeplapi"
	"io"
	"log"
	"net/http"
)

type Improver struct {
	Config godeeplapi.Config
}

func (imp *Improver) ImproveText(req RephraseRequest) (string, error) {
	// Check again if the token is still empty after init
	if imp.Config.DeeplApiToken == "" {
		return "", fmt.Errorf("error: DeepL API token is empty")
	}
	var url string
	if imp.Config.IsPro {
		url = "https://api.deepl.com/v2/translate"
	} else {
		url = "https://api-free.deepl.com/v2/translate"
	}
	jsonStr, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}

	// Fix: Set the Authorization header correctly
	request.Header.Set("Authorization", "DeepL-Auth-Key "+imp.Config.DeeplApiToken)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
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
		return "", err
	}

	// If status code is not 200, log the error
	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var response rephraseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	if len(response.Text) == 0 {
		return "", fmt.Errorf("no translations in response")
	}

	return response.Text, nil
}
