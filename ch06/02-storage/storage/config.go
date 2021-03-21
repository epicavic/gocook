package storage

import (
	"context"
	"log"
	"time"

	"github.com/sethvargo/go-retry"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoStorage initializes a MongoStorage
func NewMongoStorage(ctx context.Context, dbname, dbcoll string) (*MongoStorage, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	dburl := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, dburl)
	if err != nil {
		return nil, err
	}

	// retry database connection
	log.Println("Trying to connect to database")
	err = retry.Constant(ctx, 5*time.Second, func(ctx context.Context) error {
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			// This marks the error as retryable
			return retry.RetryableError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database")

	ms := MongoStorage{
		Client:     client,
		DB:         dbname,
		Collection: dbcoll,
	}
	return &ms, nil
}
