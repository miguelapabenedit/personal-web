package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/services"
)

type controller struct{}

var serv services.Service

func NewGratitudeController(service services.Service) Controller {
	serv = service
	return &controller{}
}

func (c *controller) GetGratitudes(ctx *fiber.Ctx) error {
	gratitudes, err := serv.GetGratitudes()

	if err != nil {
		ctx.Context().Error("Internal Server Error", 500)
		return err
	}

	ctx.JSON(gratitudes)
	return nil
}
