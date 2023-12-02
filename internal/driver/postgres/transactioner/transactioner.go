package transactioner

import (
	"context"
	"database/sql"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/bill"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/kost"
	kostrule "github.com/arvinpaundra/ngekost-api/internal/driver/postgres/kostRule"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/lessee"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/owner"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/payment"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/rent"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/room"
	roomasset "github.com/arvinpaundra/ngekost-api/internal/driver/postgres/roomAsset"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/session"
	"github.com/arvinpaundra/ngekost-api/internal/driver/postgres/user"
	"gorm.io/gorm"
)

type transactioner struct {
	db *gorm.DB
}

func NewTransactioner(db *gorm.DB) contract.Transactioner {
	return &transactioner{db: db}
}

func (t *transactioner) Begin(ctx context.Context, opts ...*sql.TxOptions) (contract.Transaction, error) {
	tx := t.db.Begin(opts...)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &transaction{db: tx}, nil
}

type transaction struct {
	db *gorm.DB
}

func (t *transaction) Commit() error {
	tx := t.db.Commit()
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (t *transaction) Rollback() error {
	tx := t.db.Rollback()
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (t *transaction) UserRepository() contract.UserRepository {
	return user.NewAuthRepository(t.db)
}

func (t *transaction) OwnerRepository() contract.OwnerRepository {
	return owner.NewOwnerRepository(t.db)
}

func (t *transaction) LesseeRepository() contract.LesseeRepository {
	return lessee.NewLesseeRepository(t.db)
}

func (t *transaction) SessionRepository() contract.SessionRepository {
	return session.NewSessionRepository(t.db)
}

func (t *transaction) KostRepository() contract.KostRepository {
	return kost.NewKostRepository(t.db)
}

func (t *transaction) RoomRepository() contract.RoomRepository {
	return room.NewRoomRepository(t.db)
}

func (t *transaction) RoomAssetRepository() contract.RoomAssetRepository {
	return roomasset.NewRoomAssetRepository(t.db)
}

func (t *transaction) KostRuleRepository() contract.KostRuleRepository {
	return kostrule.NewKostRuleRepository(t.db)
}

func (t *transaction) BillRepository() contract.BillRepository {
	return bill.NewBillRepository(t.db)
}

func (t *transaction) PaymentRepository() contract.PaymentRepository {
	return payment.NewPaymentRepository(t.db)
}

func (t *transaction) RentRepository() contract.RentRepository {
	return rent.NewRentRepository(t.db)
}
