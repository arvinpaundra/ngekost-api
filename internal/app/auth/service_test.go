package auth_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/app/auth"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract/mocks"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/util/common"
	"github.com/arvinpaundra/ngekost-api/pkg/util/token"
	jwtMock "github.com/arvinpaundra/ngekost-api/pkg/util/token/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ServiceTestTable struct {
	name string
	fn   func(*testing.T)
}

var (
	jsonWebToken      jwtMock.JSONWebToken
	transactioner     mocks.Transactioner
	transaction       mocks.Transaction
	cacheRepository   mocks.CacheRepository
	userRepository    mocks.UserRepository
	ownerRepository   mocks.OwnerRepository
	lesseRepository   mocks.LesseeRepository
	sessionRepository mocks.SessionRepository

	service auth.Service

	registerRoleOwner   request.Register
	registerRoleLessee  request.Register
	registerInvalidRole request.Register
	login               request.Login
	logout              request.Logout

	customClaims token.JWTCustomClaim
	user         entity.User
	owner        entity.Owner
	lessee       entity.Lessee
	session      entity.Session

	ctx context.Context
)

func initDataService() {
	f := factory.Factory{
		JSONWebToken:      &jsonWebToken,
		CacheRepository:   &cacheRepository,
		Transactioner:     &transactioner,
		UserRepository:    &userRepository,
		OwnerRepository:   &ownerRepository,
		LesseeRepository:  &lesseRepository,
		SessionRepository: &sessionRepository,
	}

	service = auth.NewService(&f)

	user = entity.User{
		ID:       common.GetID(),
		Username: "test",
		Password: "$2a$10$WhnKiwlPQjpYo2ScyjAFJ.bn5FtCwhvi9jN59Wk9N1sybc6BQJd1O",
		Role:     "owner",
	}

	owner = entity.Owner{
		ID:        common.GetID(),
		UserId:    user.ID,
		Fullname:  "test",
		Gender:    "test",
		Phone:     "test",
		Address:   "test",
		City:      "test",
		Birthdate: nil,
		Status:    nil,
		Photo:     nil,
	}

	lessee = entity.Lessee{
		ID:        common.GetID(),
		UserId:    user.ID,
		Fullname:  "test",
		Gender:    "test",
		Phone:     "test",
		City:      "test",
		Address:   "test",
		Birthdate: nil,
		Status:    nil,
		Photo:     nil,
	}

	session = entity.Session{
		ID:               common.GetID(),
		UserId:           user.ID,
		DeviceName:       "test",
		DeviceId:         "test",
		IPAddress:        "test",
		Platform:         "test",
		AccessToken:      "test",
		RefreshToken:     nil,
		FCMToken:         nil,
		GoogleOAuthToken: nil,
	}

	registerRoleOwner = request.Register{
		Username: "test",
		Password: "test",
		Role:     "owner",
		Fullname: "test",
		Phone:    "test",
		Gender:   "test",
		City:     "test",
		Address:  "test",
	}

	registerRoleLessee = request.Register{
		Username: "test",
		Password: "test",
		Role:     "lessee",
		Fullname: "test",
		Phone:    "test",
		Gender:   "test",
		City:     "test",
		Address:  "test",
	}

	registerInvalidRole = request.Register{
		Username: "test",
		Password: "test",
		Role:     "invalid",
		Fullname: "test",
		Phone:    "test",
		Gender:   "test",
		City:     "test",
		Address:  "test",
	}

	login = request.Login{
		Username:   "test",
		Password:   "test",
		DeviceId:   "test",
		Platform:   "test",
		DeviceName: "test",
		IPAddress:  "test",
		FCMToken:   nil,
	}

	logout = request.Logout{
		UserId:   user.ID,
		DeviceId: "test",
	}

	customClaims = token.JWTCustomClaim{
		UserId:    user.ID,
		Role:      "test",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now(),
	}

	ctx = context.Background()
}

func TestMain(m *testing.M) {

	initDataService()

	m.Run()
}

