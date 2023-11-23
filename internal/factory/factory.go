package factory

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres"
	txBeginner "github.com/arvinpaundra/ngekost-api/internal/driver/postgres/beginner"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/lessee"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/owner"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/session"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/user"
	"github.com/arvinpaundra/ngekost-api/internal/driver/redis"
	"github.com/arvinpaundra/ngekost-api/internal/driver/redis/cache"
	"github.com/arvinpaundra/ngekost-api/pkg/util/token"
)

type Factory struct {
	CacheRepository   contract.CacheRepository
	TxBeginner        contract.TxBeginner
	UserRepository    contract.UserRepository
	OwnerRepository   contract.OwnerRepository
	LesseeRepository  contract.LesseeRepository
	SessionRepository contract.SessionRepository
	JSONWebToken      token.JSONWebToken
}

func NewFactory(ctx context.Context) *Factory {
	rdb := redis.New().Connect(ctx)
	pg := postgres.New().Connect(ctx)

	return &Factory{
		JSONWebToken:      token.NewJWT(),
		CacheRepository:   cache.NewCacheRepository(rdb),
		TxBeginner:        txBeginner.NewTxBeginner(pg),
		UserRepository:    user.NewAuthRepository(pg),
		OwnerRepository:   owner.NewOwnerRepository(pg),
		LesseeRepository:  lessee.NewLesseeRepository(pg),
		SessionRepository: session.NewSessionRepository(pg),
	}
}
