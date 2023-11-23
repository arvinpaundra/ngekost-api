package auth

import "github.com/gofiber/fiber/v2"

func (c *client) RouterV1(f fiber.Router) {
	f.Post("/register", c.HandlerRegister)
	f.Post("/login", c.HandlerLogin)
	f.Post("/logout", c.HandlerLogout)
}
