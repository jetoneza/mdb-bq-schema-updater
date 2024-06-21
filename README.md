# MongoDB to BigQuery Schema Updater

Updates BigQuery schema based on the schema of a MongoDB collection.

## Prerequisites

* Setup [Go](https://go.dev/doc/install)
* BigQuery authentication is via [Google Application Default Credentials](https://cloud.google.com/docs/authentication/provide-credentials-adc#how-to).

## Instructions

* Create .env file with the following variables:

```
MONGO_DB_URI=mongodb+srv://<user>:<password>@<host>
MONGO_DB_NAME=<db_name>
MONGO_DB_COLLECTION=<collection_name>
BQ_PROJECT_ID=<bigquery_project_id>
BQ_DATASET=<bigquery_dataset>
BQ_TABLE=<bigquery_table>
```

* Run the command:

```
go run .
```

