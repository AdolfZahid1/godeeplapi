package improver

type RephraseRequest struct {
	//Text to be improved. Only UTF-8-encoded plain text is supported. Improvements are returned in the same order as they are requested.
	Text []string `json:"text"`

	//The language for the text improvement.
	// Use ImproveTextLanguageOptions struct for available options
	TargetLanguage string `json:"target_lang,omitempty"`

	//Specify a style to rephrase your text in a way that fits your audience and goals. The prefer_ prefix allows falling back to the default style if the language does not yet support styles.
	// Use WritingStyleOptions for available options
	WritingStyle string `json:"writing_style,omitempty"`

	//Specify the desired tone for your text. The prefer_ prefix allows falling back to the default tone if the language does not yet support tones.
	// Use ToneOptions for available options
	Tone string `json:"tone,omitempty"`
}
type rephraseResponse struct {
	Text           string `json:"text"`
	TargetLanguage string `json:"target_lang"`
	SourceLanguage string `json:"detected_source_language"`
}

type textLanguage struct {
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

var ImproveTextLanguageOptions = textLanguage{
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

type writingStyle struct {
	Academic       string
	Business       string
	Casual         string
	Default        string
	Simple         string
	PreferAcademic string
	PreferBusiness string
	PreferCasual   string
	PreferSimple   string
}

var WritingStyleOptions = writingStyle{
	Academic:       "academic",
	Business:       "business",
	Casual:         "casual",
	Default:        "default",
	Simple:         "simple",
	PreferAcademic: "prefer_academic",
	PreferBusiness: "prefer_business",
	PreferCasual:   "prefer_casual",
	PreferSimple:   "prefer_simple",
}

type tone struct {
	Confident          string
	Default            string
	Diplomatic         string
	Enthusiastic       string
	Friendly           string
	PreferConfident    string
	PreferDiplomatic   string
	PreferEnthusiastic string
	PreferFriendly     string
}

var ToneOptions = tone{
	Confident:          "confident",
	Default:            "default",
	Diplomatic:         "diplomatic",
	Enthusiastic:       "enthusiastic",
	Friendly:           "friendly",
	PreferConfident:    "prefer_confident",
	PreferDiplomatic:   "prefer_diplomatic",
	PreferEnthusiastic: "prefer_enthusiastic",
	PreferFriendly:     "prefer_friendly",
}
