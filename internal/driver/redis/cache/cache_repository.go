package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/pkg/util/log"
	"github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(client *redis.Client) contract.CacheRepository {
	return &cacheRepository{client: client}
}

func (c *cacheRepository) Save(ctx context.Context, key string, value any, ttl time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		log.Logging(err.Error()).Error()
		return err
	}

	err = c.client.Set(ctx, key, v, ttl).Err()
	if err != nil {
		log.Logging(err.Error()).Error()
		return err
	}

	return nil
}

func (c *cacheRepository) Get(ctx context.Context, key string) (*string, error) {
	res, err := c.client.Get(ctx, key).Result()

	if err != nil {
		log.Logging(err.Error()).Error()
		return nil, err
	}

	if err == redis.Nil {
		log.Logging(err.Error()).Error()
		return nil, err
	}

	return &res, nil
}

func (c *cacheRepository) Del(ctx context.Context, keys ...string) error {
	err := c.client.Del(ctx, keys...).Err()
	if err != nil {
		log.Logging(err.Error()).Error()
		return err
	}

	return nil
}
