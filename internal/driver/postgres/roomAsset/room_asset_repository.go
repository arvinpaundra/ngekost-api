package roomasset

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"gorm.io/gorm"
)

type roomAssetRepository struct {
	db *gorm.DB
}

func NewRoomAssetRepository(db *gorm.DB) contract.RoomAssetRepository {
	return &roomAssetRepository{db: db}
}

func (r *roomAssetRepository) Save(ctx context.Context, asset *entity.RoomAsset) error {
	err := r.db.WithContext(ctx).Model(&entity.RoomAsset{}).Create(&asset).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roomAssetRepository) SaveWithTx(ctx context.Context, tx *gorm.DB, asset *entity.RoomAsset) error {
	err := tx.WithContext(ctx).Model(&entity.RoomAsset{}).Create(&asset).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roomAssetRepository) Update(ctx context.Context, asset *entity.RoomAsset, assetId string) error {
	err := r.db.WithContext(ctx).Model(&entity.RoomAsset{}).Where("id = ?", assetId).Updates(&asset).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roomAssetRepository) UpdateWithTx(ctx context.Context, tx *gorm.DB, asset *entity.RoomAsset, assetId string) error {
	err := tx.WithContext(ctx).Model(&entity.RoomAsset{}).Where("id = ?", assetId).Updates(&asset).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roomAssetRepository) Delete(ctx context.Context, assetId string) error {
	err := r.db.WithContext(ctx).Model(&entity.RoomAsset{}).
		Where("id = ?", assetId).Delete(&entity.RoomAsset{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *roomAssetRepository) DeleteWithTx(ctx context.Context, tx *gorm.DB, assetId string) error {
	err := tx.WithContext(ctx).Model(&entity.RoomAsset{}).
		Where("id = ?", assetId).Delete(&entity.RoomAsset{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *roomAssetRepository) FindById(ctx context.Context, assetId string) (*entity.RoomAsset, error) {
	var asset entity.RoomAsset

	err := r.db.WithContext(ctx).Model(&entity.RoomAsset{}).
		Where("id = ?", assetId).
		First(&asset).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrRoomAssetNotFound
		}
		return nil, err
	}

	return &asset, nil
}

func (r *roomAssetRepository) FindByRoomId(ctx context.Context, roomId string) ([]*entity.RoomAsset, error) {
	var assets []*entity.RoomAsset

	err := r.db.WithContext(ctx).Model(&entity.RoomAsset{}).
		Where("room_id = ?", roomId).
		Order("created_at asc").
		Find(&assets).Error

	if err != nil {
		return nil, err
	}

	return assets, nil
}
