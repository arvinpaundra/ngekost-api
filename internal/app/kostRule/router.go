package kostrule

import "github.com/gofiber/fiber/v2"

func (c *Client) RouterV1(f fiber.Router) {
	f.Post("", c.HandlerCreate)
	f.Get("/:rule_id", c.HandlerFindById)
	f.Put("/:rule_id", c.HandlerUpdate)
	f.Delete("/:rule_id", c.HandlerDelete)
}
