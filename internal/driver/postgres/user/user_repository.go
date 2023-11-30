package user

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) contract.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Save(ctx context.Context, user *entity.User) error {
	err := u.db.WithContext(ctx).Model(&entity.User{}).Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User

	err := u.db.WithContext(ctx).Model(&entity.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) FindById(ctx context.Context, userId string) (*entity.User, error) {
	var user entity.User

	err := u.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Update(ctx context.Context, tx *gorm.DB, user *entity.User, userId string) error {
	err := tx.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userId).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}
