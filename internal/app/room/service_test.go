package room_test

import (
	"context"
	"errors"
	"testing"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	r "github.com/arvinpaundra/ngekost-api/internal/app/room"
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
	txBeginner          mocks.TxBeginner
	kostRepository      mocks.KostRepository
	roomRepository      mocks.RoomRepository
	roomAssetRepository mocks.RoomAssetRepository

	service r.Service

	createRoom request.CreateRoom
	updateRoom request.UpdateRoom
	query      request.Common
	roomPath   request.RoomPathParam

	kost  entity.Kost
	room  entity.Room
	asset entity.RoomAsset

	ctx context.Context
)

func initDataService() {
	f := factory.Factory{
		TxBeginner:          &txBeginner,
		KostRepository:      &kostRepository,
		RoomRepository:      &roomRepository,
		RoomAssetRepository: &roomAssetRepository,
	}

	service = r.NewService(&f)

	kost = entity.Kost{
		ID:              common.GetID(),
		OwnerId:         common.GetID(),
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

	asset = entity.RoomAsset{
		ID:     common.GetID(),
		RoomId: room.ID,
		Url:    "https://test.com",
		Type:   "test",
	}

	createRoom = request.CreateRoom{
		Name:        "test",
		Quantity:    0,
		Price:       0,
		Description: nil,
		Category:    nil,
		Image:       nil,
	}

	updateRoom = request.UpdateRoom{
		Name:        "test",
		Quantity:    0,
		Price:       0,
		Description: nil,
		Category:    nil,
		Image:       nil,
	}

	query = request.Common{
		Limit:  10,
		Offset: 0,
		Search: "",
	}

	roomPath = request.RoomPathParam{
		KostId: kost.ID,
		RoomId: room.ID,
	}

	ctx = context.Background()
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
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				err := service.Save(ctx, &roomPath, &createRoom)

				assert.NoError(t, err)
			},
		},
		{
			name: "kost not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(nil, constant.ErrKostNotFound).Once()

				err := service.Save(ctx, &roomPath, &createRoom)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrKostNotFound.Error())
			},
		},
		{
			name: "error save to db",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("Save", ctx, mock.Anything).Return(errors.New("failed")).Once()

				err := service.Save(ctx, &roomPath, &createRoom)

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

				roomRepository.On("FindById", ctx, room.ID).Return(&room, nil).Once()

				roomRepository.On("Update", ctx, mock.Anything, room.ID).Return(nil).Once()

				err := service.Update(ctx, &roomPath, &updateRoom)

				assert.NoError(t, err)
			},
		},
		{
			name: "kost not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(nil, constant.ErrKostNotFound).Once()

				err := service.Update(ctx, &roomPath, &updateRoom)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrKostNotFound.Error())
			},
		},
		{
			name: "room not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(nil, constant.ErrRoomNotFound).Once()

				err := service.Update(ctx, &roomPath, &updateRoom)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrRoomNotFound.Error())
			},
		},
		{
			name: "error update to db",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(&room, nil).Once()

				roomRepository.On("Update", ctx, mock.Anything, room.ID).Return(errors.New("failed")).Once()

				err := service.Update(ctx, &roomPath, &updateRoom)

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

				roomRepository.On("FindById", ctx, room.ID).Return(&room, nil).Once()

				roomRepository.On("Delete", ctx, room.ID).Return(nil).Once()

				err := service.Delete(ctx, &roomPath)

				assert.NoError(t, err)
			},
		},
		{
			name: "kost not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(nil, constant.ErrKostNotFound).Once()

				err := service.Delete(ctx, &roomPath)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrKostNotFound.Error())
			},
		},
		{
			name: "room not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(nil, constant.ErrRoomNotFound).Once()

				err := service.Delete(ctx, &roomPath)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrRoomNotFound.Error())
			},
		},
		{
			name: "error delete from db",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(&room, nil).Once()

				roomRepository.On("Delete", ctx, room.ID).Return(errors.New("failed")).Once()

				err := service.Delete(ctx, &roomPath)

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

				roomRepository.On("FindById", ctx, room.ID).Return(&room, nil).Once()

				roomAssetRepository.On("FindByRoomId", ctx, room.ID).Return([]*entity.RoomAsset{&asset}, nil).Once()

				res, err := service.FindById(ctx, &roomPath)

				assert.NoError(t, err)
				assert.NotEmpty(t, res)
			},
		},
		{
			name: "kost not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(nil, constant.ErrKostNotFound).Once()

				res, err := service.FindById(ctx, &roomPath)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrKostNotFound.Error())
				assert.Empty(t, res)
			},
		},
		{
			name: "room not found",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(nil, constant.ErrRoomNotFound).Once()

				res, err := service.FindById(ctx, &roomPath)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrRoomNotFound.Error())
				assert.Empty(t, res)
			},
		},
		{
			name: "error retrieve assets from db",
			fn: func(t *testing.T) {
				kostRepository.On("FindById", ctx, kost.ID).Return(&kost, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(&room, nil).Once()

				roomAssetRepository.On("FindByRoomId", ctx, room.ID).Return(nil, errors.New("failed")).Once()

				res, err := service.FindById(ctx, &roomPath)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
				assert.Empty(t, res)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
