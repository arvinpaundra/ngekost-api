package payment

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/util/db"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) contract.PaymentRepository {
	return &paymentRepository{db: db}
}

func (p *paymentRepository) Save(ctx context.Context, payment *entity.Payment) error {
	err := p.db.WithContext(ctx).Model(&entity.Payment{}).Create(&payment).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *paymentRepository) SaveWithTx(ctx context.Context, tx *gorm.DB, payment *entity.Payment) error {
	err := tx.WithContext(ctx).Model(&entity.Payment{}).Create(&payment).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *paymentRepository) Update(ctx context.Context, paymentId string, payment *entity.Payment) error {
	err := p.db.WithContext(ctx).Model(&entity.Payment{}).Where("id = ?", paymentId).Updates(&payment).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *paymentRepository) UpdateWithTx(ctx context.Context, tx *gorm.DB, paymentId string, payment *entity.Payment) error {
	err := tx.WithContext(ctx).Model(&entity.Payment{}).Where("id = ?", paymentId).Updates(&payment).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *paymentRepository) FindAll(ctx context.Context, query *request.QueryParamPayment) ([]*entity.Payment, error) {
	var payments []*entity.Payment

	err := p.db.WithContext(ctx).Model(&entity.Payment{}).
		Scopes(
			db.Exact("status", query.Status),
		).
		Offset(query.GetOffset()).Limit(query.GetLimit()).
		Find(&payments).Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (p *paymentRepository) FindById(ctx context.Context, paymentId string) (*entity.Payment, error) {
	var payment entity.Payment

	err := p.db.WithContext(ctx).Model(&entity.Payment{}).
		Where("id = ?", paymentId).
		First(&payment).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrPaymentNotFound
		}
		return nil, err
	}

	return &payment, nil
}

func (p *paymentRepository) Count(ctx context.Context, query *request.QueryParamPayment) (int, error) {
	var total int64

	err := p.db.WithContext(ctx).Model(&entity.Payment{}).
		Scopes(
			db.Exact("status", query.Status),
		).
		Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), nil
}
