package tests

import (
	"deeplapi"
	"deeplapi/translator"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestTranslator_Translate(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	type fields struct {
		config go_deeplapi.Config
	}
	type args struct {
		request go_deeplapi.TranslationRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{name: "No API key",
			fields: fields{config: go_deeplapi.Config{DeeplApiToken: ""}},
			args: args{
				request: go_deeplapi.TranslationRequest{
					Text:       []string{"Test"},
					TargetLang: go_deeplapi.TargetLanguage.EnglishUS,
				},
			},
			want: nil,
		},
		{name: "Translate to german \"Hello, World!\"",
			fields: fields{config: go_deeplapi.Config{DeeplApiToken: os.Getenv("DEEPL_API_TOKEN")}},
			args: args{
				request: go_deeplapi.TranslationRequest{
					Text:       []string{"Hello, World!"},
					TargetLang: go_deeplapi.TargetLanguage.German,
				},
			},
			want: []string{"Hallo, Welt!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &translator.Translator{
				Config: tt.fields.config,
			}
			if got, _ := tr.Translate(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
