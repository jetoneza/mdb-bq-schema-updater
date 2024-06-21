package bigquery

import (
	"context"
	"log"

	bq "cloud.google.com/go/bigquery"
	"github.com/jetoneza/mdb-bg-schema-updater/pkg/config"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type StreamWriter struct {
	Client  *bq.Client
	Context context.Context
}

type SchemaField struct {
	Name string
	Type bq.FieldType
}

func NewStreamWriter(ctx context.Context, creds *google.Credentials) *StreamWriter {
	var client *bq.Client
	var err error

	if creds != nil {
		client, err = bq.NewClient(ctx, config.BQ_PROJECT_ID, option.WithCredentials(creds))
	} else {
		client, err = bq.NewClient(ctx, config.BQ_PROJECT_ID)
	}

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
func (w *StreamWriter) UpdateSchema(schema map[string]interface{}) error {
	sampleSchema, err := createDynamicSchema(schema)
	if err != nil {
		panic(err)
	}

	log.Println("Schema successfully created")
	logSchema(sampleSchema)

	tableRef := w.Client.Dataset(config.BQ_DATASET).Table(config.BQ_TABLE)
	meta, err := tableRef.Metadata(w.Context)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}

	update := bq.TableMetadataToUpdate{
		Schema: sampleSchema,
	}

	res, err := tableRef.Update(w.Context, update, meta.ETag)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}

	log.Println("Schema successfully updated: ", res.FullID)

	return nil
}

func (w *StreamWriter) Close() {
	log.Println("Closing BigQuery stream writer...")
	if w.Client != nil {
		w.Client.Close()
	}
}

// Helpers
func createDynamicSchema(schema map[string]interface{}) (bq.Schema, error) {
	sampleSchema, err := parseSchemaInput(schema)
	if err != nil {
		return nil, err
	}

	return sampleSchema, nil
}

func parseSchemaInput(schema map[string]interface{}) (bq.Schema, error) {
	var sampleSchema bq.Schema
	for key, value := range schema {
		fieldSchema := &bq.FieldSchema{Name: key, Type: determineFieldType(value)}
		sampleSchema = append(sampleSchema, fieldSchema)
	}
	return sampleSchema, nil
}

func determineFieldType(value interface{}) bq.FieldType {
	switch value.(type) {
	case string:
		return bq.StringFieldType
	case int32, int64, float32, float64:
		return bq.IntegerFieldType
	case bool:
		return bq.BooleanFieldType
	case []interface{}:
		return bq.RecordFieldType // Adjust as needed for nested structures
	default:
		return bq.StringFieldType // Default type, adjust as needed
	}
}

// logSchema logs the created BigQuery schema
func logSchema(schema bq.Schema) {
	for _, field := range schema {
		log.Printf("Field: %s, Type: %s\n", field.Name, field.Type)
	}
}
