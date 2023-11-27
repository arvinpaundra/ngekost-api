package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/adapter/response"
	"github.com/midtrans/midtrans-go/snap"
)

type (
	MidtransTransactionRepository interface {
		Create(ctx context.Context, transaction *snap.Request) (*snap.Response, error)
	}

	MidtransBeneficiaryRepository interface {
		Create(ctx context.Context, beneficiary *request.CreateBeneficiary) error
		Update(ctx context.Context, aliasName string, beneficiary *request.UpdateBeneficiary) error
		FindAll(ctx context.Context, query *request.Common) ([]*response.MidtransBeneficiary, error)
	}

	MidtransPayoutRepository interface {
		Create(ctx context.Context, payout *request.CreatePayout) ([]*response.MidtransPayout, error)
		Approve(ctx context.Context, payout *request.ApprovePayout) error
		Reject(ctx context.Context, payout *request.RejectPayout) error
	}
)
