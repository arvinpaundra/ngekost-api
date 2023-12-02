package rent

import (
	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/helper/format"
	"github.com/arvinpaundra/ngekost-api/pkg/helper/validator"
	"github.com/gofiber/fiber/v2"
)

type Client struct {
	service Service
}

func NewClient(f *factory.Factory) *Client {
	return &Client{
		service: NewService(f),
	}
}

func (c *Client) HandlerCreate(f *fiber.Ctx) error {
	var req request.CreateRent

	_ = f.BodyParser(&req)

	validationErrors := validator.Validate(req)
	if validationErrors != nil {
		return f.Status(fiber.StatusBadRequest).JSON(format.BadRequest("invalid request body", validationErrors))
	}

	err := c.service.Save(f.Context(), &req)

	if err != nil {
		switch err {
		case constant.ErrLesseeNotFound, constant.ErrRoomNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		case constant.ErrActiveRent:
			return f.Status(fiber.StatusUnprocessableEntity).JSON(format.UnprocessableEntity(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusCreated).JSON(format.SuccessCreated("success renting kost", nil))
}

func (c *Client) HandlerUpdate(f *fiber.Ctx) error {
	rentId := f.Params("rent_id")

	err := c.service.Update(f.Context(), rentId)

	if err != nil {
		switch err {
		case constant.ErrRentNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success update status rent", nil))
}
