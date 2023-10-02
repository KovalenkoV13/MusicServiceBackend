package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	client *redis.Client
}

func New(ctx context.Context, addr, pass string, db int) (*Client, error) {
	redisCl := &Client{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pass,
			DB:       db,
		}),
	}
	if _, err := redisCl.client.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("cant ping redis: %w", err)
	}

	return redisCl, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}
