package http

import (
	"github.com/arvinpaundra/ngekost-api/internal/app/auth"
	"github.com/arvinpaundra/ngekost-api/internal/app/kost"
	kostrule "github.com/arvinpaundra/ngekost-api/internal/app/kostRule"
	"github.com/arvinpaundra/ngekost-api/internal/app/ping"
	"github.com/arvinpaundra/ngekost-api/internal/app/room"
	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/gofiber/fiber/v2"
)

func NewHttp(app *fiber.App, f *factory.Factory) {
	v1 := app.Group("/api/v1")

	auth.NewClient(f).RouterV1(v1.Group("/auth"))
	kost.NewClient(f).RouterV1(v1.Group("/kosts"))
	room.NewClient(f).RouterV1(v1.Group("/kosts/:kost_id/rooms"))
	kostrule.NewClient(f).RouterV1(v1.Group("/kosts/:kost_id/rules"))

	ping.NewClient(f).Router(app.Group(""))
}
