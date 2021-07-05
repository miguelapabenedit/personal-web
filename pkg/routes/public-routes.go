package routes

import (
	"github/miguelapabenedit/personal-page/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/gratitude", controllers.GetGratitudes)

	route.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendString("The app is running")
	})
}
