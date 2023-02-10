package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func New(opts []Option) (redis.UniversalClient, func(), error) {
	op := redisOptions{
		Addrs:    []string{"localhost:6379"},
		Password: "",
		DB:       0,
	}
	for _, option := range opts {
		option.apply(&op)
	}
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    op.Addrs,
		Password: op.Password,
		DB:       op.DB,
		PoolSize: op.PoolSize,
	})

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		client.Close()
	}
	return client, cleanup, nil
}
