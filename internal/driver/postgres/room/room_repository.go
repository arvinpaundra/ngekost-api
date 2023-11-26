package room

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/util/db"
	"gorm.io/gorm"
)

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) contract.RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) Save(ctx context.Context, room *entity.Room) error {
	err := r.db.WithContext(ctx).Model(&entity.Room{}).Create(&room).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roomRepository) SaveWithTx(ctx context.Context, tx *gorm.DB, room *entity.Room) error {
	err := tx.WithContext(ctx).Model(&entity.Room{}).Create(&room).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roomRepository) Update(ctx context.Context, room *entity.Room, roomId string) error {
	err := r.db.WithContext(ctx).Model(&entity.Room{}).Where("id = ?", roomId).Updates(&room).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roomRepository) UpdateWithTx(ctx context.Context, tx *gorm.DB, room *entity.Room, roomId string) error {
	err := tx.WithContext(ctx).Model(&entity.Room{}).Where("id = ?", roomId).Updates(&room).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roomRepository) Delete(ctx context.Context, roomId string) error {
	err := r.db.WithContext(ctx).Model(&entity.Room{}).Where("id = ?", roomId).Delete(&entity.Room{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roomRepository) DeleteWithTx(ctx context.Context, tx *gorm.DB, roomId string) error {
	err := tx.WithContext(ctx).Model(&entity.Room{}).Where("id = ?", roomId).Delete(&entity.Room{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roomRepository) FindByKostId(ctx context.Context, kostId string, query *request.Common) ([]*entity.Room, error) {
	var rooms []*entity.Room

	err := r.db.WithContext(ctx).Model(&entity.Room{}).
		Where("kost_id = ?", kostId).
		Scopes(
			db.Search("name", query.Search),
		).
		Order("price asc").
		Offset(query.GetOffset()).Limit(query.GetLimit()).
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *roomRepository) FindById(ctx context.Context, roomId string) (*entity.Room, error) {
	var room entity.Room

	err := r.db.WithContext(ctx).Model(&entity.Room{}).Preload("Assets").
		Where("id = ?", roomId).
		First(&room).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrRoomNotFound
		}
		return nil, err
	}

	return &room, nil
}

func (r *roomRepository) CountByKostId(ctx context.Context, kostId string, query *request.Common) (int, error) {
	var total int64

	err := r.db.WithContext(ctx).Model(&entity.Room{}).
		Where("kost_id = ?", kostId).
		Scopes(
			db.Search("name", query.Search),
		).
		Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), nil
}