func TestRegister(t *testing.T) {
	tests := []ServiceTestTable{
		{
			name: "success with role owner",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("OwnerRepository").Return(&ownerRepository).Once()

				ownerRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("Commit").Return(nil).Once()

				err := service.Register(ctx, &registerRoleOwner)

				assert.NoError(t, err)
			},
		},
		{
			name: "success with role lessee",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("LesseeRepository").Return(&lesseRepository).Once()

				lesseRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("Commit").Return(nil).Once()

				err := service.Register(ctx, &registerRoleLessee)

				assert.NoError(t, err)
			},
		},
		{
			name: "error begin transaction",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(nil, errors.New("failed")).Once()

				err := service.Register(ctx, &registerRoleLessee)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
		{
			name: "error find user by username",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, errors.New("failed")).Once()

				err := service.Register(ctx, &registerRoleLessee)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
		{
			name: "username already used",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(&user, nil).Once()

				err := service.Register(ctx, &registerRoleLessee)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrUsernameAlreadyUsed.Error())
			},
		},
		{
			name: "error save user to db with rollback",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(errors.New("failed")).Once()

				transaction.On("Rollback").Return(errors.New("transaction failed")).Once()

				err := service.Register(ctx, &registerRoleLessee)

				assert.Error(t, err)
				assert.EqualError(t, err, "transaction failed")
			},
		},
		{
			name: "error save user to db",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(errors.New("failed")).Once()

				transaction.On("Rollback").Return(nil).Once()

				err := service.Register(ctx, &registerRoleLessee)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
		{
			name: "error save user role owner to db with rollback",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("OwnerRepository").Return(&ownerRepository).Once()

				ownerRepository.On("Save", ctx, mock.Anything).Return(errors.New("failed")).Once()

				transaction.On("Rollback").Return(errors.New("transaction failed")).Once()

				err := service.Register(ctx, &registerRoleOwner)

				assert.Error(t, err)
				assert.EqualError(t, err, "transaction failed")
			},
		},
		{
			name: "error save user role owner to db",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("OwnerRepository").Return(&ownerRepository).Once()

				ownerRepository.On("Save", ctx, mock.Anything).Return(errors.New("failed")).Once()

				transaction.On("Rollback").Return(nil).Once()

				err := service.Register(ctx, &registerRoleOwner)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
		{
			name: "error save user role lessee to db with rollback",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("LesseeRepository").Return(&lesseRepository).Once()

				lesseRepository.On("Save", ctx, mock.Anything).Return(errors.New("failed")).Once()

				transaction.On("Rollback").Return(errors.New("transaction failed")).Once()

				err := service.Register(ctx, &registerRoleLessee)

				assert.Error(t, err)
				assert.EqualError(t, err, "transaction failed")
			},
		},
		{
			name: "error save user role lessee to db",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("LesseeRepository").Return(&lesseRepository).Once()

				lesseRepository.On("Save", ctx, mock.Anything).Return(errors.New("failed")).Once()

				transaction.On("Rollback").Return(nil).Once()

				err := service.Register(ctx, &registerRoleLessee)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
		{
			name: "error invalid role with rollback",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("Rollback").Return(errors.New("transaction failed")).Once()

				err := service.Register(ctx, &registerInvalidRole)

				assert.Error(t, err)
				assert.EqualError(t, err, "transaction failed")
			},
		},
		{
			name: "error invalid role",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("Rollback").Return(nil).Once()

				err := service.Register(ctx, &registerInvalidRole)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrInvalidRole.Error())
			},
		},
		{
			name: "error commit transaction",
			fn: func(t *testing.T) {
				transactioner.On("Begin", ctx).Return(&transaction, nil).Once()

				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("OwnerRepository").Return(&ownerRepository).Once()

				ownerRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				transaction.On("Commit").Return(errors.New("transaction failed")).Once()

				err := service.Register(ctx, &registerRoleOwner)

				assert.Error(t, err)
				assert.EqualError(t, err, "transaction failed")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func TestLogin(t *testing.T) {
	tests := []ServiceTestTable{
		{
			name: "success",
			fn: func(t *testing.T) {
				userRepository.On("FindByUsername", ctx, user.Username).Return(&user, nil).Once()

				jsonWebToken.On("Encode", mock.Anything).Return("token", nil).Once()

				sessionRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				cacheRepository.On("Save", ctx, mock.Anything, mock.Anything, time.Minute*30).Return(nil).Once()

				res, err := service.Login(ctx, &login)

				assert.NoError(t, err)
				assert.NotEmpty(t, res)
			},
		},
		{
			name: "error user not found",
			fn: func(t *testing.T) {
				userRepository.On("FindByUsername", ctx, user.Username).Return(nil, constant.ErrUserNotFound).Once()

				res, err := service.Login(ctx, &login)

				assert.Error(t, err)
				assert.Empty(t, res)
				assert.EqualError(t, err, constant.ErrUserNotFound.Error())
			},
		},
		{
			name: "error password incorrect",
			fn: func(t *testing.T) {
				userRepository.On("FindByUsername", ctx, user.Username).Return(&user, nil).Once()

				res, err := service.Login(ctx, &request.Login{Username: "test", Password: "invalid"})

				assert.Error(t, err)
				assert.Empty(t, res)
				assert.EqualError(t, err, constant.ErrPasswordIncorrect.Error())
			},
		},
		{
			name: "error encode token",
			fn: func(t *testing.T) {
				userRepository.On("FindByUsername", ctx, user.Username).Return(&user, nil).Once()

				jsonWebToken.On("Encode", mock.Anything).Return("", errors.New("failed")).Once()

				res, err := service.Login(ctx, &login)

				assert.Error(t, err)
				assert.Empty(t, res)
				assert.EqualError(t, err, "failed")
			},
		},
		{
			name: "error encode token",
			fn: func(t *testing.T) {
				userRepository.On("FindByUsername", ctx, user.Username).Return(&user, nil).Once()

				jsonWebToken.On("Encode", mock.Anything).Return("token", nil).Once()

				sessionRepository.On("Save", ctx, mock.Anything).Return(errors.New("failed")).Once()

				res, err := service.Login(ctx, &login)

				assert.Error(t, err)
				assert.Empty(t, res)
				assert.EqualError(t, err, "failed")
			},
		},
		{
			name: "error save token to redis",
			fn: func(t *testing.T) {
				userRepository.On("FindByUsername", ctx, user.Username).Return(&user, nil).Once()

				jsonWebToken.On("Encode", mock.Anything).Return("token", nil).Once()

				sessionRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				cacheRepository.On("Save", ctx, mock.Anything, mock.Anything, time.Minute*30).Return(errors.New("failed")).Once()

				res, err := service.Login(ctx, &login)

				assert.Error(t, err)
				assert.Empty(t, res)
				assert.EqualError(t, err, "failed")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func TestLogout(t *testing.T) {
	tests := []ServiceTestTable{
		{
			name: "success",
			fn: func(t *testing.T) {
				userRepository.On("FindById", ctx, logout.UserId).Return(&user, nil).Once()

				sessionRepository.On("FindByDeviceId", ctx, logout.DeviceId).Return(&session, nil).Once()

				key := "session::" + session.ID
				cacheRepository.On("Del", ctx, key).Return(nil).Once()

				sessionRepository.On("DeleteById", ctx, session.ID).Return(nil).Once()

				err := service.Logout(ctx, &logout)

				assert.NoError(t, err)
			},
		},
		{
			name: "user not found",
			fn: func(t *testing.T) {
				userRepository.On("FindById", ctx, logout.UserId).Return(nil, constant.ErrUserNotFound).Once()

				err := service.Logout(ctx, &logout)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrUserNotFound.Error())
			},
		},
		{
			name: "session not found",
			fn: func(t *testing.T) {
				userRepository.On("FindById", ctx, logout.UserId).Return(&user, nil).Once()

				sessionRepository.On("FindByDeviceId", ctx, logout.DeviceId).Return(nil, constant.ErrSessionNotFound).Once()

				err := service.Logout(ctx, &logout)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrSessionNotFound.Error())
			},
		},
		{
			name: "error delete cache session from redis",
			fn: func(t *testing.T) {
				userRepository.On("FindById", ctx, logout.UserId).Return(&user, nil).Once()

				sessionRepository.On("FindByDeviceId", ctx, logout.DeviceId).Return(&session, nil).Once()

				key := "session::" + session.ID
				cacheRepository.On("Del", ctx, key).Return(errors.New("failed")).Once()

				err := service.Logout(ctx, &logout)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
		{
			name: "error delete session from db",
			fn: func(t *testing.T) {
				userRepository.On("FindById", ctx, logout.UserId).Return(&user, nil).Once()

				sessionRepository.On("FindByDeviceId", ctx, logout.DeviceId).Return(&session, nil).Once()

				key := "session::" + session.ID
				cacheRepository.On("Del", ctx, key).Return(nil).Once()

				sessionRepository.On("DeleteById", ctx, session.ID).Return(errors.New("failed")).Once()

				err := service.Logout(ctx, &logout)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
