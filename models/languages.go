package models

// TargetLanguageCode defines all supported target languages for translation.
type targetLanguageCode struct {
	EnglishUS    string
	EnglishGB    string
	Bulgarian    string
	Czech        string
	Danish       string
	German       string
	Greek        string
	Spanish      string
	Estonian     string
	Finnish      string
	French       string
	Hungarian    string
	Indonesian   string
	Italian      string
	Japanese     string
	Korean       string
	Lithuanian   string
	Latvian      string
	Norwegian    string
	Dutch        string
	Polish       string
	Russian      string
	PortugueseBR string
	Portuguese   string
	Romanian     string
	Slovak       string
	Slovenian    string
	Swedish      string
	Turkish      string
	Ukrainian    string
	ChineseSimpl string
	ChineseHans  string
}

// TargetLanguage is a predefined instance of TargetLanguageCode.
var TargetLanguage = targetLanguageCode{
	EnglishUS:    "EN-US",
	EnglishGB:    "EN-GB",
	Bulgarian:    "BG",
	Czech:        "CS",
	Danish:       "DA",
	German:       "DE",
	Greek:        "EL",
	Spanish:      "ES",
	Estonian:     "ET",
	Finnish:      "FI",
	French:       "FR",
	Hungarian:    "HU",
	Indonesian:   "ID",
	Italian:      "IT",
	Japanese:     "JA",
	Korean:       "KO",
	Lithuanian:   "LT",
	Latvian:      "LV",
	Norwegian:    "NB",
	Dutch:        "NL",
	Polish:       "PL",
	Russian:      "RU",
	PortugueseBR: "PT-BR",
	Portuguese:   "PT-PT",
	Romanian:     "RO",
	Slovak:       "SK",
	Slovenian:    "SL",
	Swedish:      "SV",
	Turkish:      "TR",
	Ukrainian:    "UK",
	ChineseSimpl: "ZH",
	ChineseHans:  "ZH-HANS",
}

// SourceLanguageCode defines all supported source languages for translation.
type sourceLanguageCode struct {
	English      string
	Bulgarian    string
	Czech        string
	Danish       string
	German       string
	Greek        string
	Spanish      string
	Estonian     string
	Finnish      string
	French       string
	Hungarian    string
	Indonesian   string
	Italian      string
	Japanese     string
	Korean       string
	Lithuanian   string
	Latvian      string
	Norwegian    string
	Dutch        string
	Polish       string
	Russian      string
	Portuguese   string
	Romanian     string
	Slovak       string
	Slovenian    string
	Swedish      string
	Turkish      string
	Ukrainian    string
	ChineseSimpl string
}

// SourceLanguage is a predefined instance of SourceLanguageCode.
var SourceLanguage = sourceLanguageCode{
	English:      "EN",
	Bulgarian:    "BG",
	Czech:        "CS",
	Danish:       "DA",
	German:       "DE",
	Greek:        "EL",
	Spanish:      "ES",
	Estonian:     "ET",
	Finnish:      "FI",
	French:       "FR",
	Hungarian:    "HU",
	Indonesian:   "ID",
	Italian:      "IT",
	Japanese:     "JA",
	Korean:       "KO",
	Lithuanian:   "LT",
	Latvian:      "LV",
	Norwegian:    "NB",
	Dutch:        "NL",
	Polish:       "PL",
	Russian:      "RU",
	Portuguese:   "PT",
	Romanian:     "RO",
	Slovak:       "SK",
	Slovenian:    "SL",
	Swedish:      "SV",
	Turkish:      "TR",
	Ukrainian:    "UK",
	ChineseSimpl: "ZH",
}

// SupportedLanguage represents a language supported by DeepL API.
type SupportedLanguage struct {
	Language          string `json:"language"`
	Name              string `json:"name"`
	SupportsFormality bool   `json:"supports_formality,omitempty"`
}

// LanguagesResponse represents the response from the languages API.
type LanguagesResponse struct {
	Languages []SupportedLanguage `json:"languages"`
}
