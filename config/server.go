package config

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func InitServer(app *fiber.App) {
	appEnv := Config("APP_ENV")
	if appEnv == "development" || appEnv == "staging" {
		startServer(app)
	} else {
		startServerWithGracefulShutdown(app)
	}
}

// StartServer func for starting a simple server.
func startServer(a *fiber.App) {
	fiberConnectURL, _ := ConnectionURLBuilder("fiber")

	if err := a.Listen(fiberConnectURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func startServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
