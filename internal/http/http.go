package http

import (
	"github.com/arvinpaundra/ngekost-api/internal/app/ping"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/gofiber/fiber/v2"
)

func NewHttp(app *fiber.App, f *factory.Factory) {
	_ = app.Group("/api/v1")

	ping.NewClient(f).Router(app.Group(""))
}
