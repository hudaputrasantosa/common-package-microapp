package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudaputrasantosacommon-package-microapp/config"
)

func main() {
	app := fiber.New()

	// Start server
	config.InitServer(app)

}
