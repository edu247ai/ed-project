package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Client struct {
	rdb *redis.Client
}

func New(addr string) *Client {
	return &Client{
		rdb: redis.NewClient(&redis.Options{
			Addr:         addr,
			DB:           0,
			PoolSize:     20,
			MinIdleConns: 5,
		}),
	}
}

//
// CACHE
//

func (c *Client) SetCache(key string, value []byte, ttl time.Duration) error {
	return c.rdb.Set(ctx, key, value, ttl).Err()
}

func (c *Client) GetCache(key string) ([]byte, error) {
	return c.rdb.Get(ctx, key).Bytes()
}

//
// QUEUE (FIFO)
//

func (c *Client) Enqueue(queue string, payload []byte) error {
	return c.rdb.RPush(ctx, queue, payload).Err()
}

func (c *Client) Dequeue(queue string, block bool, timeout time.Duration) ([]byte, error) {
	if block {
		res, err := c.rdb.BLPop(ctx, timeout, queue).Result()
		if err != nil {
			return nil, err
		}
		return []byte(res[1]), nil
	}
	return c.rdb.LPop(ctx, queue).Bytes()
}
