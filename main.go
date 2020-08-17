package main

import (
	"hexago/adapters/infrastructure/mongo"
	"hexago/adapters/infrastructure/redis"
	"hexago/config"
	"log"
)

func main() {

	// Initialize driven adapters (infrastructure).
	mongoRepository, err := mongo.NewMongoRepository(config.MongoURL, config.MongoDatabaseName, config.MongoTimeout)
	if err != nil {
		log.Fatalf("cannot initialize MongoDB adapter: %v", err)
		return
	}
	_ = mongoRepository

	redisRepository, err := redis.NewRedisRepository(config.RedisURL)
	if err != nil {
		log.Fatalf("cannot initialize Redis adapter: %v", err)
		return
	}
	_ = redisRepository

	// Initialize domain services.

	// Initialize driving adapters (interface).
}
