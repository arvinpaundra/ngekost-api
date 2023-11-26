package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"gorm.io/gorm"
)

type RoomRepository interface {
	Save(ctx context.Context, room *entity.Room) error
	SaveWithTx(ctx context.Context, tx *gorm.DB, room *entity.Room) error
	Update(ctx context.Context, room *entity.Room, roomId string) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, room *entity.Room, roomId string) error
	Delete(ctx context.Context, roomId string) error
	DeleteWithTx(ctx context.Context, tx *gorm.DB, roomId string) error
	FindByKostId(ctx context.Context, kostId string, query *request.Common) ([]*entity.Room, error)
	FindById(ctx context.Context, roomId string) (*entity.Room, error)
	CountByKostId(ctx context.Context, kostId string, query *request.Common) (int, error)
}
