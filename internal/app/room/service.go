package room

import (
	"context"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/adapter/response"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/util/common"
	"github.com/arvinpaundra/ngekost-api/pkg/util/log"
)

type Service interface {
	Save(ctx context.Context, path *request.RoomPathParam, req *request.CreateRoom) error
	Update(ctx context.Context, path *request.RoomPathParam, req *request.UpdateRoom) error
	Delete(ctx context.Context, path *request.RoomPathParam) error
	FindById(ctx context.Context, path *request.RoomPathParam) (*response.RoomDetail, error)
}

type service struct {
	txBeginner          contract.TxBeginner
	kostRepository      contract.KostRepository
	roomRepository      contract.RoomRepository
	roomAssetRepository contract.RoomAssetRepository
}

func NewService(f *factory.Factory) Service {
	return &service{
		txBeginner:          f.TxBeginner,
		kostRepository:      f.KostRepository,
		roomRepository:      f.RoomRepository,
		roomAssetRepository: f.RoomAssetRepository,
	}
}

func (s *service) Save(ctx context.Context, path *request.RoomPathParam, req *request.CreateRoom) error {
	_, err := s.kostRepository.FindById(ctx, path.KostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	newRoom := entity.Room{
		ID:          common.GetID(),
		KostId:      path.KostId,
		Name:        req.Name,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Category:    req.Category,
		Description: req.Description,
		Image:       nil,
	}

	err = s.roomRepository.Save(ctx, &newRoom)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, path *request.RoomPathParam, req *request.UpdateRoom) error {
	_, err := s.kostRepository.FindById(ctx, path.KostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	_, err = s.roomRepository.FindById(ctx, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	updatedRoom := entity.Room{
		Name:        req.Name,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Category:    req.Category,
		Description: req.Description,
		Image:       nil,
	}

	err = s.roomRepository.Update(ctx, &updatedRoom, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, path *request.RoomPathParam) error {
	_, err := s.kostRepository.FindById(ctx, path.KostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	_, err = s.roomRepository.FindById(ctx, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	err = s.roomRepository.Delete(ctx, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) FindById(ctx context.Context, path *request.RoomPathParam) (*response.RoomDetail, error) {
	_, err := s.kostRepository.FindById(ctx, path.KostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	room, err := s.roomRepository.FindById(ctx, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	assets, err := s.roomAssetRepository.FindByRoomId(ctx, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	return response.ToResponseRoom(room, assets), nil
}
