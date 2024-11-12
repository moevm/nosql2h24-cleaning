package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const mongodb_uri = "mongodb://%s:%s@%s:%d"

func MongoURI(user, password, hostname string, port int) string {
	return fmt.Sprintf(mongodb_uri, user, password, hostname, port)
}

func Connect(
	user string,
	password string,
	hostname string,
	port int,
) (*mongo.Client, error) {
	uri := MongoURI(user, password, hostname, port)

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
