package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	GetGratitudes(ctx *fiber.Ctx) error
}
