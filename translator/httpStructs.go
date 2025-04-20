package translator

import "io"

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

type FileTranslationRequest struct {
	// Language of the text to be translated.
	// If this parameter is omitted, the API will attempt to detect the language of the text and translate it.
	SourceLang string `json:"source_lang,omitempty"`

	// The language into which the text should be translated.
	// Use TargetLanguageCode struct to define the lang
	TargetLang string `json:"target_lang"`

	//The document file to be translated. The file name should be included in this part's content disposition. As an alternative, the filename parameter can be used. The following file types and extensions are supported:
	//
	//    docx - Microsoft Word Document
	//    pptx - Microsoft PowerPoint Document
	//    xlsx - Microsoft Excel Document
	//    pdf - Portable Document Format
	//    htm / html - HTML Document
	//    txt - Plain Text Document
	//    xlf / xliff - XLIFF Document, version 2.1
	//    srt - SRT Document
	File io.Reader `json:"-"`

	//The name of the uploaded file. Can be used as an alternative to including the file name in the file part's content disposition.
	FileName string `json:"filename,omitempty"`

	//File extension of desired format of translated file, for example: docx. If unspecified, by default the translated file will be in the same format as the input file.
	OutputFormat string `json:"output_format,omitempty"`

	// Sets whether the translated text should lean towards formal or informal language.
	// This feature is only available for certain target languages.
	// Setting this parameter with a target language that does not support formality will fail,
	// unless one of the prefer_... options are used. Use FormalityOptions struct for possible options
	Formality string `json:"formality,omitempty"`

	// Specify the glossary to use for the translation.
	// Important: This requires the source_lang parameter to be set.
	// The language pair of the glossary has to match the language pair of the request.
	// Example: def3a26b-3e84-45b3-84ae-0c0aaf3525f7
	GlossaryId string `json:"glossary_id,omitempty"`
}

// TranslationResponse represents the response from the translation API.
type translationResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

// Formality defines formality options for translations.
type Formality struct {
	Default        string `default:"default"`
	Formal         string `default:"more"`
	Informal       string `default:"less"`
	PreferFormal   string `default:"prefer_more"`
	PreferInformal string `default:"prefer_less"`
}

// ModelType defines options for the DeepL translation model.
type ModelType struct {
	Quality       string `default:"quality_optimized"`
	PreferQuality string `default:"prefer_quality_optimized"`
	Latency       string `default:"latency_optimized"`
}

// Tag defines options for tag handling in translations.
type Tag struct {
	Xml  string `default:"xml"`
	Html string `default:"html"`
}

// SplitSentences defines options for sentence splitting in translations.
type SplitSentences struct {
	No         string `default:"0"`
	Yes        string `default:"1"`
	NoNewLines string `default:"nonewlines"`
}

// FormalityOptions Predefined formality options
var FormalityOptions = Formality{
	Default:        "default",
	Formal:         "more",
	Informal:       "less",
	PreferFormal:   "prefer_more",
	PreferInformal: "prefer_less",
}

// ModelTypeOptions Predefined model type options
var ModelTypeOptions = ModelType{
	Quality:       "quality_optimized",
	PreferQuality: "prefer_quality_optimized",
	Latency:       "latency_optimized",
}

// TagOptions Predefined tag handling options
var TagOptions = Tag{
	Xml:  "xml",
	Html: "html",
}

// SplitSentencesOptions Predefined sentence splitting options
var SplitSentencesOptions = SplitSentences{
	No:         "0",
	Yes:        "1",
	NoNewLines: "nonewlines",
}
