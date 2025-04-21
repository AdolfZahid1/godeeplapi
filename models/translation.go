package models

// TranslationRequest represents parameters for a translation API request.
// It contains all available options for customizing the translation output.
type TranslationRequest struct {
	// Text to be translated. Only UTF-8-encoded plain text is supported.
	// The parameter may be specified many times in a single request, within the request size limit (128KiB).
	// Translations are returned in the same order as they are requested.
	Text []string `json:"text"`

	// The language into which the text should be translated.
	TargetLang string `json:"target_lang"`

	// Language of the text to be translated.
	// If this parameter is omitted, the API will attempt to detect the language of the text and translate it.
	SourceLang string `json:"source_lang,omitempty"`

	// Additional context that can influence a translation but is not translated itself.
	// Characters included in the context parameter will not be counted toward billing.
	Context string `json:"context,omitempty"`

	// When true, the response will include the billed_characters parameter,
	// giving the number of characters from the request that will be counted by DeepL for billing purposes.
	ShowBilledChars bool `json:"show_billed_characters,omitempty" default:"false"`

	// Sets whether the translation engine should first split the input into sentences.
	// Possible values are:
	//   0 - no splitting at all, whole input is treated as one sentence
	//   1 (default when tag_handling is not set to html) - splits on punctuation and on newlines
	//   nonewlines (default when tag_handling=html) - splits on punctuation only, ignoring newlines
	SplitSentences string `json:"split_sentences,omitempty"`

	// Sets whether the translation engine should respect the original formatting,
	// even if it would usually correct some aspects.
	PreserveFormatting bool `json:"preserve_formatting,omitempty" default:"false"`

	// Sets whether the translated text should lean towards formal or informal language.
	// This feature is only available for certain target languages.
	// Setting this parameter with a target language that does not support formality will fail,
	// unless one of the prefer_... options are used. Use FormalityOptions struct for possible options
	Formality string `json:"formality,omitempty"`

	// Specifies which DeepL model should be used for translation.
	// Available options: quality_optimized, prefer_quality_optimized, latency_optimized
	ModelType string `json:"model_type,omitempty"`

	// Specify the glossary to use for the translation.
	// Important: This requires the source_lang parameter to be set.
	// The language pair of the glossary has to match the language pair of the request.
	// Example: def3a26b-3e84-45b3-84ae-0c0aaf3525f7
	GlossaryId string `json:"glossary_id,omitempty"`

	// Sets which kind of tags should be handled. Options currently available:
	// xml
	// html
	TagHandling string `json:"tag_handling,omitempty"`

	// Disable the automatic detection of XML structure by setting the outline_detection parameter to false
	// and selecting the tags that should be considered structure tags.
	// This will split sentences using the splitting_tags parameter.
	OutlineDetection bool `json:"outline_detection,omitempty" default:"true"`

	// Comma-separated list of XML tags which never split sentences.
	NonSplittingTags []string `json:"non_splitting_tags,omitempty"`

	// Comma-separated list of XML tags which always cause splits.
	SplittingTags []string `json:"splitting_tags,omitempty"`

	// Comma-separated list of XML tags that indicate text not to be translated.
	IgnoreTags []string `json:"ignore_tags,omitempty"`
}

// TranslationResponse represents the response from the translation API.
type TranslationResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

// Option constants for formality
var (
	FormalityDefault    = "default"
	FormalityMore       = "more"        // More formal
	FormalityLess       = "less"        // More informal
	FormalityPreferMore = "prefer_more" // Prefer formal, fall back to default
	FormalityPreferLess = "prefer_less" // Prefer informal, fall back to default
)

// Option constants for model type
var (
	ModelQuality       = "quality_optimized"
	ModelPreferQuality = "prefer_quality_optimized"
	ModelLatency       = "latency_optimized"
)

// Option constants for tag handling
var (
	TagXML  = "xml"
	TagHTML = "html"
)

// Option constants for sentence splitting
var (
	SplitNone       = "0"
	SplitDefault    = "1"
	SplitNoNewlines = "nonewlines"
)
