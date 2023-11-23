package auth

import (
	"github.com/arvinpaundra/ngekost-api/internal/adapter/request"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/pkg/constant"
	"github.com/arvinpaundra/ngekost-api/pkg/helper/format"
	"github.com/arvinpaundra/ngekost-api/pkg/helper/validator"
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

func (c *client) HandlerRegister(f *fiber.Ctx) error {
	var req request.Register

	_ = f.BodyParser(&req)

	validationErrors := validator.Validate(req)
	if validationErrors != nil {
		return f.Status(fiber.StatusBadRequest).JSON(format.BadRequest("invalid request body", validationErrors))
	}

	err := c.service.Register(f.Context(), &req)

	if err != nil {
		switch err {
		case constant.ErrInvalidRole:
			return f.Status(fiber.StatusUnprocessableEntity).JSON(format.UnprocessableEntity(err.Error()))
		case constant.ErrUsernameAlreadyUsed:
			return f.Status(fiber.StatusConflict).JSON(format.Conflict(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusCreated).JSON(format.SuccessCreated("register successful", nil))
}

func (c *client) HandlerLogin(f *fiber.Ctx) error {
	var req request.Login

	_ = f.BodyParser(&req)

	validationErrors := validator.Validate(req)
	if validationErrors != nil {
		return f.Status(fiber.StatusBadRequest).JSON(format.BadRequest("invalid request body", validationErrors))
	}

	res, err := c.service.Login(f.Context(), &req)

	if err != nil {
		switch err {
		case constant.ErrUserNotFound, constant.ErrPasswordIncorrect:
			return f.Status(fiber.StatusBadRequest).JSON(format.BadRequest("username or password incorrect", validator.ValidationError{
				"username": "username or password incorrect",
				"password": "username or password incorrect",
			}))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("login successful", res))
}

func (c *client) HandlerLogout(f *fiber.Ctx) error {
	var req request.Logout

	_ = f.BodyParser(&req)

	validationErrors := validator.Validate(req)
	if validationErrors != nil {
		return f.Status(fiber.StatusBadRequest).JSON(format.BadRequest("invalid request body", validationErrors))
	}

	err := c.service.Logout(f.Context(), &req)

	if err != nil {
		switch err {
		case constant.ErrUserNotFound, constant.ErrSessionNotFound:
			return f.Status(fiber.StatusNotFound).JSON(format.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(format.SuccessOK("logout successful", nil))
}
