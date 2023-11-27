package bill

import (
	"context"
	"strings"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/util/db"
	"gorm.io/gorm"
)

type billRepository struct {
	db *gorm.DB
}

func NewBillRepository(db *gorm.DB) contract.BillRepository {
	return &billRepository{db: db}
}

func (b *billRepository) Save(ctx context.Context, bill *entity.Bill) error {
	err := b.db.WithContext(ctx).Model(&entity.Bill{}).Create(&bill).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *billRepository) SaveWithTx(ctx context.Context, tx *gorm.DB, bill *entity.Bill) error {
	err := tx.WithContext(ctx).Model(&entity.Bill{}).Create(&bill).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *billRepository) Update(ctx context.Context, billId string, bill *entity.Bill) error {
	err := b.db.WithContext(ctx).Model(&entity.Bill{}).Where("id = ?", billId).Updates(&bill).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *billRepository) UpdateWithTx(ctx context.Context, tx *gorm.DB, billId string, bill *entity.Bill) error {
	err := tx.WithContext(ctx).Model(&entity.Bill{}).Where("id = ?", billId).Updates(&bill).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *billRepository) FindAll(ctx context.Context, query *request.QueryParamBill) ([]*entity.Bill, error) {
	var bills []*entity.Bill

	err := b.db.WithContext(ctx).Model(&entity.Bill{}).Preload("Payment").Preload("Lessee.User").
		Joins("LEFT JOIN payments ON bills.id = payments.bill_id").
		Joins("JOIN lessees ON bills.lessee_id = lessees.id").
		Joins("JOIN users ON lessees.user_id = users.id").
		Scopes(
			func(db *gorm.DB) *gorm.DB {
				if query.Search != "" {
					return db.Where(
						db.Where("LOWER(bills.invoice) LIKE ?", "%"+strings.ToLower(query.Search)+"%").Or("LOWER(users.username) LIKE ?", "%"+strings.ToLower(query.Search)+"%"),
					)
				}
				return db
			},
			db.Exact("bills.status", query.Status),
		).
		Where("DATE(bills.created_at) >= ?", query.StartDate).
		Where("DATE(bills.created_at) <= ?", query.EndDate).
		Order("bills.created_at").
		Offset(query.GetOffset()).Limit(query.GetLimit()).
		Find(&bills).Error

	if err != nil {
		return nil, err
	}

	return bills, nil
}

func (b *billRepository) FindByLesseeId(ctx context.Context, lesseeId string, query *request.QueryParamBill) ([]*entity.Bill, error) {
	var bills []*entity.Bill

	err := b.db.WithContext(ctx).Model(&entity.Bill{}).Preload("Payment").Preload("Lessee.User").
		Joins("LEFT JOIN payments ON bills.id = payments.bill_id").
		Joins("JOIN lessees ON bills.lessee_id = lessees.id").
		Joins("JOIN users ON lessees.user_id = users.id").
		Where("bills.lessee_id = ?", lesseeId).
		Scopes(
			db.Search("bills.invoice", query.Search),
			db.Exact("bills.status", query.Status),
		).
		Where("DATE(bills.created_at) >= ?", query.StartDate).
		Where("DATE(bills.created_at) <= ?", query.EndDate).
		Order("bills.created_at").
		Offset(query.GetOffset()).Limit(query.GetLimit()).
		Find(&bills).Error

	if err != nil {
		return nil, err
	}

	return bills, nil
}

func (b *billRepository) FindById(ctx context.Context, billId string) (*entity.Bill, error) {
	var bill entity.Bill

	err := b.db.WithContext(ctx).Preload("Payment").Preload("Lessee.User").
		Joins("LEFT JOIN payments ON bills.id = payments.bill_id").
		Joins("JOIN lessees ON bills.lessee_id = lessees.id").
		Joins("JOIN users ON lessees.user_id = users.id").
		Where("bills.id = ?", billId).
		First(&bill).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrBillNotFound
		}
		return nil, err
	}

	return &bill, nil
}

func (b *billRepository) Count(ctx context.Context, query *request.QueryParamBill) (int, error) {
	var total int64

	err := b.db.WithContext(ctx).Model(&entity.Bill{}).Preload("Payment").Preload("Lessee.User").
		Joins("LEFT JOIN payments ON bills.id = payments.bill_id").
		Joins("JOIN lessees ON bills.lessee_id = lessees.id").
		Joins("JOIN users ON lessees.user_id = users.id").
		Scopes(
			func(db *gorm.DB) *gorm.DB {
				if query.Search != "" {
					return db.Where(
						db.Where("LOWER(bills.invoice) LIKE ?", "%"+strings.ToLower(query.Search)+"%").Or("LOWER(users.username) LIKE ?", "%"+strings.ToLower(query.Search)+"%"),
					)
				}
				return db
			},
			db.Exact("bills.status", query.Status),
		).
		Where("DATE(bills.created_at) >= ?", query.StartDate).
		Where("DATE(bills.created_at) <= ?", query.EndDate).
		Order("bills.created_at").
		Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), nil
}

func (b *billRepository) CountByLesseId(ctx context.Context, lesseeId string, query *request.QueryParamBill) (int, error) {
	var total int64

	err := b.db.WithContext(ctx).Model(&entity.Bill{}).Preload("Payment").Preload("Lessee.User").
		Joins("LEFT JOIN payments ON bills.id = payments.bill_id").
		Joins("JOIN lessees ON bills.lessee_id = lessees.id").
		Joins("JOIN users ON lessees.user_id = users.id").
		Where("bills.lessee_id = ?", lesseeId).
		Scopes(
			db.Search("bills.invoice", query.Search),
			db.Exact("bills.status", query.Status),
		).
		Where("DATE(bills.created_at) >= ?", query.StartDate).
		Where("DATE(bills.created_at) <= ?", query.EndDate).
		Order("bills.created_at").
		Count(&total).Error

	if err != nil {
		return 0, nil
	}

	return int(total), nil
}
