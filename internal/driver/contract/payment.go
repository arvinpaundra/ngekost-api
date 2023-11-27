package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *entity.Payment) error
	SaveWithTx(ctx context.Context, tx *gorm.DB, payment *entity.Payment) error
	Update(ctx context.Context, paymentId string, payment *entity.Payment) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, paymentId string, payment *entity.Payment) error
	FindAll(ctx context.Context, query *request.QueryParamPayment) ([]*entity.Payment, error)
	FindById(ctx context.Context, paymentId string) (*entity.Payment, error)
	Count(ctx context.Context, query *request.QueryParamPayment) (int, error)
}
