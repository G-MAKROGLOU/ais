package config

// Config represents the app configuration
type Config struct {
	Port             string
	Env              string
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresPort     string
	PostgresDB       string
	RabbitMQHost     string
	RabbitMQUser     string
	RabbitMQPassword string
	RabbitMQPort     string
	RabbitMQTopic    string
}
