package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/controllers"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/repository"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/services"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/pkg/infrastructure"
)

var (
	repo       repository.Repository  = infrastructure.NewPostgreSQL()
	service    services.Service       = services.NewGratitudeService(repo)
	controller controllers.Controller = controllers.NewGratitudeController(service)
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/gratitudes", controller.GetGratitudes)

	route.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendString("The app is running")
	})
}
