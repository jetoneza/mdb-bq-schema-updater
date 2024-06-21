package mongodb

import (
	"context"
	"log"

	"github.com/jetoneza/mdb-bg-schema-updater/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client  *mongo.Client
	Context context.Context
}

func New(ctx context.Context) *MongoDB {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.MONGO_DB_URI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return &MongoDB{
		Client:  client,
		Context: ctx,
	}
}

func (m *MongoDB) Close() {
	err := m.Client.Disconnect(m.Context)
	if err != nil {
		panic(err)
	}
}

func (m *MongoDB) GetSchemaFromMongoDB(ctx context.Context) (map[string]interface{}, error) {
	coll := m.Client.Database(config.MONGO_DB_NAME).Collection(config.MONGO_DB_COLLECTION)

	var result map[string]interface{}
	err := coll.FindOne(ctx, bson.D{}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
