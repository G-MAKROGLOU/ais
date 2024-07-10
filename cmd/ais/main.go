package main

import (
	"fmt"
	"log"

	"github.com/G-MAKROGLOU/ais/internal/config"
	"github.com/G-MAKROGLOU/ais/internal/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

type Test struct {
	Message string `json:"message"`
}

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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(Test{
			Message: "Hello World!",
		})
	})
	// register middlewares
	//core.SetupCoreMiddlewares(app)

	// register routes and route specific middlewares

	// listen per environment
	if config.AppConfig.Env != "production" {
		log.Fatalln(app.Listen(fmt.Sprintf(":%s", config.AppConfig.Port)))
	}

	if config.AppConfig.Env == "production" {
		log.Fatalln(app.ListenTLS(":443", "./cert.pem", "./cert.key"))
	}
}
