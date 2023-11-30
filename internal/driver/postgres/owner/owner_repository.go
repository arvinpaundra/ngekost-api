package owner

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"gorm.io/gorm"
)

type ownerRepository struct {
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) contract.OwnerRepository {
	return &ownerRepository{db: db}
}

func (o *ownerRepository) Save(ctx context.Context, owner *entity.Owner) error {
	err := o.db.WithContext(ctx).Model(&entity.Owner{}).Create(&owner).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *ownerRepository) Update(ctx context.Context, owner *entity.Owner, ownerId string) error {
	err := o.db.WithContext(ctx).Model(&entity.Owner{}).Where("id = ?", ownerId).Updates(&owner).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *ownerRepository) Find(ctx context.Context, keyword string) ([]*entity.Owner, error) {
	var owners []*entity.Owner

	err := o.db.WithContext(ctx).Model(&entity.Owner{}).
		Joins("JOIN users ON owners.user_id = users.id").
		Scopes(
			func(db *gorm.DB) *gorm.DB {
				if keyword != "" {
					return db.Where(
						db.Where("owners.fullname LIKE ?", "%"+keyword+"%").
							Or("users.username LIKE ?", "%"+keyword+"%"),
					)
				}
				return db
			},
		).
		Find(&owners).Error

	if err != nil {
		return nil, err
	}

	return owners, nil
}

func (o *ownerRepository) FindById(ctx context.Context, ownerId string) (*entity.Owner, error) {
	var owner entity.Owner

	err := o.db.WithContext(ctx).Model(&entity.Owner{}).
		Joins("JOIN users ON owners.user_id = users.id").
		Where("owners.id = ?", ownerId).
		First(&owner).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrOwnerNotFound
		}
		return nil, err
	}

	return &owner, nil
}

func (o *ownerRepository) Count(ctx context.Context, keyword string) (int, error) {
	var total int64

	err := o.db.WithContext(ctx).Model(&entity.Owner{}).
		Joins("JOIN users ON owners.user_id = users.id").
		Scopes(
			func(db *gorm.DB) *gorm.DB {
				if keyword != "" {
					return db.Where(
						db.Where("owners.fullname LIKE ?", "%"+keyword+"%").Or("users.username LIKE ?", "%"+keyword+"%"),
					)
				}
				return db
			},
		).
		Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), nil
}
