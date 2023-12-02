package lessee

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"gorm.io/gorm"
)

type lesseeRepository struct {
	db *gorm.DB
}

func NewLesseeRepository(db *gorm.DB) contract.LesseeRepository {
	return &lesseeRepository{db: db}
}

func (l *lesseeRepository) Save(ctx context.Context, lessee *entity.Lessee) error {
	err := l.db.WithContext(ctx).Model(&entity.Lessee{}).Create(&lessee).Error
	if err != nil {
		return err
	}

	return nil
}

func (l *lesseeRepository) Update(ctx context.Context, lessee *entity.Lessee, lesseeId string) error {
	err := l.db.WithContext(ctx).Model(&entity.Lessee{}).Where("id = ?", lesseeId).Updates(&lessee).Error
	if err != nil {
		return err
	}

	return nil
}

func (l *lesseeRepository) Find(ctx context.Context, keyword string) ([]*entity.Lessee, error) {
	var lessees []*entity.Lessee

	err := l.db.WithContext(ctx).Model(&entity.Lessee{}).
		Joins("JOIN users ON lessees.user_id = users.id").
		Scopes(
			func(db *gorm.DB) *gorm.DB {
				if keyword != "" {
					return db.Where(
						db.Where("lessees.fullname LIKE ?", "%"+keyword+"%").Or("users.username LIKE ?", "%"+keyword+"%"),
					)
				}
				return db
			},
		).
		Find(&lessees).Error

	if err != nil {
		return nil, err
	}

	return lessees, nil
}

func (l *lesseeRepository) FindById(ctx context.Context, lesseId string) (*entity.Lessee, error) {
	var lessee entity.Lessee

	err := l.db.WithContext(ctx).Model(&entity.Lessee{}).
		Joins("JOIN users ON lessees.user_id = users.id").
		Where("lessees.id = ?", lesseId).
		First(&lessee).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrLesseeNotFound
		}
		return nil, err
	}

	return &lessee, nil
}

func (l *lesseeRepository) Count(ctx context.Context, keyword string) (int, error) {
	var total int64

	err := l.db.WithContext(ctx).Model(&entity.Lessee{}).
		Joins("JOIN users ON lessees.user_id = users.id").
		Scopes(
			func(db *gorm.DB) *gorm.DB {
				if keyword != "" {
					return db.Where(
						db.Where("lessees.fullname LIKE ?", "%"+keyword+"%").Or("users.username LIKE ?", "%"+keyword+"%"),
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
