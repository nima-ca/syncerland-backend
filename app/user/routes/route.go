package routes

import "github.com/gofiber/fiber/v2"

func Register(router *fiber.App) {
	userGroup := router.Group("/user")

	userGroup.Get("/register", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})
}
