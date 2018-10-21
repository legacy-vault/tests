package redis

import (
	"github.com/go-redis/redis"
)

// Connects to Redis Server.
func Connect() (*redis.Client, error) {

	var client *redis.Client
	var err error

	client, err = Connect2()

	return client, err
}

// Connect2.
func Connect2() (*redis.Client, error) {

	var client *redis.Client
	var err error
	var options redis.Options

	options.Addr = "localhost:6379"

	client = redis.NewClient(&options)

	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
