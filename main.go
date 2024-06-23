package main

import (
	"context"
	"log"

	"github.com/jetoneza/mdb-bg-schema-updater/pkg/bigquery"
	"github.com/jetoneza/mdb-bg-schema-updater/pkg/mongodb"
)

func main() {
	ctx := context.Background()

	schema := getMongoDBSchema(ctx)

	updateBigquerySchema(ctx, schema)
}

func getMongoDBSchema(ctx context.Context) map[string]interface{} {
	log.Println("Fetching schema from MongoDB")

	mdbClient := mongodb.New(ctx)
	defer mdbClient.Close()

	schema, err := mdbClient.GetSchemaFromMongoDB(ctx)
	if err != nil {
		log.Printf("Error fetching schema from MongoDB: %v\n", err)
		return nil
	}

	if schema == nil {
		log.Println("No documents found in MongoDB collection")
		return nil
	}

	return schema
}

func updateBigquerySchema(ctx context.Context, schema map[string]interface{}) {
	log.Println("Updating BigQuery schema")

	// TODO: Support for service account credentials json file
	// creds := bigquery.InitializeCredentials(ctx, "")

	writer := bigquery.New(ctx, nil)
	defer writer.Close()

	err := writer.UpdateSchema(schema)
	if err != nil {
		panic(err)
	}
}
