package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/url"
)

type Redis struct {
	host        string
	username    string
	password    string
	port        string
	dbName      string
	RedisClient *redis.Client
}

func New(opts ...Option) (*Redis, error) {
	r := new(Redis)

	for _, opt := range opts {
		opt(r)
	}

	q := url.Values{}

	u := url.URL{
		Scheme:   "redis",
		Host:     fmt.Sprintf("%s:%s", r.host, r.port),
		Path:     r.dbName,
		RawQuery: q.Encode(),
	}

	opt, err := redis.ParseURL(u.String())
	if err != nil {
		log.Fatalf("Error creating Redis client config: %v", err)
		return nil, err
	}

	client := redis.NewClient(opt)

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
		return nil, err
	}

	r.RedisClient = client

	return r, nil
}
