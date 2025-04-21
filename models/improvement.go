package models

// RephraseRequest represents parameters for a text improvement request.
type RephraseRequest struct {
	// Text to be improved. Only UTF-8-encoded plain text is supported.
	Text []string `json:"text"`

	// The language for the text improvement.
	TargetLanguage string `json:"target_lang,omitempty"`

	// Specify a style to rephrase your text.
	WritingStyle string `json:"writing_style,omitempty"`

	// Specify the desired tone for your text.
	Tone string `json:"tone,omitempty"`
}

// RephraseResponse represents the response from the text improvement API.
type RephraseResponse struct {
	Text           string `json:"text"`
	TargetLanguage string `json:"target_lang"`
	SourceLanguage string `json:"detected_source_language"`
}

// Writing style constants
var (
	StyleDefault        = "default"
	StyleAcademic       = "academic"
	StyleBusiness       = "business"
	StyleCasual         = "casual"
	StyleSimple         = "simple"
	StylePreferAcademic = "prefer_academic"
	StylePreferBusiness = "prefer_business"
	StylePreferCasual   = "prefer_casual"
	StylePreferSimple   = "prefer_simple"
)

// Tone constants
var (
	ToneDefault            = "default"
	ToneConfident          = "confident"
	ToneDiplomatic         = "diplomatic"
	ToneEnthusiastic       = "enthusiastic"
	ToneFriendly           = "friendly"
	TonePreferConfident    = "prefer_confident"
	TonePreferDiplomatic   = "prefer_diplomatic"
	TonePreferEnthusiastic = "prefer_enthusiastic"
	TonePreferFriendly     = "prefer_friendly"
)

// ImproveTextLanguage defines language codes for the improve text API.
type ImproveTextLanguage struct {
	German       string
	EnglishUS    string
	EnglishGB    string
	English      string
	Spanish      string
	French       string
	Italian      string
	PortugueseBR string
	Portuguese   string
}

// ImproveTextLanguages provides language codes for text improvement.
var ImproveTextLanguages = ImproveTextLanguage{
	German:       "de",
	EnglishUS:    "en",
	EnglishGB:    "en-GB",
	English:      "en-US",
	Spanish:      "es",
	French:       "fr",
	Italian:      "it",
	PortugueseBR: "pt-BR",
	Portuguese:   "pt",
}
