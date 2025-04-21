package tests

import (
	"context"
	"github.com/AdolfZahid1/godeeplapi"
	"github.com/AdolfZahid1/godeeplapi/models"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestClient_Translate(t *testing.T) {
	type args struct {
		request models.TranslationRequest
	}

	tests := []struct {
		name    string
		client  *godeeplapi.Client
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "No API key",
			client: godeeplapi.NewClient(
				"", // Empty API key
				false,
			),
			args: args{
				request: models.TranslationRequest{
					Text:       []string{"Test"},
					TargetLang: models.TargetLanguage.EnglishUS,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Translate to german \"Hello, World!\"",
			client: godeeplapi.NewClient(
				os.Getenv("DEEPL_API_TOKEN"),
				false,
			),
			args: args{
				request: models.TranslationRequest{
					Text:       []string{"Hello, World!"},
					TargetLang: models.TargetLanguage.German,
				},
			},
			want:    []string{"Hallo, Welt!"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a context with timeout for each test
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			got, err := tt.client.Translate(ctx, tt.args.request)

			// Check error expectations
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Only check response content if we didn't expect an error
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
