package transaction

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/midtrans/midtrans-go/snap"
)

type transactionRepository struct {
	sc *snap.Client
}

func NewTransactionRepository(sc *snap.Client) contract.MidtransTransactionRepository {
	return &transactionRepository{sc: sc}
}

func (t *transactionRepository) Create(ctx context.Context, transaction *snap.Request) (*snap.Response, error) {
	res, err := t.sc.CreateTransaction(transaction)

	if err != nil {
		return nil, err
	}

	return res, nil
}
