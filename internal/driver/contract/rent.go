package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
)

type RentRepository interface {
	Save(ctx context.Context, rent *entity.Rent) error
	Update(ctx context.Context, rentId string, rent *entity.Rent) error
	Delete(ctx context.Context, rentId string) error
	CheckExistRent(ctx context.Context, lesseeId string) (bool, error)
	FindById(ctx context.Context, rentId string) (*entity.Rent, error)
	FindByLesseeId(ctx context.Context, lesseId string, query *request.Common) ([]*entity.Rent, error)
	FindByKostId(ctx context.Context, kostId string, query *request.Common) ([]*entity.Rent, error)
	CountByLesseeId(ctx context.Context, lesseeId string, query *request.Common) (int, error)
	CountByKostId(ctx context.Context, kostId string, query *request.Common) (int, error)
}
