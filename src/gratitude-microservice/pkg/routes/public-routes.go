package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/controllers"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/services"
)

var (
	service    services.Service       = services.NewGratitudeService()
	controller controllers.Controller = controllers.NewGratitudeController(service)
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/gratitudes", controller.GetGratitudes)

	route.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendString("The app is running")
	})
}
