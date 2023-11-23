package http

import (
	"github.com/arvinpaundra/ngekost-api/internal/app/auth"
	"github.com/arvinpaundra/ngekost-api/internal/app/ping"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/gofiber/fiber/v2"
)

func NewHttp(app *fiber.App, f *factory.Factory) {
	v1 := app.Group("/api/v1")

	auth.NewClient(f).RouterV1(v1.Group("/auth"))

	ping.NewClient(f).Router(app.Group(""))
}
