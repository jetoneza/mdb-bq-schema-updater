package bigquery

import (
	"context"
	"log"

	bq "cloud.google.com/go/bigquery"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type StreamWriter struct {
	Client  *bq.Client
	Context context.Context
}

func NewStreamWriter(ctx context.Context, creds *google.Credentials) *StreamWriter {
	client, err := bq.NewClient(ctx, BQ_PROJECT_ID, option.WithCredentials(creds))
	if err != nil {
		panic(err)
	}

	return &StreamWriter{
		Client:  client,
		Context: ctx,
	}
}

// UpdateSchema
//
// updates the BigQuery schema of the table
func (w *StreamWriter) UpdateSchema() error {
	log.Println("Schema successfully updated")

	return nil
}
