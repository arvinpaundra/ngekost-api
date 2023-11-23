package auth

import (
	"context"
	"time"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/adapter/response"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/helper/format"
	"github.com/arvinpaundra/ngekost-api/pkg/util/common"
	"github.com/arvinpaundra/ngekost-api/pkg/util/log"
	"github.com/arvinpaundra/ngekost-api/pkg/util/token"
)

type Service interface {
	Register(ctx context.Context, req *request.Register) error
	Login(ctx context.Context, req *request.Login) (*response.Login, error)
	Logout(ctx context.Context, req *request.Logout) error
}

type service struct {
	jsonWebToken      token.JSONWebToken
	cacheRepository   contract.CacheRepository
	txBeginner        contract.TxBeginner
	userRepository    contract.UserRepository
	ownerRepository   contract.OwnerRepository
	lesseRepository   contract.LesseeRepository
	sessionRepository contract.SessionRepository
}

func NewService(f *factory.Factory) Service {
	return &service{
		jsonWebToken:      f.JSONWebToken,
		cacheRepository:   f.CacheRepository,
		txBeginner:        f.TxBeginner,
		userRepository:    f.UserRepository,
		ownerRepository:   f.OwnerRepository,
		lesseRepository:   f.LesseeRepository,
		sessionRepository: f.SessionRepository,
	}
}

func (s *service) Register(ctx context.Context, req *request.Register) error {
	tx := s.txBeginner.Begin()

	user, err := s.userRepository.FindByUsername(ctx, req.Username)
	if err != nil && err != constant.ErrUserNotFound {
		log.Logging(err.Error()).Error()
		return err
	}

	if user != nil {
		log.Logging(constant.ErrUsernameAlreadyUsed.Error()).Error()
		return constant.ErrUsernameAlreadyUsed
	}

	newUser := entity.User{
		ID:       common.GetID(),
		Username: req.Username,
		Password: common.HashPassword(req.Password),
		Role:     req.Role,
	}

	err = s.userRepository.SaveWithTx(ctx, tx, &newUser)
	if err != nil {
		if errRollback := tx.Rollback().Error; errRollback != nil {
			log.Logging(errRollback.Error()).Error()
			return errRollback
		}
		log.Logging(err.Error()).Error()
		return err
	}

	switch req.Role {
	case "owner":
		newOwner := entity.Owner{
			ID:       common.GetID(),
			UserId:   newUser.ID,
			Fullname: req.Fullname,
			Gender:   req.Gender,
			Phone:    req.Phone,
			Address:  req.Address,
			City:     req.City,
		}

		err := s.ownerRepository.SaveWithTx(ctx, tx, &newOwner)
		if err != nil {
			if errRollback := tx.Rollback().Error; errRollback != nil {
				log.Logging(errRollback.Error()).Error()
				return errRollback
			}
			log.Logging(err.Error()).Error()
			return err
		}
	case "lessee":
		newLessee := entity.Lessee{
			ID:       common.GetID(),
			UserId:   newUser.ID,
			Fullname: req.Fullname,
			Gender:   req.Gender,
			Phone:    req.Phone,
			City:     req.City,
			Address:  req.Address,
		}

		err := s.lesseRepository.SaveWithTx(ctx, tx, &newLessee)
		if err != nil {
			if errRollback := tx.Rollback().Error; errRollback != nil {
				log.Logging(errRollback.Error()).Error()
				return errRollback
			}
			log.Logging(err.Error()).Error()
			return err
		}
	default:
		if errRollback := tx.Rollback().Error; errRollback != nil {
			log.Logging(errRollback.Error()).Error()
			return errRollback
		}
		return constant.ErrInvalidRole
	}

	if errCommit := tx.Commit().Error; errCommit != nil {
		log.Logging(errCommit.Error()).Error()
		return errCommit
	}

	return nil
}

func (s *service) Login(ctx context.Context, req *request.Login) (*response.Login, error) {
	user, err := s.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		log.Logging(err.Error()).Error()
		return nil, err
	}

	if !common.ComparePassword(user.Password, req.Password) {
		return nil, constant.ErrPasswordIncorrect
	}

	// TODO: create token and refresh token
	now := time.Now()

	accessTokenClaims := token.JWTCustomClaim{
		UserId:    user.ID,
		Role:      user.Role,
		IssuedAt:  now,
		ExpiredAt: now.Add(24 * time.Hour),
	}

	accessToken, err := s.jsonWebToken.Encode(&accessTokenClaims)
	if err != nil {
		log.Logging(err.Error()).Error()
		return nil, err
	}

	// refreshTokenClaims := token.JWTCustomClaim{
	// 	UserId:    user.ID,
	// 	Role:      user.Role,
	// 	IssuedAt:  now,
	// 	ExpiredAt: now.AddDate(0, 6, 0),
	// }

	// refreshToken, err := s.jsonWebToken.Encode(&refreshTokenClaims)
	// if err != nil {
	// 	log.Logging(err.Error()).Error()
	// 	return nil, err
	// }

	session := entity.Session{
		ID:           common.GetID(),
		UserId:       user.ID,
		DeviceId:     req.DeviceId,
		DeviceName:   req.DeviceName,
		IPAddress:    req.IPAddress,
		Platform:     req.Platform,
		AccessToken:  accessToken,
		RefreshToken: nil,
		FCMToken:     req.FCMToken,
	}

	err = s.sessionRepository.Save(ctx, &session)
	if err != nil {
		log.Logging(err.Error()).Error()
		return nil, err
	}

	key := "session::" + session.ID
	err = s.cacheRepository.Save(ctx, key, session, time.Minute*30)
	if err != nil {
		log.Logging(err.Error()).Error()
		return nil, err
	}

	res := response.Login{
		UserId:       user.ID,
		Role:         user.Role,
		IssuedAt:     format.DatetimeToString(accessTokenClaims.IssuedAt),
		ExpiredAt:    format.DatetimeToString(accessTokenClaims.ExpiredAt),
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
	}

	return &res, nil
}

func (s *service) Logout(ctx context.Context, req *request.Logout) error {
	// check is user exist
	_, err := s.userRepository.FindById(ctx, req.UserId)
	if err != nil {
		log.Logging(err.Error()).Error()
		return err
	}

	// check if session exist for current user
	session, err := s.sessionRepository.FindByDeviceId(ctx, req.DeviceId)
	if err != nil {
		log.Logging(err.Error()).Error()
		return err
	}

	// delete session from cache
	key := "session::" + session.ID
	err = s.cacheRepository.Del(ctx, key)
	if err != nil {
		log.Logging(err.Error()).Error()
		return err
	}

	// delete session from database
	err = s.sessionRepository.DeleteById(ctx, session.ID)
	if err != nil {
		log.Logging(err.Error()).Error()
		return err
	}

	return nil
}
