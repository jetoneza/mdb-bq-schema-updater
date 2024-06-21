package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	MONGO_DB_URI        = GetEnv("MONGO_DB_URI", "")
	MONGO_DB_NAME       = GetEnv("MONGO_DB_NAME", "")
	MONGO_DB_COLLECTION = GetEnv("MONGO_DB_COLLECTION", "")
	BQ_PROJECT_ID       = GetEnv("BQ_PROJECT_ID", "")
	BQ_DATASET          = GetEnv("BQ_DATASET", "")
	BQ_TABLE            = GetEnv("BQ_TABLE", "")

	// TODO: Use `ldflags` for a cleaner approach
	// VERSION=$(git tag --list 'v*' --sort=-v:refname | head -n 1 | sed 's/^v//')
	// go build -ldflags "-X github.com/jetrooper/mdb-bq-schema-updater/pkg/config.Version=${VERSION}" -v -o ./bin/api
	//
	// The issue currently of the above is that the CodeBuild environment
	// does not have access to the Git repository context
	VERSION = "0.1.0"
)

func GetEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		log.Printf("Environment variable found :: %v: %v", name, value)
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf("Environment variable not found :: %v", name))
}
