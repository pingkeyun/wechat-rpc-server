package config

// queue basic info
type QueueBase struct {
	Host string
	Port string
	User string
	Password string
}
var QueueBaseConfig = &QueueBase{}

// rabbitmq
type QueueRabbitMq struct {
	QueueBase
	Vhost string
	Exchange string
	Queue string
	RoutingKey string
}
var QueueRabbitMqConfig = &QueueRabbitMq{}
