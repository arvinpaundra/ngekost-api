package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"gorm.io/gorm"
)

type KostRuleRepository interface {
	Save(ctx context.Context, rule *entity.KostRule) error
	SaveWithTx(ctx context.Context, tx *gorm.DB, rule *entity.KostRule) error
	Update(ctx context.Context, rule *entity.KostRule, ruleId string) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, rule *entity.KostRule, ruleId string) error
	Delete(ctx context.Context, ruleId string) error
	DeleteWithTx(ctx context.Context, tx *gorm.DB, ruleId string) error
	FindByKostId(ctx context.Context, kostId string, query *request.Common) ([]*entity.KostRule, error)
	FindById(ctx context.Context, ruleId string) (*entity.KostRule, error)
}
