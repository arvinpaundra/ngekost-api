package session

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) contract.SessionRepository {
	return &sessionRepository{db: db}
}

func (s *sessionRepository) Save(ctx context.Context, session *entity.Session) error {
	err := s.db.WithContext(ctx).Model(&entity.Session{}).Create(&session).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionRepository) Update(ctx context.Context, session *entity.Session, sessionId string) error {
	err := s.db.WithContext(ctx).Model(&entity.Session{}).Where("id = ?", sessionId).Updates(&session).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionRepository) FindByUserId(ctx context.Context, userId string) ([]*entity.Session, error) {
	var sessions []*entity.Session
	err := s.db.WithContext(ctx).Model(&entity.Session{}).
		Where("user_id = ?", userId).
		Find(&sessions).Error

	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *sessionRepository) FindById(ctx context.Context, sessionId string) (*entity.Session, error) {
	var session entity.Session

	err := s.db.WithContext(ctx).Model(&entity.Session{}).
		Where("id = ?", sessionId).
		First(&session).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrSessionNotFound
		}
		return nil, err
	}

	return &session, nil
}

func (s *sessionRepository) FindByDeviceId(ctx context.Context, deviceId string) (*entity.Session, error) {
	var session entity.Session

	err := s.db.WithContext(ctx).Model(&entity.Session{}).
		Where("device_id = ?", deviceId).
		First(&session).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constant.ErrSessionNotFound
		}
		return nil, err
	}

	return &session, nil
}

func (s *sessionRepository) DeleteByUserId(ctx context.Context, userId string) error {
	err := s.db.WithContext(ctx).Model(&entity.Session{}).
		Where("user_id = ?", userId).
		Delete(&entity.Session{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (s *sessionRepository) DeleteById(ctx context.Context, sessionId string) error {
	err := s.db.WithContext(ctx).Model(&entity.Session{}).
		Where("id = ?", sessionId).
		Delete(&entity.Session{}).Error

	if err != nil {
		return err
	}

	return nil
}
