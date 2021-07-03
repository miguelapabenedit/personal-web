package controllers

import "github.com/gofiber/fiber/v2"

func GetAll(ctx *fiber.Ctx) {
	ctx.SendString("Here are all the gratitudes")
}
