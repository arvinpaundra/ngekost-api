package rent

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/util/db"
	"gorm.io/gorm"
)

type rentRepository struct {
	db *gorm.DB
}

func NewRentRepository(db *gorm.DB) contract.RentRepository {
	return &rentRepository{db: db}
}

func (r *rentRepository) Save(ctx context.Context, rent *entity.Rent) error {
	err := r.db.WithContext(ctx).Model(&entity.Rent{}).Create(&rent).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *rentRepository) Update(ctx context.Context, rentId string, rent *entity.Rent) error {
	err := r.db.WithContext(ctx).Model(&entity.Rent{}).Where("id = ?", rentId).Updates(&rent).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *rentRepository) Delete(ctx context.Context, rentId string) error {
	err := r.db.WithContext(ctx).Model(&entity.Rent{}).Where("id = ?", rentId).Delete(&entity.Rent{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *rentRepository) CheckExistRent(ctx context.Context, lesseeId string) (bool, error) {
	var rents []*entity.Rent

	err := r.db.WithContext(ctx).Model(&entity.Rent{}).
		Where("end_date IS NULL").
		Where("lessee_id = ?", lesseeId).
		Find(&rents).Error

	if err != nil {
		return false, err
	}

	return len(rents) > 0, nil
}

func (r *rentRepository) FindById(ctx context.Context, rentId string) (*entity.Rent, error) {
	var rent entity.Rent

	err := r.db.WithContext(ctx).Model(&entity.Rent{}).Preload("Room.Kost").Preload("Lessee").
		Joins("JOIN lessees ON rents.lessee_id = lessees.id").
		Joins("JOIN rooms ON rents.room_id = rooms.id").
		Joins("JOIN kosts ON rooms.kost_id = kosts.id").
		Where("rents.id = ?", rentId).
		First(&rent).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrRentNotFound
		}
		return nil, err
	}

	return &rent, nil
}

func (r *rentRepository) FindByLesseeId(ctx context.Context, lesseId string, query *request.Common) ([]*entity.Rent, error) {
	var rents []*entity.Rent

	err := r.db.WithContext(ctx).Model(&entity.Rent{}).Preload("Room.Kost").Preload("Lessee").Unscoped().
		Joins("JOIN lessees ON rents.lessee_id = lessees.id").
		Joins("JOIN rooms ON rents.room_id = rooms.id").
		Joins("JOIN kosts ON rooms.kost_id = kosts.id").
		Where("rents.lessee_id = ?", lesseId).
		Scopes(
			db.Search("kost.name", query.Search),
		).
		Order("rents.start_date DESC").
		Offset(query.GetOffset()).Limit(query.GetLimit()).
		Find(&rents).Error

	if err != nil {
		return nil, err
	}

	return rents, nil
}

func (r *rentRepository) FindByKostId(ctx context.Context, kostId string, query *request.Common) ([]*entity.Rent, error) {
	var rents []*entity.Rent

	err := r.db.WithContext(ctx).Model(&entity.Rent{}).Preload("Room.Kost").Preload("Lessee").Unscoped().
		Joins("JOIN lessees ON rents.lessee_id = lessees.id").
		Joins("JOIN rooms ON rents.room_id = rooms.id").
		Joins("JOIN kosts ON rooms.kost_id = kosts.id").
		Where("rooms.kost_id = ?", kostId).
		Scopes(
			db.Search("lessees.fullname", query.Search),
		).
		Order("rents.start_date DESC").
		Offset(query.GetOffset()).Limit(query.GetLimit()).
		Find(&rents).Error

	if err != nil {
		return nil, err
	}

	return rents, nil
}

func (r *rentRepository) CountByLesseeId(ctx context.Context, lesseeId string, query *request.Common) (int, error) {
	var total int64

	err := r.db.WithContext(ctx).Model(&entity.Rent{}).Preload("Room.Kost").Preload("Lessee").Unscoped().
		Joins("JOIN lessees ON rents.lessee_id = lessees.id").
		Joins("JOIN rooms ON rents.room_id = rooms.id").
		Joins("JOIN kosts ON rooms.kost_id = kosts.id").
		Where("rents.lessee_id = ?", lesseeId).
		Scopes(
			db.Search("kost.name", query.Search),
		).
		Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), nil
}

func (r *rentRepository) CountByKostId(ctx context.Context, kostId string, query *request.Common) (int, error) {
	var total int64

	err := r.db.WithContext(ctx).Model(&entity.Rent{}).Preload("Room.Kost").Preload("Lessee").Unscoped().
		Joins("JOIN lessees ON rents.lessee_id = lessees.id").
		Joins("JOIN rooms ON rents.room_id = rooms.id").
		Joins("JOIN kosts ON rooms.kost_id = kosts.id").
		Where("rooms.kost_id = ?", kostId).
		Scopes(
			db.Search("lessees.fullname", query.Search),
		).
		Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), nil
}
