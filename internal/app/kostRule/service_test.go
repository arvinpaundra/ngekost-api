package kostrule_test

import (
	"context"
	"testing"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	k "github.com/arvinpaundra/ngekost-api/internal/app/kostRule"
	"github.com/arvinpaundra/ngekost-api/internal/driver/contract/mocks"
	"github.com/arvinpaundra/ngekost-api/internal/entity"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/util/common"
)

type ServiceTestTable struct {
	name string
	fn   func(*testing.T)
}

var (
	txBeginer          mocks.TxBeginner
	kostRepository     mocks.KostRepository
	kostRuleRepository mocks.KostRuleRepository

	service k.Service

	createRule request.CreateKostRule
	updateRule request.UpdateKostRule
	query      request.Common
	rulePath   request.KostRulePathParam

	kost entity.Kost
	rule entity.KostRule

	ctx context.Context
)

func initDataService() {
	f := factory.Factory{
		TxBeginner:         &txBeginer,
		KostRepository:     &kostRepository,
		KostRuleRepository: &kostRuleRepository,
	}

	service = k.NewService(&f)

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

	rule = entity.KostRule{
		ID:          common.GetID(),
		KostId:      kost.ID,
		Title:       "test",
		Priority:    "test",
		Description: nil,
	}

	createRule = request.CreateKostRule{
		Title:       "test",
		Priority:    "test",
		Description: nil,
	}

	updateRule = request.UpdateKostRule{
		Title:       "test",
		Priority:    "test",
		Description: nil,
	}

	query = request.Common{
		Limit:  10,
		Offset: 1,
		Search: "",
	}

	rulePath = request.KostRulePathParam{
		KostId:     kost.ID,
		KostRuleId: rule.ID,
	}

	ctx = context.Background()
}

func TestMain(m *testing.M) {

	initDataService()

	m.Run()
}
