package kost_test

import (
	"context"
	"errors"
	"testing"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	k "github.com/arvinpaundra/ngekost-api/internal/app/kost"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract/mocks"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/util/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ServiceTableTest struct {
	name string
	fn   func(*testing.T)
}

var (
	txBeginner         mocks.TxBeginner
	kostRepository     mocks.KostRepository
	roomRepository     mocks.RoomRepository
	kostRuleRepository mocks.KostRuleRepository
	ownerRepository    mocks.OwnerRepository

	service k.Service

	createKost request.CreateKost
	updateKost request.UpdateKost
	query      request.Common
	roomPath   request.RoomPathParam
	rulePath   request.KostRulePathParam

	kost  entity.Kost
	rule  entity.KostRule
	room  entity.Room
	owner entity.Owner

	ctx context.Context
)

func initDataService() {
	f := factory.Factory{
		TxBeginner:         &txBeginner,
		OwnerRepository:    &ownerRepository,
		KostRepository:     &kostRepository,
		RoomRepository:     &roomRepository,
		KostRuleRepository: &kostRuleRepository,
	}

	service = k.NewService(&f)

	ctx = context.Background()

	owner = entity.Owner{
		ID:        common.GetID(),
		UserId:    common.GetID(),
		Fullname:  "test",
		Gender:    "test",
		Phone:     "test",
		Address:   "test",
		City:      "test",
		Birthdate: nil,
		Status:    nil,
		Photo:     nil,
	}

	kost = entity.Kost{
		ID:              common.GetID(),
		OwnerId:         owner.ID,
		Name:            "test",
		Description:     "test",
		Type:            "test",
		PaymentInterval: "test",
		Province:        "test",
		City:            "test",
		District:        "test",
		Subdistrict:     "test",
		Latitude:        0,
		Longitude:       0,
		Image:           nil,
	}

	room = entity.Room{
		ID:          common.GetID(),
		KostId:      kost.ID,
		Name:        "test",
		Quantity:    0,
		Price:       0,
		Category:    nil,
		Description: nil,
		Image:       nil,
	}

	rule = entity.KostRule{
		ID:          common.GetID(),
		KostId:      kost.ID,
		Title:       "test",
		Priority:    "test",
		Description: nil,
	}

	createKost = request.CreateKost{
		OwnerId:         owner.ID,
		Name:            kost.Name,
		Description:     kost.Description,
		Type:            kost.Type,
		PaymentInterval: kost.PaymentInterval,
		Province:        kost.Province,
		City:            kost.City,
		District:        kost.District,
		Subdistrict:     kost.Subdistrict,
		Latitude:        kost.Latitude,
		Longitude:       kost.Longitude,
	}

	updateKost = request.UpdateKost{
		Name:            kost.Name,
		Description:     kost.Description,
		Type:            kost.Type,
		PaymentInterval: kost.PaymentInterval,
		Province:        kost.Province,
		City:            kost.City,
		District:        kost.District,
		Subdistrict:     kost.Subdistrict,
		Latitude:        kost.Latitude,
		Longitude:       kost.Longitude,
	}

	query = request.Common{
		Limit:  10,
		Offset: 1,
		Search: "",
	}

	roomPath = request.RoomPathParam{
		KostId: kost.ID,
		RoomId: room.ID,
	}

	rulePath = request.KostRulePathParam{
		KostId:     kost.ID,
		KostRuleId: rule.ID,
	}
}

func TestMain(m *testing.M) {

	initDataService()

	m.Run()
}

