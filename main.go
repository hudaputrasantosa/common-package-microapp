package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudaputrasantosa/common-package-microapp/config"
)

func main() {
	app := fiber.New()

	// Start server
	config.InitServer(app)

}
