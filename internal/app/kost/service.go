package kost

import (
	"context"
	"mime/multipart"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/adapter/response"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/helper/format"
	"github.com/arvinpaundra/ngekost-api/pkg/util/common"
	"github.com/arvinpaundra/ngekost-api/pkg/util/log"
)

type Service interface {
	Save(ctx context.Context, req *request.CreateKost, image *multipart.FileHeader) error
	Update(ctx context.Context, kostId string, req *request.UpdateKost, image *multipart.FileHeader) error
	Delete(ctx context.Context, kostId string) error
	FindAll(ctx context.Context, query *request.Common) (*response.WithPagination, error)
	FindByOwnerId(ctx context.Context, ownerId string, query *request.Common) (*response.WithPagination, error)
	FindById(ctx context.Context, kostId string) (*response.Kost, error)
	FindRoomsByKost(ctx context.Context, path *request.RoomPathParam, query *request.Common) (*response.WithPagination, error)
	FindRulesByKost(ctx context.Context, path *request.KostRulePathParam, query *request.Common) ([]*response.KostRule, error)
}

type service struct {
	txBeginner         contract.TxBeginner
	kostRepository     contract.KostRepository
	roomRepository     contract.RoomRepository
	kostRuleRepository contract.KostRuleRepository
	ownerRepository    contract.OwnerRepository
}

func NewService(f *factory.Factory) Service {
	return &service{
		txBeginner:         f.TxBeginner,
		kostRepository:     f.KostRepository,
		roomRepository:     f.RoomRepository,
		kostRuleRepository: f.KostRuleRepository,
		ownerRepository:    f.OwnerRepository,
	}
}

func (s *service) Save(ctx context.Context, req *request.CreateKost, image *multipart.FileHeader) error {
	_, err := s.ownerRepository.FindById(ctx, req.OwnerId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	newKost := entity.Kost{
		ID:              common.GetID(),
		OwnerId:         req.OwnerId,
		Name:            req.Name,
		Description:     req.Description,
		Type:            req.Type,
		PaymentInterval: req.PaymentInterval,
		Province:        req.Province,
		City:            req.City,
		District:        req.District,
		Subdistrict:     req.Subdistrict,
		Latitude:        req.Latitude,
		Longitude:       req.Longitude,
		Image:           nil,
	}

	err = s.kostRepository.Save(ctx, &newKost)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, kostId string, req *request.UpdateKost, image *multipart.FileHeader) error {
	_, err := s.kostRepository.FindById(ctx, kostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	updatedKost := entity.Kost{
		Name:            req.Name,
		Description:     req.Description,
		Type:            req.Type,
		PaymentInterval: req.PaymentInterval,
		Province:        req.Province,
		City:            req.City,
		District:        req.District,
		Subdistrict:     req.Subdistrict,
		Latitude:        req.Latitude,
		Longitude:       req.Longitude,
		Image:           nil,
	}

	err = s.kostRepository.Update(ctx, &updatedKost, kostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) FindAll(ctx context.Context, query *request.Common) (*response.WithPagination, error) {
	kosts, err := s.kostRepository.Find(ctx, query)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	total, err := s.kostRepository.Count(ctx, query)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	res := response.WithPagination{
		Results:    response.ToResponseKosts(kosts),
		Pagination: format.NewPagination(query.Offset, query.Limit, total),
	}

	return &res, nil
}

func (s *service) FindByOwnerId(ctx context.Context, ownerId string, query *request.Common) (*response.WithPagination, error) {
	kosts, err := s.kostRepository.FindByOwnerId(ctx, ownerId, query)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	total, err := s.kostRepository.CountByOwnerId(ctx, ownerId, query)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	pagination := format.NewPagination(query.Offset, query.Limit, total)

	res := response.WithPagination{
		Results:    response.ToResponseKosts(kosts),
		Pagination: pagination,
	}

	return &res, nil
}

func (s *service) FindById(ctx context.Context, kostId string) (*response.Kost, error) {
	kost, err := s.kostRepository.FindById(ctx, kostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	return response.ToResponseKost(kost), nil
}

func (s *service) Delete(ctx context.Context, kostId string) error {
	_, err := s.kostRepository.FindById(ctx, kostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	err = s.kostRepository.Delete(ctx, kostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) FindRoomsByKost(ctx context.Context, path *request.RoomPathParam, query *request.Common) (*response.WithPagination, error) {
	_, err := s.kostRepository.FindById(ctx, path.KostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	rooms, err := s.roomRepository.FindByKostId(ctx, path.KostId, query)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	total, err := s.roomRepository.CountByKostId(ctx, path.KostId, query)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	res := response.WithPagination{
		Results:    response.ToResponseRooms(rooms),
		Pagination: format.NewPagination(query.Offset, query.Limit, total),
	}

	return &res, nil
}

func (s *service) FindRulesByKost(ctx context.Context, path *request.KostRulePathParam, query *request.Common) ([]*response.KostRule, error) {
	_, err := s.kostRepository.FindById(ctx, path.KostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	rules, err := s.kostRuleRepository.FindByKostId(ctx, path.KostId, query)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	return response.ToResponseKostRules(rules), nil
}
