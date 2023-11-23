package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"gorm.io/gorm"
)

type SessionRepository interface {
	Save(ctx context.Context, session *entity.Session) error
	Update(ctx context.Context, session *entity.Session, sessionId string) error
	SaveWithTx(ctx context.Context, tx *gorm.DB, session *entity.Session) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, session *entity.Session, sessionId string) error
	FindByUserId(ctx context.Context, userId string) ([]*entity.Session, error)
	FindByDeviceId(ctx context.Context, deviceId string) (*entity.Session, error)
	FindById(ctx context.Context, sessionId string) (*entity.Session, error)
	DeleteByUserId(ctx context.Context, userId string) error
	DeleteById(ctx context.Context, sessionId string) error
}
