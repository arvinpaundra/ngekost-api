package ping

import (
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/gofiber/fiber/v2"
)

type client struct {
	service Service
}

func NewClient(f *factory.Factory) *client {
	return &client{
		service: NewService(f),
	}
}

func (c *client) Ping(f *fiber.Ctx) error {
	res, err := c.service.Ping(f.Context())
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return f.Status(fiber.StatusOK).JSON(res)
}
