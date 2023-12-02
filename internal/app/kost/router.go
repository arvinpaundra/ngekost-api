package kost

import "github.com/gofiber/fiber/v2"

func (c *Client) RouterV1(f fiber.Router) {
	f.Post("", c.HandlerCreate)
	f.Get("", c.HandlerFindAll)
	f.Get("/:kost_id", c.HandlerFindById)
	f.Put("/:kost_id", c.HandlerUpdate)
	f.Delete("/:kost_id", c.HandlerDelete)

	f.Get("/:kost_id/lessees", c.HandlerFindLesseesByKost)
	f.Get("/:kost_id/rooms", c.HandlerFindRoomsByKost)
	f.Get("/:kost_id/rules", c.HandlerFindRulesByKost)
}
