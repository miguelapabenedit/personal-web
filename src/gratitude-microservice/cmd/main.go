package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/pkg/middlewares"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/pkg/routes"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/pkg/utils"
)

func main() {
	app := fiber.New()

	middlewares.FiberMiddleware(app)

	routes.PublicRoutes(app)

	utils.StartServer(app)
}
