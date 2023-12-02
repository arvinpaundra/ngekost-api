package rent_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	r "github.com/arvinpaundra/ngekost-api/internal/app/rent"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract/mocks"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/util/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	transactioner    mocks.Transactioner
	transaction      mocks.Transaction
	lesseeRepository mocks.LesseeRepository
	kostRepository   mocks.KostRepository
	roomRepository   mocks.RoomRepository
	rentRepository   mocks.RentRepository

	service r.Service

	createRent request.CreateRent

	lessee entity.Lessee
	kost   entity.Kost
	room   entity.Room
	rent   entity.Rent

	ctx context.Context
)

func initDataService() {
	f := factory.Factory{
		Transactioner:    &transactioner,
		LesseeRepository: &lesseeRepository,
		KostRepository:   &kostRepository,
		RoomRepository:   &roomRepository,
		RentRepository:   &rentRepository,
	}

	service = r.NewService(&f)

	lessee = entity.Lessee{
		ID:        common.GetID(),
		UserId:    common.GetID(),
		Fullname:  "test",
		Gender:    "test",
		Phone:     "test",
		City:      "test",
		Address:   "test",
		Birthdate: nil,
		Status:    nil,
		Photo:     nil,
	}

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

	rent = entity.Rent{
		ID:        common.GetID(),
		RoomId:    room.ID,
		LesseeId:  lessee.ID,
		StartDate: time.Now(),
		EndDate:   nil,
	}

	createRent = request.CreateRent{
		LesseeId: lessee.ID,
		RoomId:   room.ID,
	}

	ctx = context.Background()
}

type ServiceTestTable struct {
	name string
	fn   func(*testing.T)
}

func TestMain(m *testing.M) {

	initDataService()

	m.Run()
}

func TestSave(t *testing.T) {
	tests := []ServiceTestTable{
		{
			name: "success",
			fn: func(t *testing.T) {
				lesseeRepository.On("FindById", ctx, lessee.ID).Return(&lessee, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(&room, nil).Once()

				rentRepository.On("CheckExistRent", ctx, lessee.ID).Return(false, nil).Once()

				rentRepository.On("Save", ctx, mock.Anything).Return(nil).Once()

				err := service.Save(ctx, &createRent)

				assert.NoError(t, err)
			},
		},
		{
			name: "lesse not found",
			fn: func(t *testing.T) {
				lesseeRepository.On("FindById", ctx, lessee.ID).Return(nil, constant.ErrLesseeNotFound).Once()

				err := service.Save(ctx, &createRent)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrLesseeNotFound.Error())
			},
		},
		{
			name: "room not found",
			fn: func(t *testing.T) {
				lesseeRepository.On("FindById", ctx, lessee.ID).Return(&lessee, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(nil, constant.ErrRoomNotFound).Once()

				err := service.Save(ctx, &createRent)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrRoomNotFound.Error())
			},
		},
		{
			name: "error check exist rent",
			fn: func(t *testing.T) {
				lesseeRepository.On("FindById", ctx, lessee.ID).Return(&lessee, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(&room, nil).Once()

				rentRepository.On("CheckExistRent", ctx, lessee.ID).Return(false, errors.New("failed")).Once()

				err := service.Save(ctx, &createRent)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
		{
			name: "error contain exist rent active",
			fn: func(t *testing.T) {
				lesseeRepository.On("FindById", ctx, lessee.ID).Return(&lessee, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(&room, nil).Once()

				rentRepository.On("CheckExistRent", ctx, lessee.ID).Return(true, nil).Once()

				err := service.Save(ctx, &createRent)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrActiveRent.Error())
			},
		},
		{
			name: "error save rent to db",
			fn: func(t *testing.T) {
				lesseeRepository.On("FindById", ctx, lessee.ID).Return(&lessee, nil).Once()

				roomRepository.On("FindById", ctx, room.ID).Return(&room, nil).Once()

				rentRepository.On("CheckExistRent", ctx, lessee.ID).Return(false, nil).Once()

				rentRepository.On("Save", ctx, mock.Anything).Return(errors.New("failed")).Once()

				err := service.Save(ctx, &createRent)

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
	tests := []ServiceTestTable{
		{
			name: "success",
			fn: func(t *testing.T) {
				rentRepository.On("FindById", ctx, rent.ID).Return(&rent, nil).Once()

				rentRepository.On("Update", ctx, rent.ID, mock.Anything).Return(nil).Once()

				err := service.Update(ctx, rent.ID)

				assert.NoError(t, err)
			},
		},
		{
			name: "rent not found",
			fn: func(t *testing.T) {
				rentRepository.On("FindById", ctx, rent.ID).Return(nil, constant.ErrRentNotFound).Once()

				err := service.Update(ctx, rent.ID)

				assert.Error(t, err)
				assert.EqualError(t, err, constant.ErrRentNotFound.Error())
			},
		},
		{
			name: "error update rent to db",
			fn: func(t *testing.T) {
				rentRepository.On("FindById", ctx, rent.ID).Return(&rent, nil).Once()

				rentRepository.On("Update", ctx, rent.ID, mock.Anything).Return(errors.New("failed")).Once()

				err := service.Update(ctx, rent.ID)

				assert.Error(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
