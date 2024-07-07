package main

import (
	"github.com/G-MAKROGLOU/ais/internal/config"
)

// ENV is env
func main() {
	// load env

	config.LoadEnv()

	// init db

	// init rabbitmq
	//if err := rabbitmq.RMQConnect(""); err != nil {
	//	log.Println("RabbitMQ connection error: ", err.Error())
	//}
	//defer rabbitmq.RabbitMQ.Close()

	//// init app
	//app := fiber.New(fiber.Config{
	//	CaseSensitive:     true,
	//	ETag:              true,
	//	EnablePrintRoutes: true,
	//	StrictRouting:     true,
	//	UnescapePath:      true,
	//})

	//// register middlewares
	//SetupCoreMiddlewares(app)

	//// register routes and route specific middlewares

	//// listen per environment
	//if ENV == "development" {
	//	log.Fatalln(app.Listen(":3000"))
	//}

	//if ENV == "staging" {
	//	log.Fatalln(app.Listen(":80"))
	//}

	//if ENV == "production" {
	//	log.Fatalln(app.ListenTLS(":443", "./cert.pem", "./cert.key"))
	//}
}
