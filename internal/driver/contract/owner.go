package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/entity"
)

type OwnerRepository interface {
	Save(ctx context.Context, owner *entity.Owner) error
	Update(ctx context.Context, owner *entity.Owner, ownerId string) error
	Find(ctx context.Context, keyword string) ([]*entity.Owner, error)
	FindById(ctx context.Context, ownerId string) (*entity.Owner, error)
	Count(ctx context.Context, keyword string) (int, error)
}
