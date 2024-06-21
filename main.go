package main

import (
	"context"
	"log"

	"github.com/jetoneza/mdb-bg-schema-updater/pkg/bigquery"
)

func main() {
    // TODO: Get schema from MongoDB
    // TODO: Create bigquery schema

	log.Println("Updating Schema...")
	updateBigquerySchema()
}

func updateBigquerySchema() {
	ctx := context.Background()

	creds := bigquery.InitializeCredentials(ctx, "")

	writer := bigquery.NewStreamWriter(ctx, creds)

	err := writer.UpdateSchema()
	if err != nil {
		panic(err)
	}
}
