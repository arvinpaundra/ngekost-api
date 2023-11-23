package contract

import (
	"context"
	"time"
)

type CacheRepository interface {
	Save(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string) (*string, error)
	Del(ctx context.Context, keys ...string) error
}
