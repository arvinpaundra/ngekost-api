package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/arvinpaundra/ngekost-api/internal/factory"
	"github.com/arvinpaundra/ngekost-api/internal/http"
	"github.com/arvinpaundra/ngekost-api/pkg/util/config"
	"github.com/arvinpaundra/ngekost-api/pkg/util/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	f := factory.NewFactory()

	app := fiber.New(fiber.Config{})

	http.NewHttp(app, f)

	ch := make(chan os.Signal, 1)

	defer close(ch)

	go func(ch chan os.Signal) {
		if err := app.Listen(config.GetString("APP_ADDR")); err != nil {
			log.Logging(err.Error()).Error()

			os.Exit(1)
		}

		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	}(ch)

	fmt.Println("application started")
	<-ch
}
