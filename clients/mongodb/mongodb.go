package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(dbHost string, dbName string) (*mongo.Database, *context.Context) {
	ctx := context.Background()

	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
		NilSliceAsEmpty:   true,
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	mongoClientOptions := options.Client()
	mongoClientOptions.ApplyURI(dbHost).SetBSONOptions(bsonOpts).SetServerAPIOptions(serverAPI)

	mongoClient, err := mongo.Connect(ctx, mongoClientOptions)
	if err != nil {
		panic(err)
	}

	if err := mongoClient.Ping(ctx, nil); err != nil {
		panic(err)
	}

	log.Println("mongodb database connected successfully")

	return mongoClient.Database(dbName), &ctx
}
