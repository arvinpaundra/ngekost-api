package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"gorm.io/gorm"
)

type BillRepository interface {
	Save(ctx context.Context, bill *entity.Bill) error
	SaveWithTx(ctx context.Context, tx *gorm.DB, bill *entity.Bill) error
	Update(ctx context.Context, billId string, bill *entity.Bill) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, billId string, bill *entity.Bill) error
	FindAll(ctx context.Context, query *request.QueryParamBill) ([]*entity.Bill, error)
	FindByLesseeId(ctx context.Context, lesseeId string, query *request.QueryParamBill) ([]*entity.Bill, error)
	FindById(ctx context.Context, billId string) (*entity.Bill, error)
	Count(ctx context.Context, query *request.QueryParamBill) (int, error)
	CountByLesseId(ctx context.Context, lesseeId string, query *request.QueryParamBill) (int, error)
}