func TestSave(t *testing.T) {
	tests := []ServiceTableTest{
		{
			name: "success",
			fn: func(t *testing.T) {
				ownerRepository.On("FindById", ctx, owner.ID).Return(&owner, nil).Once()

				kostRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				err := service.Save(ctx, &createKost, nil)

				assert.NoError(t, err)
			},
		},
		{
			name: "owner not found",
			fn: func(t *testing.T) {
				ownerRepository.On("FindById", ctx, owner.ID).Return(nil, constant.ErrOwnerNotFound).Once()

				err := service.Save(ctx, &createKost, nil)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrOwnerNotFound.Error())
			},
		},
		{
			name: "error save to db",
			fn: func(t *testing.T) {
				ownerRepository.On("FindById", ctx, owner.ID).Return(&owner, nil).Once()

				kostRepository.On("Save", ctx, mock.Anything).Return(errors.New("failed")).Once()

				err := service.Save(ctx, &createKost, nil)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func TestUpdate(t *testing.T) {
	tests := []ServiceTableTest{
		{
			name: "success",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				kostRepository.On("Update", ctx, mock.Anything, kost.ID).Return(nil).Once()

				err := service.Update(ctx, kost.ID, &updateKost, nil)

				assert.NoError(t, err)
			},
		},
		{
			name: "kost not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(nil, constant.ErrKostNotFound).Once()

				err := service.Update(ctx, kost.ID, &updateKost, nil)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrKostNotFound.Error())
			},
		},
		{
			name: "error update to db",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				kostRepository.On("Update", ctx, mock.Anything, kost.ID).Return(errors.New("failed")).Once()

				err := service.Update(ctx, kost.ID, &updateKost, nil)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func TestDelete(t *testing.T) {
	tests := []ServiceTableTest{
		{
			name: "success",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				kostRepository.On("Delete", ctx, kost.ID).Return(nil).Once()

				err := service.Delete(ctx, kost.ID)

				assert.NoError(t, err)
			},
		},
		{
			name: "kost not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(nil, constant.ErrKostNotFound).Once()

				err := service.Delete(ctx, kost.ID)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrKostNotFound.Error())
			},
		},
		{
			name: "error delete from db",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				kostRepository.On("Delete", ctx, kost.ID).Return(errors.New("failed")).Once()

				err := service.Delete(ctx, kost.ID)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func TestFindById(t *testing.T) {
	tests := []ServiceTableTest{
		{
			name: "success",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				res, err := service.FindById(ctx, kost.ID)

				assert.NoError(t, err)
				assert.NotEmpty(t, res)
			},
		},
		{
			name: "kost not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(nil, constant.ErrKostNotFound).Once()

				res, err := service.FindById(ctx, kost.ID)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrKostNotFound.Error())
				assert.Empty(t, res)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func TestFindAll(t *testing.T) {
	tests := []ServiceTableTest{
		{
			name: "success",
			fn: func(t *testing.T) {
				kostRepository.On("Find", ctx, &query).Return([]*entity.Kost{&kost}, nil).Once()

				kostRepository.On("Count", ctx, &query).Return(1, nil).Once()

				res, err := service.FindAll(ctx, &query)

				assert.NoError(t, err)
				assert.NotEmpty(t, res)
			},
		},
		{
			name: "error retrieve kosts from db",
			fn: func(t *testing.T) {
				kostRepository.On("Find", ctx, &query).Return(nil, errors.New("failed")).Once()

				res, err := service.FindAll(ctx, &query)

				assert.Error(t, err)
				assert.Empty(t, res)
			},
		},
		{
			name: "error retrieve amount of kosts from db",
			fn: func(t *testing.T) {
				kostRepository.On("Find", ctx, &query).Return([]*entity.Kost{&kost}, nil).Once()

				kostRepository.On("Count", ctx, &query).Return(0, errors.New("failed"))

				res, err := service.FindAll(ctx, &query)

				assert.Error(t, err)
				assert.Empty(t, res)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func TestFindByOwnerId(t *testing.T) {
	tests := []ServiceTableTest{
		{
			name: "success",
			fn: func(t *testing.T) {
				kostRepository.On("FindByOwnerId", ctx, owner.ID, &query).Return([]*entity.Kost{&kost}, nil).Once()

				kostRepository.On("CountByOwnerId", ctx, owner.ID, &query).Return(1, nil).Once()

				res, err := service.FindByOwnerId(ctx, owner.ID, &query)

				assert.NoError(t, err)
				assert.NotEmpty(t, res)
			},
		},
		{
			name: "error retrieve kosts from db",
			fn: func(t *testing.T) {
				kostRepository.On("FindByOwnerId", ctx, owner.ID, &query).Return(nil, errors.New("failed")).Once()

				res, err := service.FindByOwnerId(ctx, owner.ID, &query)

				assert.Error(t, err)
				assert.Empty(t, res)
			},
		},
		{
			name: "error retrieve amount of kosts from db",
			fn: func(t *testing.T) {
				kostRepository.On("FindByOwnerId", ctx, owner.ID, &query).Return([]*entity.Kost{&kost}, nil).Once()

				kostRepository.On("CountByOwnerId", ctx, owner.ID, &query).Return(0, errors.New("failed"))

				res, err := service.FindByOwnerId(ctx, owner.ID, &query)

				assert.Error(t, err)
				assert.Empty(t, res)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func TestFindRoomsByKostId(t *testing.T) {
	tests := []ServiceTableTest{
		{
			name: "success",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("FindByKostId", ctx, kost.ID, &query).Return([]*entity.Room{&room}, nil).Once()

				roomRepository.On("CountByKostId", ctx, kost.ID, &query).Return(1, nil).Once()

				res, err := service.FindRoomsByKost(ctx, &roomPath, &query)

				assert.NoError(t, err)
				assert.NotEmpty(t, res)
			},
		},
		{
			name: "kost not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(nil, constant.ErrKostNotFound).Once()

				res, err := service.FindRoomsByKost(ctx, &roomPath, &query)

				assert.Error(t, err)
				assert.Empty(t, res)
			},
		},
		{
			name: "error retrieve rooms from db",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("FindByKostId", ctx, kost.ID, &query).Return(nil, errors.New("failed")).Once()

				res, err := service.FindRoomsByKost(ctx, &roomPath, &query)

				assert.Error(t, err)
				assert.Empty(t, res)
			},
		},
		{
			name: "error retrieve amount of rooms from db",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("FindByKostId", ctx, kost.ID, &query).Return([]*entity.Room{&room}, nil).Once()

				roomRepository.On("CountByKostId", ctx, kost.ID, &query).Return(0, errors.New("failed")).Once()

				res, err := service.FindRoomsByKost(ctx, &roomPath, &query)

				assert.Error(t, err)
				assert.Empty(t, res)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func TestFindRulesByKost(t *testing.T) {
	tests := []ServiceTableTest{
		{
			name: "success",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				kostRuleRepository.On("FindByKostId", ctx, kost.ID, &query).Return([]*entity.KostRule{&rule}, nil).Once()

				res, err := service.FindRulesByKost(ctx, &rulePath, &query)

				assert.NoError(t, err)
				assert.NotEmpty(t, res)
			},
		},
		{
			name: "kost not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(nil, constant.ErrKostNotFound).Once()

				res, err := service.FindRulesByKost(ctx, &rulePath, &query)

				assert.Error(t, err)
				assert.Empty(t, res)
			},
		},
		{
			name: "error retrieve kost rules from db",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				kostRuleRepository.On("FindByKostId", ctx, kost.ID, &query).Return(nil, errors.New("failed")).Once()

				res, err := service.FindRulesByKost(ctx, &rulePath, &query)

				assert.Error(t, err)
				assert.Empty(t, res)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
