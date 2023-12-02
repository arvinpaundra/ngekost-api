package contract

import (
	"context"
	"database/sql"
)

type Transactioner interface {
	Begin(ctx context.Context, opts ...*sql.TxOptions) (Transaction, error)
}

type Transaction interface {
	Commit() error
	Rollback() error

	// methods return repositories with transaction
	UserRepository() UserRepository
	OwnerRepository() OwnerRepository
	LesseeRepository() LesseeRepository
	SessionRepository() SessionRepository
	KostRepository() KostRepository
	RoomRepository() RoomRepository
	RoomAssetRepository() RoomAssetRepository
	KostRuleRepository() KostRuleRepository
	BillRepository() BillRepository
	PaymentRepository() PaymentRepository
	RentRepository() RentRepository
}
