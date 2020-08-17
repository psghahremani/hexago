package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"hexago/domain/ports"
	"time"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, timeout time.Duration) (*mongo.Client, error) {
	// Make a context based on timeout.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Create a client using given URL.
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	// Ping database to ensure connection.
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL, database string, timeout time.Duration) (ports.MongoRepository, error) {
	// Create a MongoDB Client using given URL.
	client, err := newMongoClient(mongoURL, timeout)
	if err != nil {
		return nil, err
	}

	// Create and return repository.
	repository := &mongoRepository{
		client:   client,
		database: database,
		timeout:  timeout,
	}
	return repository, nil
}

/* TODO: Implement repository methods on *mongoRepository as they are declared in "hexago/domain/ports" here. */
