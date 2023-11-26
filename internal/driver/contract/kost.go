package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"gorm.io/gorm"
)

type KostRepository interface {
	Save(ctx context.Context, kost *entity.Kost) error
	SaveWithTx(ctx context.Context, tx *gorm.DB, kost *entity.Kost) error
	Update(ctx context.Context, kost *entity.Kost, kostId string) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, kost *entity.Kost, kostId string) error
	Delete(ctx context.Context, kostId string) error
	DeleteWithTx(ctx context.Context, tx *gorm.DB, kostId string) error
	Find(ctx context.Context, query *request.Common) ([]*entity.Kost, error)
	FindByOwnerId(ctx context.Context, ownerId string, query *request.Common) ([]*entity.Kost, error)
	FindById(ctx context.Context, kostId string) (*entity.Kost, error)
	Count(ctx context.Context, query *request.Common) (int, error)
	CountByOwnerId(ctx context.Context, ownerId string, query *request.Common) (int, error)
}
