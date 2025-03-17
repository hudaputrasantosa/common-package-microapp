package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string, fallback ...string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error load env")
	}
	value := os.Getenv(key)
	if value == "" && len(fallback) > 0 {
		return fallback[0]
	}
	return value

}

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(name string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch name {
	case "redis":
		url = fmt.Sprintf(
			"%s:%s",
			Config("REDIS_HOST"),
			Config("REDIS_PORT"),
		)
	case "fiber":
		url = fmt.Sprintf(
			"%s:%s",
			Config("SERVER_HOST"),
			Config("SERVER_PORT"),
		)
	case "cloudinary":
		url = fmt.Sprintf(
			"cloudinary://%s:%s@%s",
			Config("CLOUDINARY_API_KEY"),
			Config("CLOUDINARY_API_SECRET_KEY"),
			Config("CLOUDINARY_CLOUD_NAME"),
		)
	case "rabbitmq":
		url = fmt.Sprintf(
			"amqp://%s:%s@%s:%s",
			Config("RABBITMQ_USER"),
			Config("RABBITMQ_PASSWORD"),
			Config("RABBITMQ_HOST"),
			Config("RABBITMQ_PORT"),
		)
	default:
		return "", fmt.Errorf("connection name '%v' is not supported", name)
	}

	return url, nil
}
