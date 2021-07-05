package controllers

import "github.com/gofiber/fiber/v2"

func GetGratitudes(ctx *fiber.Ctx) error {
	ctx.SendString("Here are all the gratitudes")
	return nil
}
