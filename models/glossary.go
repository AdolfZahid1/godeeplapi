package models

// GlossaryListResponse contains supported language pairs for glossaries.
type GlossaryListResponse struct {
	SupportedLanguages []GlossaryLangPair `json:"supported_languages"`
}

// GlossaryLangPair represents a source-target language pair supported by glossaries.
type GlossaryLangPair struct {
	// The language for source texts. Use SourceLanguageCode struct for options.
	SourceLanguage string `json:"source_lang"`
	// The language for target texts. Use TargetLanguageCode struct for options.
	TargetLanguage string `json:"target_lang"`
}

// Dictionary represents a single language pair dictionary within a glossary.
type Dictionary struct {
	// The language for source texts.
	SourceLanguage string `json:"source_lang"`
	// The language for target texts.
	TargetLanguage string `json:"target_lang"`
	// Optional. The entries of the glossary in the specified format.
	Entries string `json:"entries,omitempty"`
	// Optional. Format: "tsv" (default) or "csv".
	EntriesFormat string `json:"entries_format,omitempty"`
}

// Glossary represents a glossary resource with its metadata.
type Glossary struct {
	// A unique ID assigned to the glossary
	GlossaryID string `json:"glossary_id,omitempty"`
	// Name associated with the glossary
	Name string `json:"name,omitempty"`
	// Dictionaries contained in this glossary
	Dictionaries []Dictionary `json:"dictionaries,omitempty"`
	// Creation time in ISO 8601-1:2019 format
	CreatedAt string `json:"creation_time,omitempty"`
}

// CreateGlossaryRequest holds data for glossary creation.
type CreateGlossaryRequest struct {
	// Name for the new glossary
	Name string `json:"name"`
	// Dictionaries to populate glossary with
	Dictionaries []Dictionary `json:"dictionaries"`
}

// CreateGlossaryResponse represents API response after glossary creation.
type CreateGlossaryResponse struct {
	Glossary
}

// AllGlossaryListResponse contains all glossaries for the account.
type AllGlossaryListResponse struct {
	// List of all glossaries.
	Glossaries []Glossary `json:"glossaries,omitempty"`
}

// EditGlossaryRequest hold data for glossary edit.
type EditGlossaryRequest struct {
	// Name associated with the glossary.
	Name string `json:"name,omitempty"`
	//Dictionaries to edit the glossary with. Currently only supports 0 or 1 dictionaries in the array.
	Dictionaries []Dictionary `json:"dictionaries,omitempty"`
}
type GlossaryEntriesResponse struct {
	Dictionaries []Dictionary `json:"dictionaries,omitempty"`
}

// EditOrCreateDictionaryInGlossaryResponse represents answer from API.
type EditOrCreateDictionaryInGlossaryResponse struct {
	//The language in which the source texts in the glossary are specified.
	SourceLanguage string `json:"source_lang,omitempty"`
	//The language in which the target texts in the glossary are specified.
	TargetLanguage string `json:"target_lang,omitempty"`
	//The number of entries in the glossary.
	EntryCount int `json:"entry_count,omitempty"`
}
