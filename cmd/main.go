package main

import (
	"github/miguelapabenedit/personal-page/pkg/middlewares"
	"github/miguelapabenedit/personal-page/pkg/routes"
	"github/miguelapabenedit/personal-page/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middlewares.FiberMiddleware(app)

	routes.PublicRoutes(app)

	utils.StartServer(app)
}
