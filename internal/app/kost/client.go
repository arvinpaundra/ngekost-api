package kost

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
	var req request.CreateKost

	_ = f.BodyParser(&req)

	validationErrors := validator.Validate(req)
	if validationErrors != nil {
		return f.Status(fiber.StatusBadRequest).JSON(format.BadRequest("invalid request body", validationErrors))
	}

	err := c.service.Save(f.Context(), &req, nil)

	if err != nil {
		switch err {
		case constant.ErrOwnerNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusCreated).JSON(format.SuccessCreated("success create kost", nil))
}

func (c *Client) HandlerUpdate(f *fiber.Ctx) error {
	var req request.UpdateKost

	kostId := f.Params("kost_id")

	_ = f.BodyParser(&req)

	validationErrors := validator.Validate(req)
	if validationErrors != nil {
		return f.Status(fiber.StatusBadRequest).JSON(format.BadRequest("invalid request body", validationErrors))
	}

	err := c.service.Update(f.Context(), kostId, &req, nil)

	if err != nil {
		switch err {
		case constant.ErrKostNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success update kost", nil))
}

func (c *Client) HandlerFindAll(f *fiber.Ctx) error {
	query := request.Common{
		Limit:  f.QueryInt("per_page", 10),
		Offset: f.QueryInt("page", 1),
		Search: f.Query("search"),
	}

	res, err := c.service.FindAll(f.Context(), &query)

	if err != nil {
		switch err {
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success get kosts", res.Results, res.Pagination))
}

func (c *Client) HandlerFindByOwner(f *fiber.Ctx) error {
	query := request.Common{
		Limit:  f.QueryInt("per_page", 10),
		Offset: f.QueryInt("page", 1),
		Search: f.Query("search"),
	}

	ownerId := f.Params("owner_id")

	res, err := c.service.FindByOwnerId(f.Context(), ownerId, &query)

	if err != nil {
		switch err {
		case constant.ErrOwnerNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success get kosts by owner", res.Results, res.Pagination))
}

func (c *Client) HandlerFindById(f *fiber.Ctx) error {
	kostId := f.Params("kost_id")

	res, err := c.service.FindById(f.Context(), kostId)

	if err != nil {
		switch err {
		case constant.ErrKostNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success get kost by id", res))
}

func (c *Client) HandlerFindRoomsByKost(f *fiber.Ctx) error {
	path := request.RoomPathParam{
		KostId: f.Params("kost_id"),
	}

	query := request.Common{
		Limit:  f.QueryInt("per_page", 10),
		Offset: f.QueryInt("page", 1),
		Search: f.Query("search"),
	}

	res, err := c.service.FindRoomsByKost(f.Context(), &path, &query)

	if err != nil {
		switch err {
		case constant.ErrKostNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success get rooms by kost", res.Results, res.Pagination))
}

func (c *Client) HandlerDelete(f *fiber.Ctx) error {
	kostId := f.Params("kost_id")

	err := c.service.Delete(f.Context(), kostId)

	if err != nil {
		switch err {
		case constant.ErrKostNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success delete kost by id", nil))
}

func (c *Client) HandlerFindRulesByKost(f *fiber.Ctx) error {
	path := request.KostRulePathParam{
		KostId: f.Params("kost_id"),
	}

	query := request.Common{
		Limit:  f.QueryInt("per_page", 10),
		Offset: f.QueryInt("page", 1),
	}

	res, err := c.service.FindRulesByKost(f.Context(), &path, &query)

	if err != nil {
		switch err {
		case constant.ErrKostNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("success get rules of kost", res))
}
