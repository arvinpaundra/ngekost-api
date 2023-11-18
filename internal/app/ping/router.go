package ping

import "github.com/gofiber/fiber/v2"

func (c *client) Router(f fiber.Router) {
	f.Get("", c.Ping)
}
