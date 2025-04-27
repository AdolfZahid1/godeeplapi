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
		{
			name: "Text improved",
			fields: fields{
				client: godeeplapi.NewClient(apiKey, false),
			},
			args: args{
				ctx: context.Background(),
				req: models.RephraseRequest{
					Text:           []string{"The meeting was very good and we talk about many important things. The project is making progress and the team is working hard to finish it soon. I think we will meet the deadlines if we continue to work like this. There was some problems with the database but John fixed it. The clients are happy with what we shown them last week. We need to focus on the user interface because it's not so good and needs improvements. Everyone must submit their reports by Friday so we can review them on Monday meeting. I will sent an email with more details tomorrow."},
					TargetLanguage: models.TargetLanguage.EnglishUS,
				},
			},
			want:    "The meeting was very good and we talked about many important things. The project is progressing and the team is working hard to finish it soon. I think we will meet the deadlines if we keep working like this. There were some problems with the database, but John fixed them. The clients are happy with what we showed them last week. We need to focus on the user interface because it's not so good and needs to be improved. Everyone needs to submit their reports by Friday so we can review them at the Monday meeting. I will send an email tomorrow with more details.",
			wantErr: false,
		},
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
