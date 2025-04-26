# DeepL API Go Wrapper

A lightweight, easy-to-use Go wrapper for the DeepL Translation API that simplifies making translation requests.

## Roadmap
- Add examples for all methods
- Make more tests
  
## Features

- Simple interface for translating text using DeepL's powerful machine translation
- Support for all languages offered by DeepL
- Easy configuration with environment variables
- Comprehensive error handling
- Fully tested implementation

## Installation

```bash
go get github.com/AdolfZahid1e/godeeplapi
```

## Prerequisites

- Go 1.18 or higher
- A DeepL API key (get one at [DeepL API]([https://www.deepl.com/](https://www.deepl.com/en/your-account/keys)))

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "godeeplapi/translator"
)

func main() {
    // Create a translator with your DeepL API key
    tr := &translator.Translator{
        Config: translator.Config{
            DeeplApiToken: "Your DeepL API Token",
        },
    }
    
    // Create a translation request
    request := translator.TranslationRequest{
        Text:       []string{"Hello, world!", "How are you today?"},
        TargetLang: translator.TargetLanguage.German,
    }
    
    // Get translations
    translations, err := tr.Translate(request)
    if err != nil {
        log.Fatalf("Translation failed: %v", err)
    }
    
    // Print the translations
    for i, translation := range translations {
        fmt.Printf("Original: %s\nTranslation: %s\n\n", 
                   request.Text[i], translation)
    }
}
```

### Supported Languages

The wrapper supports all languages provided by the DeepL API:

```go
// Examples of target languages
translator.TargetLanguage.German     // "DE"
translator.TargetLanguage.EnglishUS  // "EN-US"
translator.TargetLanguage.Spanish    // "ES"
translator.TargetLanguage.Japanese   // "JA"
translator.TargetLanguage.Russian    // "RU"
// ... and many more
```

### Specifying Source Language (Optional)

```go
request := translator.TranslationRequest{
    Text:       []string{"Hello, world!"},
    SourceLang: "EN",            // Optional: specify source language
    TargetLang: translator.TargetLanguage.French,
}
```

## API Reference

### Translator

The main struct for handling translations.

```go
type Translator struct {
    Config Config
}
```

### Config

Configuration settings for the DeepL API.

```go
type Config struct {
    DeeplApiToken string
}
```

### TranslationRequest

The request structure sent to DeepL API.

```go
type TranslationRequest struct {
    Text       []string `json:"text"`
    TargetLang string   `json:"target_lang"`
    SourceLang string   `json:"source_lang,omitempty"`
}
```

## Error Handling

The library provides detailed error messages for common issues:

- Missing API key
- Network errors
- API response errors
- JSON parsing errors

## Running Tests

```bash
go test -v ./tests
```

Make sure you have a `.env` file with your DeepL API key in the project root before running tests.

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
