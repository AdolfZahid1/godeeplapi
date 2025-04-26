package tests

import (
	"context"
	"github.com/AdolfZahid1/godeeplapi"
	"github.com/AdolfZahid1/godeeplapi/models"
	"net/http"
	"os"
	"reflect"
	"testing"
)

var LanguagesPossibleResponse = []models.SupportedLanguage{
	{
		Language:          "BG",
		Name:              "Bulgarian",
		SupportsFormality: false,
	},
	{
		Language:          "CS",
		Name:              "Czech",
		SupportsFormality: false,
	},
	{
		Language:          "DA",
		Name:              "Danish",
		SupportsFormality: false,
	},
	{
		Language:          "DE",
		Name:              "German",
		SupportsFormality: true,
	},
	{
		Language:          "EL",
		Name:              "Greek",
		SupportsFormality: false,
	},
	{
		Language:          "EN-GB",
		Name:              "English (British)",
		SupportsFormality: false,
	},
	{
		Language:          "EN-US",
		Name:              "English (American)",
		SupportsFormality: false,
	},
	{
		Language:          "ES",
		Name:              "Spanish",
		SupportsFormality: true,
	},
	{
		Language:          "ET",
		Name:              "Estonian",
		SupportsFormality: false,
	},
	{
		Language:          "FI",
		Name:              "Finnish",
		SupportsFormality: false,
	},
	{
		Language:          "FR",
		Name:              "French",
		SupportsFormality: true,
	},
	{
		Language:          "HU",
		Name:              "Hungarian",
		SupportsFormality: false,
	},
	{
		Language:          "ID",
		Name:              "Indonesian",
		SupportsFormality: false,
	},
	{
		Language:          "IT",
		Name:              "Italian",
		SupportsFormality: true,
	},
	{
		Language:          "JA",
		Name:              "Japanese",
		SupportsFormality: true,
	},
	{
		Language:          "KO",
		Name:              "Korean",
		SupportsFormality: false,
	},
	{
		Language:          "LT",
		Name:              "Lithuanian",
		SupportsFormality: false,
	},
	{
		Language:          "LV",
		Name:              "Latvian",
		SupportsFormality: false,
	},
	{
		Language:          "NB",
		Name:              "Norwegian (Bokmål)",
		SupportsFormality: false,
	},
	{
		Language:          "NL",
		Name:              "Dutch",
		SupportsFormality: true,
	},
	{
		Language:          "PL",
		Name:              "Polish",
		SupportsFormality: true,
	},
	{
		Language:          "PT-BR",
		Name:              "Portuguese (Brazilian)",
		SupportsFormality: true,
	},
	{
		Language:          "PT-PT",
		Name:              "Portuguese (European)",
		SupportsFormality: true,
	},
	{
		Language:          "RO",
		Name:              "Romanian",
		SupportsFormality: false,
	},
	{
		Language:          "RU",
		Name:              "Russian",
		SupportsFormality: true,
	},
	{
		Language:          "SK",
		Name:              "Slovak",
		SupportsFormality: false,
	},
	{
		Language:          "SL",
		Name:              "Slovenian",
		SupportsFormality: false,
	},
	{
		Language:          "SV",
		Name:              "Swedish",
		SupportsFormality: false,
	},
	{
		Language:          "TR",
		Name:              "Turkish",
		SupportsFormality: false,
	},
	{
		Language:          "UK",
		Name:              "Ukrainian",
		SupportsFormality: false,
	},
	{
		Language:          "ZH",
		Name:              "Chinese (simplified)",
		SupportsFormality: false,
	},
	{
		Language:          "ZH-HANS",
		Name:              "Chinese (simplified)",
		SupportsFormality: false,
	},
}

func TestClient_GetLanguages(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	client := godeeplapi.NewClient(apiKey, false)
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
		want    *models.LanguagesResponse
		wantErr bool
	}{
		{
			name: "Get languages",
			fields: fields{
				client: client,
			},
			args: args{
				ctx: context.Background(),
			},
			want: &models.LanguagesResponse{
				Languages: LanguagesPossibleResponse,
			},
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
			c := godeeplapi.NewClient(apiKey, false)
			got, err := c.GetLanguages(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLanguages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLanguages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUsageAndLimits(t *testing.T) {
	type fields struct {
		baseURL    string
		authKey    string
		httpClient *http.Client
		logger     godeeplapi.Logger
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.UsageAndLimitResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(os.Getenv("DEEPL_API_TOKEN"), false)
			got, err := c.GetUsageAndLimits(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsageAndLimits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsageAndLimits() got = %v, want %v", got, tt.want)
			}
		})
	}
}
