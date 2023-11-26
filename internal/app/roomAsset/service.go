package roomasset

import (
	"context"
	"mime/multipart"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/adapter/response"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/util/common"
	"github.com/arvinpaundra/ngekost-api/pkg/util/log"
)

type Service interface {
	Save(ctx context.Context, path *request.RoomAssetPathParam, file *multipart.FileHeader) error
	Update(ctx context.Context, path *request.RoomAssetPathParam, file *multipart.FileHeader) error
	Delete(ctx context.Context, path *request.RoomAssetPathParam) error
	FindByRoomId(ctx context.Context, path *request.RoomAssetPathParam) ([]*response.RoomAsset, error)
}

type service struct {
	txBeginner          contract.TxBeginner
	roomRepository      contract.RoomRepository
	roomAssetRepository contract.RoomAssetRepository
}

func NewService(f *factory.Factory) Service {
	return &service{
		txBeginner:          f.TxBeginner,
		roomRepository:      f.RoomRepository,
		roomAssetRepository: f.RoomAssetRepository,
	}
}

func (s *service) Save(ctx context.Context, path *request.RoomAssetPathParam, file *multipart.FileHeader) error {
	_, err := s.roomRepository.FindById(ctx, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	newAsset := entity.RoomAsset{
		ID:     common.GetID(),
		RoomId: path.RoomId,
		Url:    "",
		Type:   "",
	}

	err = s.roomAssetRepository.Save(ctx, &newAsset)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, path *request.RoomAssetPathParam, file *multipart.FileHeader) error {
	_, err := s.roomRepository.FindById(ctx, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	_, err = s.roomAssetRepository.FindById(ctx, path.RoomAssetId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	updatedAsset := entity.RoomAsset{
		Url:  "",
		Type: "",
	}

	err = s.roomAssetRepository.Save(ctx, &updatedAsset)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, path *request.RoomAssetPathParam) error {
	_, err := s.roomRepository.FindById(ctx, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	_, err = s.roomAssetRepository.FindById(ctx, path.RoomAssetId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	err = s.roomAssetRepository.Delete(ctx, path.RoomAssetId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) FindByRoomId(ctx context.Context, path *request.RoomAssetPathParam) ([]*response.RoomAsset, error) {
	_, err := s.roomRepository.FindById(ctx, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	assets, err := s.roomAssetRepository.FindByRoomId(ctx, path.RoomId)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	return response.ToResponseRoomAssets(assets), nil
}
