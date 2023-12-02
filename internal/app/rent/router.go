package rent

import "github.com/gofiber/fiber/v2"

func (c *Client) RouterV1(f fiber.Router) {
	f.Post("", c.HandlerCreate)
	f.Patch("/:rent_id", c.HandlerUpdate)
}
