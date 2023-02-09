package lungo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connection *mongo.Client

func Connect() error {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configuration.URI))

	if err != nil {
		log.Fatal(err)

		return err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)

		return err
	}

	connection = client

	return nil
}

func Disconnect() error {
	ctx := context.Background()

	if err := connection.Disconnect(ctx); err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}
