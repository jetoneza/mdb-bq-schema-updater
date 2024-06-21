package bigquery

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
	"golang.org/x/oauth2/google"
)

func InitializeCredentials(ctx context.Context, credentials string) *google.Credentials {
	creds, err := google.CredentialsFromJSON(ctx, []byte(credentials), bigquery.Scope)
	if err != nil {
		log.Fatal(err)
	}

	return creds
}
