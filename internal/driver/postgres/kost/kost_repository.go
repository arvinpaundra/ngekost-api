package kost

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/util/db"
	"gorm.io/gorm"
)

type kostRepository struct {
	db *gorm.DB
}

func NewKostRepository(db *gorm.DB) contract.KostRepository {
	return &kostRepository{db: db}
}

func (k *kostRepository) Save(ctx context.Context, kost *entity.Kost) error {
	err := k.db.WithContext(ctx).Model(&entity.Kost{}).Create(&kost).Error
	if err != nil {
		return err
	}
	return nil
}

func (k *kostRepository) SaveWithTx(ctx context.Context, tx *gorm.DB, kost *entity.Kost) error {
	err := tx.WithContext(ctx).Model(&entity.Kost{}).Create(&kost).Error
	if err != nil {
		return err
	}
	return nil
}

func (k *kostRepository) Update(ctx context.Context, kost *entity.Kost, kostId string) error {
	err := k.db.WithContext(ctx).Model(&entity.Kost{}).Where("id = ?", kostId).Updates(&kost).Error
	if err != nil {
		return err
	}
	return nil
}

func (k *kostRepository) UpdateWithTx(ctx context.Context, tx *gorm.DB, kost *entity.Kost, kostId string) error {
	err := tx.WithContext(ctx).Model(&entity.Kost{}).Where("id = ?", kostId).Updates(&kost).Error
	if err != nil {
		return err
	}
	return nil
}

func (k *kostRepository) Delete(ctx context.Context, kostId string) error {
	err := k.db.WithContext(ctx).Model(&entity.Kost{}).Where("id = ?", kostId).Delete(&entity.Kost{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (k *kostRepository) DeleteWithTx(ctx context.Context, tx *gorm.DB, kostId string) error {
	err := tx.WithContext(ctx).Model(&entity.Kost{}).Where("id = ?", kostId).Delete(&entity.Kost{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (k *kostRepository) Find(ctx context.Context, query *request.Common) ([]*entity.Kost, error) {
	var kosts []*entity.Kost

	err := k.db.WithContext(ctx).Model(&entity.Kost{}).Preload("Owner").
		Joins("JOIN owners ON kosts.owner_id = owners.id").
		Scopes(
			db.Search("kosts.name", query.Search),
		).
		Order("kosts.created_at desc").
		Offset(query.GetOffset()).Limit(query.GetLimit()).
		Find(&kosts).Error

	if err != nil {
		return nil, err
	}

	return kosts, nil
}

func (k *kostRepository) FindByOwnerId(ctx context.Context, ownerId string, query *request.Common) ([]*entity.Kost, error) {
	var kosts []*entity.Kost

	err := k.db.WithContext(ctx).Model(&entity.Kost{}).Preload("Owner").
		Joins("JOIN owners ON kosts.owner_id = owners.id").
		Where("kosts.owner_id = ?", ownerId).
		Scopes(
			db.Search("kosts.name", query.Search),
		).
		Order("kosts.created_at desc").
		Offset(query.GetOffset()).Limit(query.GetLimit()).
		Find(&kosts).Error

	if err != nil {
		return nil, err
	}

	return kosts, nil
}

func (k *kostRepository) FindById(ctx context.Context, kostId string) (*entity.Kost, error) {
	var kost entity.Kost

	err := k.db.WithContext(ctx).Model(&entity.Kost{}).Preload("Owner").
		Where("id = ?", kostId).
		First(&kost).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrKostNotFound
		}
		return nil, err
	}

	return &kost, nil
}

func (k *kostRepository) Count(ctx context.Context, query *request.Common) (int, error) {
	var total int64

	err := k.db.WithContext(ctx).Model(&entity.Kost{}).
		Joins("JOIN owners ON kosts.owner_id = owners.id").
		Scopes(
			db.Search("kosts.name", query.Search),
		).
		Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), nil
}

func (k *kostRepository) CountByOwnerId(ctx context.Context, ownerId string, query *request.Common) (int, error) {
	var total int64

	err := k.db.WithContext(ctx).Model(&entity.Kost{}).
		Joins("JOIN owners ON kosts.owner_id = owners.id").
		Where("kosts.owner_id = ?", ownerId).
		Scopes(
			db.Search("kosts.name", query.Search),
		).
		Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), nil
}
