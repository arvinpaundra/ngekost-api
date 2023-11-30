package contract

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindById(ctx context.Context, userId string) (*entity.User, error)
	Update(ctx context.Context, tx *gorm.DB, user *entity.User, userId string) error
}
