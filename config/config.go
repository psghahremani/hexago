package config

import (
	"time"
)

// MongoDB configurations.
const MongoHost = "172.17.0.2"
const MongoPort = "27017"
const MongoURL = "mongodb://" + MongoHost + ":" + MongoPort
const MongoTimeout = 5 * time.Second
const MongoDatabaseName = "hexago"

const RedisHost = "172.17.0.3"
const RedisPort = "6379"
const RedisURL = "redis://" + RedisHost + ":" + RedisPort
