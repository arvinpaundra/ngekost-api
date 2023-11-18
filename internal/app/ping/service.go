package ping

import (
	"context"
	"os"

	"github.com/arvinpaundra/ngekost-api/internal/adapter/response"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
)

type Service interface {
	Ping(ctx context.Context) (*response.Ping, error)
}

type service struct {
}

func NewService(f *factory.Factory) Service {
	return &service{}
}

func (p *service) Ping(ctx context.Context) (*response.Ping, error) {
	version, err := os.ReadFile("version.txt")
	if err != nil {
		return nil, err
	}

	res := response.Ping{
		Name:    "ngekost-api",
		Version: string(version),
		Author:  "Arvin Paundra Ardana",
		Github:  "https://github.com/arvinpaundra/ngekost-api",
	}

	return &res, nil
}
