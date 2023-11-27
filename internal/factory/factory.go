package factory

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/driver/midtrans"
	"github.com/arvinpaundra/ngekost-api/internal/driver/midtrans/transaction"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres"
	txBeginner "github.com/arvinpaundra/ngekost-api/internal/driver/postgres/beginner"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/bill"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/kost"
	kostrule "github.com/arvinpaundra/ngekost-api/internal/driver/postgres/kostRule"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/lessee"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/owner"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/payment"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/room"
	roomasset "github.com/arvinpaundra/ngekost-api/internal/driver/postgres/roomAsset"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/session"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/user"
	"github.com/arvinpaundra/ngekost-api/internal/driver/redis"
	"github.com/arvinpaundra/ngekost-api/internal/driver/redis/cache"
	"github.com/arvinpaundra/ngekost-api/pkg/util/token"
)

type Factory struct {
	JSONWebToken                  token.JSONWebToken
	CacheRepository               contract.CacheRepository
	MidtransTransactionRepository contract.MidtransTransactionRepository
	TxBeginner                    contract.TxBeginner
	UserRepository                contract.UserRepository
	OwnerRepository               contract.OwnerRepository
	LesseeRepository              contract.LesseeRepository
	SessionRepository             contract.SessionRepository
	KostRepository                contract.KostRepository
	RoomRepository                contract.RoomRepository
	KostRuleRepository            contract.KostRuleRepository
	RoomAssetRepository           contract.RoomAssetRepository
	BillRepository                contract.BillRepository
	PaymentRepository             contract.PaymentRepository
}

func NewFactory(ctx context.Context) *Factory {
	rdb := redis.New().Connect(ctx)
	pg := postgres.New().Connect(ctx)
	sc := midtrans.New().Snap()

	return &Factory{
		JSONWebToken:                  token.NewJWT(),
		CacheRepository:               cache.NewCacheRepository(rdb),
		MidtransTransactionRepository: transaction.NewTransactionRepository(sc),
		TxBeginner:                    txBeginner.NewTxBeginner(pg),
		UserRepository:                user.NewAuthRepository(pg),
		OwnerRepository:               owner.NewOwnerRepository(pg),
		LesseeRepository:              lessee.NewLesseeRepository(pg),
		SessionRepository:             session.NewSessionRepository(pg),
		KostRepository:                kost.NewKostRepository(pg),
		RoomRepository:                room.NewRoomRepository(pg),
		KostRuleRepository:            kostrule.NewKostRuleRepository(pg),
		RoomAssetRepository:           roomasset.NewRoomAssetRepository(pg),
		BillRepository:                bill.NewBillRepository(pg),
		PaymentRepository:             payment.NewPaymentRepository(pg),
	}
}
