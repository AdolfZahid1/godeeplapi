package models

import "io"

// FileTranslationRequest represents parameters for a document translation request.
type FileTranslationRequest struct {
	// Language of the text to be translated.
	// If omitted, the API will attempt to detect the language.
	SourceLang string `json:"source_lang,omitempty"`

	// The language into which the text should be translated.
	TargetLang string `json:"target_lang"`

	// The document file to be translated.
	// This is handled separately in the multipart form upload.
	File io.Reader `json:"-"`

	// The name of the uploaded file.
	FileName string `json:"filename,omitempty"`

	// File extension of desired format of translated file.
	OutputFormat string `json:"output_format,omitempty"`

	// Sets formality level (formal/informal language).
	Formality string `json:"formality,omitempty"`

	// Specify the glossary to use for the translation.
	GlossaryId string `json:"glossary_id,omitempty"`
}

// DocumentResponse represents the initial response when uploading a document.
type DocumentResponse struct {
	DocumentId  string `json:"document_id"`
	DocumentKey string `json:"document_key"`
}

// DocumentStatusResponse represents the status of a document translation.
type DocumentStatusResponse struct {
	DocumentId       string `json:"document_id"`
	DocumentStatus   string `json:"document_status"`
	SecondsRemaining int    `json:"seconds_remaining,omitempty"`
	BilledChars      int    `json:"billed_characters,omitempty"`
	ErrorMessage     string `json:"error_message,omitempty"`
}

// Document status constants
var (
	DocumentStatusQueued      = "queued"
	DocumentStatusTranslating = "translating"
	DocumentStatusDone        = "done"
	DocumentStatusError       = "error"
)
