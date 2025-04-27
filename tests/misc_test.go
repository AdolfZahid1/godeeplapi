package tests

import (
	"context"
	"github.com/AdolfZahid1/godeeplapi"
	"github.com/AdolfZahid1/godeeplapi/models"
	"os"
	"testing"
)

var LanguagesPossibleResponse = []models.SupportedLanguage{
	{
		Language: "BG",
		Name:     "Bulgarian",
	},
	{
		Language: "CS",
		Name:     "Czech",
	},
	{
		Language: "DA",
		Name:     "Danish",
	},
	{
		Language: "DE",
		Name:     "German",
	},
	{
		Language: "EL",
		Name:     "Greek",
	},
	{
		Language: "EN-GB",
		Name:     "English (British)",
	},
	{
		Language: "EN-US",
		Name:     "English (American)",
	},
	{
		Language: "ES",
		Name:     "Spanish",
	},
	{
		Language: "ET",
		Name:     "Estonian",
	},
	{
		Language: "FI",
		Name:     "Finnish",
	},
	{
		Language: "FR",
		Name:     "French",
	},
	{
		Language: "HU",
		Name:     "Hungarian",
	},
	{
		Language: "ID",
		Name:     "Indonesian",
	},
	{
		Language: "IT",
		Name:     "Italian",
	},
	{
		Language: "JA",
		Name:     "Japanese",
	},
	{
		Language: "KO",
		Name:     "Korean",
	},
	{
		Language: "LT",
		Name:     "Lithuanian",
	},
	{
		Language: "LV",
		Name:     "Latvian",
	},
	{
		Language: "NB",
		Name:     "Norwegian (Bokmål)",
	},
	{
		Language: "NL",
		Name:     "Dutch",
	},
	{
		Language: "PL",
		Name:     "Polish",
	},
	{
		Language: "PT-BR",
		Name:     "Portuguese (Brazilian)",
	},
	{
		Language: "PT-PT",
		Name:     "Portuguese (European)",
	},
	{
		Language: "RO",
		Name:     "Romanian",
	},
	{
		Language: "RU",
		Name:     "Russian",
	},
	{
		Language: "SK",
		Name:     "Slovak",
	},
	{
		Language: "SL",
		Name:     "Slovenian",
	},
	{
		Language: "SV",
		Name:     "Swedish",
	},
	{
		Language: "TR",
		Name:     "Turkish",
	},
	{
		Language: "UK",
		Name:     "Ukrainian",
	},
	{
		Language: "ZH",
		Name:     "Chinese (simplified)",
	},
	{
		Language: "ZH-HANS",
		Name:     "Chinese (simplified)",
	},
}

func TestClient_GetLanguages(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	if apiKey == "" {
		t.Skip("DEEPL_API_TOKEN environment variable not set")
	}

	type fields struct {
		client *godeeplapi.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.SupportedLanguage // Updated to match new return type
		wantErr bool
	}{
		{
			name: "Get languages",
			fields: fields{
				client: godeeplapi.NewClient(apiKey, false),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    LanguagesPossibleResponse, // This is now an array directly
			wantErr: false,
		},
		{
			name: "Get languages Error",
			fields: fields{
				client: nil,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use the client from the test case
			c := tt.fields.client

			var handle []models.SupportedLanguage
			var err error

			handle, err = c.GetLanguages(tt.args.ctx)

			// Check error condition using got/want pattern
			if got, want := err != nil, tt.wantErr; got != want {
				t.Fatalf("error condition: got=%v, want=%v, err=%v", got, want, err)
			}

			// Only check response if we didn't expect an error
			if !tt.wantErr {
				if got := handle; got == nil {
					t.Errorf("handle: got=nil, want=non-nil response")
				} else if got := len(handle); got == 0 {
					t.Errorf("handle length: got=%d, want=non-empty response", got)
				} else {
					t.Logf("Received %d languages from DeepL API", len(handle))
				}
			}
		})
	}

}

func TestClient_GetUsageAndLimits(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	if apiKey == "" {
		t.Skip("DEEPL_API_TOKEN environment variable not set")
	}

	type fields struct {
		client *godeeplapi.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Get usage and limits",
			fields: fields{
				client: godeeplapi.NewClient(apiKey, false),
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.client

			got, err := c.GetUsageAndLimits(tt.args.ctx)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsageAndLimits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got == nil {
					t.Fatal("GetUsageAndLimits() returned nil response")
				}

				// Validate the structure by checking field types/constraints
				if got.CharLimit < 0 {
					t.Errorf("CharLimit should be non-negative, got: %d", got.CharLimit)
				}

				if got.CharCount < 0 {
					t.Errorf("CharCount should be non-negative, got: %d", got.CharCount)
				}

				// CharCount should be less than or equal to CharLimit
				if got.CharCount > got.CharLimit {
					t.Errorf("CharCount (%d) should not exceed CharLimit (%d)",
						got.CharCount, got.CharLimit)
				}

				// Log the values for informational purposes
				t.Logf("Usage stats - Character Count: %d, Character Limit: %d",
					got.CharCount, got.CharLimit)
			}
		})
	}
}
