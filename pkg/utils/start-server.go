package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(":3000"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
