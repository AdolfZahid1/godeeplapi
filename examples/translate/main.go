package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AdolfZahid1/godeeplapi"
	"github.com/AdolfZahid1/godeeplapi/models"
)

func main() {
	// Create a new client with options
	client := godeeplapi.NewClient(
		os.Getenv("DEEPL_API_TOKEN"),
		false, // Using free API
		godeeplapi.WithTimeout(30*time.Second),
	)

	// Create a translation request
	request := models.TranslationRequest{ // Use models directly
		Text:       []string{"Hello, world!", "How are you today?"},
		TargetLang: models.TargetLanguage.German,
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get translations
	translations, err := client.Translate(ctx, request)
	if err != nil {
		log.Fatalf("Translation failed: %v", err)
	}

	// Print the translations
	for i, translation := range translations {
		fmt.Printf("Original: %s\nTranslation: %s\n\n",
			request.Text[i], translation)
	}
}
