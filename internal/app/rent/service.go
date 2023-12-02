package rent

import (
	"context"
	"time"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/util/common"
	"github.com/arvinpaundra/ngekost-api/pkg/util/log"
)

type Service interface {
	Save(ctx context.Context, req *request.CreateRent) error
	Update(ctx context.Context, rentId string) error
}

type service struct {
	transactioner    contract.Transactioner
	lesseeRepository contract.LesseeRepository
	kostRepository   contract.KostRepository
	roomRepository   contract.RoomRepository
	rentRepository   contract.RentRepository
}

func NewService(f *factory.Factory) Service {
	return &service{
		transactioner:    f.Transactioner,
		lesseeRepository: f.LesseeRepository,
		kostRepository:   f.KostRepository,
		roomRepository:   f.RoomRepository,
		rentRepository:   f.RentRepository,
	}
}

func (s *service) Save(ctx context.Context, req *request.CreateRent) error {
	_, err := s.lesseeRepository.FindById(ctx, req.LesseeId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	_, err = s.roomRepository.FindById(ctx, req.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	activeRent, err := s.rentRepository.CheckExistRent(ctx, req.LesseeId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	if activeRent {
		log.Logging().Error(constant.ErrActiveRent)
		return constant.ErrActiveRent
	}

	newRent := entity.Rent{
		ID:        common.GetID(),
		RoomId:    req.RoomId,
		LesseeId:  req.LesseeId,
		StartDate: time.Now(),
		EndDate:   nil,
	}

	err = s.rentRepository.Save(ctx, &newRent)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, rentId string) error {
	_, err := s.rentRepository.FindById(ctx, rentId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	endDate := time.Now()

	updatedRent := entity.Rent{
		EndDate: &endDate,
	}

	err = s.rentRepository.Update(ctx, rentId, &updatedRent)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}
