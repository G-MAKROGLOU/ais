package main

import (
	"log"
	"strings"

	"github.com/G-MAKROGLOU/ais/internal/config"
	"github.com/G-MAKROGLOU/ais/internal/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// load env
	config.LoadEnv()

	// init db

	// init rabbitmq
	if err := rabbitmq.RMQConnect(""); err != nil {
		log.Println("RabbitMQ connection error: ", err.Error())
	}
	defer rabbitmq.RabbitMQ.Close()

	// init app
	app := fiber.New(fiber.Config{
		CaseSensitive:     true,
		ETag:              true,
		EnablePrintRoutes: true,
		StrictRouting:     true,
		UnescapePath:      true,
	})

	// register middlewares
	SetupCoreMiddlewares(app)

	// register routes and route specific middlewares

	// listen per environment
	if config.AppConfig.Env == "development" || config.AppConfig.Env == "staging" || strings.TrimSpace(config.AppConfig.Env) == "" {
		log.Fatalln(app.Listen(config.AppConfig.Port))
	}

	if config.AppConfig.Env == "production" {
		log.Fatalln(app.ListenTLS(":443", "./cert.pem", "./cert.key"))
	}
}
