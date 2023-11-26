package room

import "github.com/gofiber/fiber/v2"

func (c *Client) RouterV1(f fiber.Router) {
	f.Post("", c.HandlerCreate)
	f.Get("/:room_id", c.HandlerFindById)
	f.Put("/:room_id", c.HandlerUpdate)
	f.Delete("/:room_id", c.HandlerDelete)
}
