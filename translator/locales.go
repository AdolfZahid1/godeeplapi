package translator

// TargetLanguageCode defines all supported target languages for translation.
// Each field represents a language code that can be used as a target language.
type TargetLanguageCode struct {
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

// TargetLanguage is a predefined singleton instance of TargetLanguageCode.
// This provides easy access to all supported target language codes.
var TargetLanguage = TargetLanguageCode{
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
// Each field represents a language code that can be used as a source language.
type SourceLanguageCode struct {
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

// SourceLanguage is a predefined singleton instance of SourceLanguageCode.
// This provides easy access to all supported source language codes.
var SourceLanguage = SourceLanguageCode{
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
