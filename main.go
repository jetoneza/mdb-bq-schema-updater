package main

import (
	"context"
	"log"

	"github.com/jetoneza/mdb-bg-schema-updater/pkg/bigquery"
	"github.com/jetoneza/mdb-bg-schema-updater/pkg/mongodb"
)

func main() {
	ctx := context.Background()

	log.Println("Getting Schema from MongoDB...")
	mdbClient := mongodb.New(ctx)
	defer mdbClient.Close()

	schema, err := mdbClient.GetSchemaFromMongoDB(ctx)
	if err != nil {
		log.Printf("Error fetching schema from MongoDB: %v\n", err)
		return
	}

	if schema == nil {
		log.Println("No documents found in MongoDB collection")
		return
	}

	log.Println("Updating Schema...")
	updateBigquerySchema(ctx, schema)
}

func updateBigquerySchema(ctx context.Context, schema map[string]interface{}) {
	// creds := bigquery.InitializeCredentials(ctx, "")

	writer := bigquery.NewStreamWriter(ctx, nil)
	defer writer.Close()

	err := writer.UpdateSchema(schema)
	if err != nil {
		panic(err)
	}
}
