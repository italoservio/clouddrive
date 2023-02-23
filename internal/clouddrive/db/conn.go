package db

import (
	"context"
	"os"
	"time"

	"github.com/italoservio/clouddrive/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var UsersDatabase *mongo.Database

const (
	USERS_COLLECTION = "users"
)

func SetupConnection() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		logger.Error("Environment variable 'MONGODB_URI' not found")
		os.Exit(1)
	}

	ctx, cancel := Context()
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	Client = client
	UsersDatabase = client.Database("users_db")
}

func Context() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	return ctx, cancel
}
