package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/internal/http"
	"github.com/arvinpaundra/ngekost-api/pkg/util/common"
	"github.com/arvinpaundra/ngekost-api/pkg/util/config"
	l "github.com/arvinpaundra/ngekost-api/pkg/util/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()

	f := factory.NewFactory(ctx)

	app := fiber.New(fiber.Config{})

	http.NewHttp(app, f)

	go func() {
		if err := app.Listen(config.GetString("APP_ADDR")); err != nil {
			l.Logging().Error(err.Error())

			os.Exit(0)
		}
	}()

	log.Println("application started")

	wait := common.GracefulShutdown(ctx, 10*time.Second, app.ShutdownWithContext)

	<-wait

	log.Println("application stopped")
}
