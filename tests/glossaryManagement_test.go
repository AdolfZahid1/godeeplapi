package tests

import (
	"context"
	"github.com/AdolfZahid1/godeeplapi"
	"github.com/AdolfZahid1/godeeplapi/models"
	"os"
	"reflect"
	"testing"
)

func TestClient_CreateGlossary(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	type fields struct {
		client *godeeplapi.Client
	}
	type args struct {
		ctx context.Context
		req models.CreateGlossaryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.CreateGlossaryResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(apiKey, false)
			got, err := c.CreateGlossary(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGlossary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGlossary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DeleteAllLangDictionaries(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	type fields struct {
		client *godeeplapi.Client
	}
	type args struct {
		ctx   context.Context
		id    string
		query models.GlossaryLangPair
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(apiKey, false)
			if err := c.DeleteAllLangDictionaries(tt.args.ctx, tt.args.id, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAllLangDictionaries() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_DeleteGlossary(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	type fields struct {
		client *godeeplapi.Client
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(apiKey, false)
			if err := c.DeleteGlossary(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteGlossary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_EditGlossary(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	type fields struct {
		client *godeeplapi.Client
	}
	type args struct {
		ctx context.Context
		id  string
		req models.EditGlossaryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Glossary
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(apiKey, false)
			got, err := c.EditGlossary(tt.args.ctx, tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditGlossary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EditGlossary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetGlossaryByID(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	type fields struct {
		client *godeeplapi.Client
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Glossary
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(apiKey, false)
			got, err := c.GetGlossaryByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGlossaryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGlossaryByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetGlossaryEntries(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	type fields struct {
		client *godeeplapi.Client
	}
	type args struct {
		ctx   context.Context
		id    string
		query models.GlossaryLangPair
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.GlossaryEntriesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(apiKey, false)
			got, err := c.GetGlossaryEntries(tt.args.ctx, tt.args.id, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGlossaryEntries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGlossaryEntries() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ListAllGlossaries(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
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
		want    *models.AllGlossaryListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(apiKey, false)
			got, err := c.ListAllGlossaries(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAllGlossaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllGlossaries() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ListLangPairsSupportedByGlossaries(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")

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
		want    *models.GlossaryListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(apiKey, false)
			got, err := c.ListLangPairsSupportedByGlossaries(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListLangPairsSupportedByGlossaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListLangPairsSupportedByGlossaries() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ReplaceOrCreateDictionaryInGlossary(t *testing.T) {
	apiKey := os.Getenv("DEEPL_API_TOKEN")
	type fields struct {
		client *godeeplapi.Client
	}
	type args struct {
		ctx context.Context
		id  string
		req models.Dictionary
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.EditOrCreateDictionaryInGlossaryResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := godeeplapi.NewClient(apiKey, false)
			got, err := c.ReplaceOrCreateDictionaryInGlossary(tt.args.ctx, tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReplaceOrCreateDictionaryInGlossary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReplaceOrCreateDictionaryInGlossary() got = %v, want %v", got, tt.want)
			}
		})
	}
}
