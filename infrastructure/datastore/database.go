package datastore

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewDatabase will create new database instance.
func NewDatabase(ctx context.Context, databaseURL string) (client *mongo.Client, err error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(databaseURL).
		SetServerAPIOptions(serverAPIOptions)
	client, err = mongo.Connect(ctx, clientOptions)

	return
}
