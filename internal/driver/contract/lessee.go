package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"gorm.io/gorm"
)

type LesseeRepository interface {
	Save(ctx context.Context, lessee *entity.Lessee) error
	Update(ctx context.Context, lessee *entity.Lessee, ownerId string) error
	SaveWithTx(ctx context.Context, tx *gorm.DB, lessee *entity.Lessee) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, lessee *entity.Lessee, ownerId string) error
	Find(ctx context.Context, keyword string) ([]*entity.Lessee, error)
	FindById(ctx context.Context, id string) (*entity.Lessee, error)
	Count(ctx context.Context, keyword string) (int, error)
}
