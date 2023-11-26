package kostrule

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
	Save(ctx context.Context, path *request.KostRulePathParam, req *request.CreateKostRule) error
	Update(ctx context.Context, path *request.KostRulePathParam, req *request.UpdateKostRule) error
	Delete(ctx context.Context, path *request.KostRulePathParam) error
	FindById(ctx context.Context, path *request.KostRulePathParam) (*response.KostRule, error)
}

type service struct {
	txBeginner         contract.TxBeginner
	kostRepository     contract.KostRepository
	kostRuleRepository contract.KostRuleRepository
}

func NewService(f *factory.Factory) Service {
	return &service{
		txBeginner:         f.TxBeginner,
		kostRepository:     f.KostRepository,
		kostRuleRepository: f.KostRuleRepository,
	}
}

func (s *service) Save(ctx context.Context, path *request.KostRulePathParam, req *request.CreateKostRule) error {
	_, err := s.kostRepository.FindById(ctx, path.KostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	newRule := entity.KostRule{
		ID:          common.GetID(),
		KostId:      path.KostId,
		Title:       req.Title,
		Priority:    req.Priority,
		Description: req.Description,
	}

	err = s.kostRuleRepository.Save(ctx, &newRule)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, path *request.KostRulePathParam, req *request.UpdateKostRule) error {
	_, err := s.kostRepository.FindById(ctx, path.KostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	_, err = s.kostRuleRepository.FindById(ctx, path.KostRuleId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	updatedRule := entity.KostRule{
		Title:       req.Title,
		Priority:    req.Priority,
		Description: req.Description,
	}

	err = s.kostRuleRepository.Update(ctx, &updatedRule, path.KostRuleId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, path *request.KostRulePathParam) error {
	_, err := s.kostRepository.FindById(ctx, path.KostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	_, err = s.kostRuleRepository.FindById(ctx, path.KostRuleId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	err = s.kostRuleRepository.Delete(ctx, path.KostRuleId)
	if err != nil {
		log.Logging().Error(err.Error())
		return err
	}

	return nil
}

func (s *service) FindById(ctx context.Context, path *request.KostRulePathParam) (*response.KostRule, error) {
	_, err := s.kostRepository.FindById(ctx, path.KostId)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	rule, err := s.kostRuleRepository.FindById(ctx, path.KostRuleId)
	if err != nil {
		log.Logging().Error(err.Error())
		return nil, err
	}

	return response.ToResponseKostRule(rule), nil
}
