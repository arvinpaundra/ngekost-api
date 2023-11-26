package kostrule

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"gorm.io/gorm"
)

type kostRuleRepository struct {
	db *gorm.DB
}

func NewKostRuleRepository(db *gorm.DB) contract.KostRuleRepository {
	return &kostRuleRepository{db: db}
}

func (k *kostRuleRepository) Save(ctx context.Context, rule *entity.KostRule) error {
	err := k.db.WithContext(ctx).Model(&entity.KostRule{}).Create(&rule).Error
	if err != nil {
		return err
	}

	return nil
}

func (k *kostRuleRepository) SaveWithTx(ctx context.Context, tx *gorm.DB, rule *entity.KostRule) error {
	err := tx.WithContext(ctx).Model(&entity.KostRule{}).Create(&rule).Error
	if err != nil {
		return err
	}

	return nil
}

func (k *kostRuleRepository) Update(ctx context.Context, rule *entity.KostRule, ruleId string) error {
	err := k.db.WithContext(ctx).Model(&entity.KostRule{}).Where("id = ?", ruleId).Updates(&rule).Error
	if err != nil {
		return err
	}

	return nil
}

func (k *kostRuleRepository) UpdateWithTx(ctx context.Context, tx *gorm.DB, rule *entity.KostRule, ruleId string) error {
	err := tx.WithContext(ctx).Model(&entity.KostRule{}).Where("id = ?", ruleId).Updates(&rule).Error
	if err != nil {
		return err
	}

	return nil
}

func (k *kostRuleRepository) Delete(ctx context.Context, ruleId string) error {
	err := k.db.WithContext(ctx).Model(&entity.KostRule{}).Where("id = ?", ruleId).Delete(&entity.KostRule{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (k *kostRuleRepository) DeleteWithTx(ctx context.Context, tx *gorm.DB, ruleId string) error {
	err := tx.WithContext(ctx).Model(&entity.KostRule{}).Where("id = ?", ruleId).Delete(&entity.KostRule{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (k *kostRuleRepository) FindByKostId(ctx context.Context, kostId string, query *request.Common) ([]*entity.KostRule, error) {
	var rules []*entity.KostRule

	err := k.db.WithContext(ctx).Model(&entity.KostRule{}).
		Where("kost_id = ?", kostId).
		Order("created_at asc").
		Find(&rules).Error

	if err != nil {
		return nil, err
	}

	return rules, nil
}

func (k *kostRuleRepository) FindById(ctx context.Context, ruleId string) (*entity.KostRule, error) {
	var rule entity.KostRule

	err := k.db.WithContext(ctx).Model(&entity.KostRule{}).
		Where("id = ?", ruleId).
		First(&rule).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrKostRuleNotFound
		}
		return nil, err
	}

	return &rule, nil
}
