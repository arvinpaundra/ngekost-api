package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"gorm.io/gorm"
)

type RoomAssetRepository interface {
	Save(ctx context.Context, asset *entity.RoomAsset) error
	SaveWithTx(ctx context.Context, tx *gorm.DB, asset *entity.RoomAsset) error
	Update(ctx context.Context, asset *entity.RoomAsset, assetId string) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, asset *entity.RoomAsset, assetId string) error
	Delete(ctx context.Context, assetId string) error
	DeleteWithTx(ctx context.Context, tx *gorm.DB, assetId string) error
	FindById(ctx context.Context, assetId string) (*entity.RoomAsset, error)
	FindByRoomId(ctx context.Context, roomId string) ([]*entity.RoomAsset, error)
}
