package redis

import (
	"github.com/go-redis/redis"
	"hexago/domain/ports"
)

type redisRepository struct {
	// TODO: DEV: Add timeout for Redis too?
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	// Parse URL.
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	// Create new Redis client.
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepository(redisURL string) (ports.RedisRepository, error) {
	// Create a Redis client using given URL.
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, err
	}

	// Create and return repository.
	repo := &redisRepository{
		client: client,
	}
	return repo, nil
}

/* TODO: Implement repository methods on *mongoRepository as they are declared in "hexago/domain/ports" here. */
