package kostrule

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
	var req request.CreateKostRule

	_ = f.BodyParser(&req)

	validationErrors := validator.Validate(req)
	if validationErrors != nil {
		return f.Status(fiber.StatusBadRequest).JSON(format.BadRequest("invalid request body", validationErrors))
	}

	path := request.KostRulePathParam{
		KostId: f.Params("kost_id"),
	}

	err := c.service.Save(f.Context(), &path, &req)

	if err != nil {
		switch err {
		case constant.ErrKostNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusCreated).JSON(format.SuccessCreated("success create kost rule", nil))
}

func (c *Client) HandlerUpdate(f *fiber.Ctx) error {
	var req request.UpdateKostRule

	_ = f.BodyParser(&req)

	validationErrors := validator.Validate(req)
	if validationErrors != nil {
		return f.Status(fiber.StatusBadRequest).JSON(format.BadRequest("invalid request body", validationErrors))
	}

	path := request.KostRulePathParam{
		KostId:     f.Params("kost_id"),
		KostRuleId: f.Params("rule_id"),
	}

	err := c.service.Update(f.Context(), &path, &req)

	if err != nil {
		switch err {
		case constant.ErrKostNotFound, constant.ErrKostRuleNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success update kost rule", nil))
}

func (c *Client) HandlerDelete(f *fiber.Ctx) error {
	path := request.KostRulePathParam{
		KostId:     f.Params("kost_id"),
		KostRuleId: f.Params("rule_id"),
	}

	err := c.service.Delete(f.Context(), &path)

	if err != nil {
		switch err {
		case constant.ErrKostNotFound, constant.ErrKostRuleNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success delete kost rule", nil))
}

func (c *Client) HandlerFindById(f *fiber.Ctx) error {
	path := request.KostRulePathParam{
		KostId:     f.Params("kost_id"),
		KostRuleId: f.Params("rule_id"),
	}

	res, err := c.service.FindById(f.Context(), &path)

	if err != nil {
		switch err {
		case constant.ErrKostNotFound, constant.ErrKostRuleNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success get kost rule by id", res))
}
