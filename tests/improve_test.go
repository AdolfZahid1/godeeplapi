package tests

import (
	"context"
	"github.com/AdolfZahid1/godeeplapi"
	"github.com/AdolfZahid1/godeeplapi/models"
	"os"
	"testing"
)

func TestClient_ImproveText(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	type fields struct {
		client *godeeplapi.Client
	}
	type args struct {
		ctx context.Context
		req models.RephraseRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(apiKey, false)
			got, err := c.ImproveText(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImproveText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ImproveText() got = %v, want %v", got, tt.want)
			}
		})
	}
}
