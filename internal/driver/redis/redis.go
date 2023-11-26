package redis

import (
	"context"
	// "crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arvinpaundra/ngekost-api/pkg/util/config"
	l "github.com/arvinpaundra/ngekost-api/pkg/util/log"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	user     string
	password string
	host     string
	port     int
}

func New() *Redis {
	return &Redis{
		user:     config.GetString("REDIS_USER"),
		password: config.GetString("REDIS_PASSWORD"),
		host:     config.GetString("REDIS_HOST"),
		port:     config.GetInt("REDIS_PORT"),
	}
}

func (r *Redis) Connect(ctx context.Context) *redis.Client {
	addr := fmt.Sprintf("%s:%d", r.host, r.port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: r.user,
		Password: r.password,
		DB:       0,
		// TLSConfig: &tls.Config{},
	})

	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if _, err := rdb.Ping(timeout).Result(); err != nil {
		l.Logging().Error(err.Error())
		os.Exit(1)
	}

	log.Println("connected to redis")

	return rdb
}
