package service

import (
	"fmt"
	"github.com/pingkeyun/wechat-rpc-server/internal/app/config"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

type WechatPushServer struct {
	conn *amqp.Connection
	ch *amqp.Channel
}

var once sync.Once
var singletonWechatPushServer *WechatPushServer

func NewWechatPushServer() *WechatPushServer {
	once.Do(func() {
		// 连接rabbitmq
		dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.QueueBaseConfig.User,
			config.QueueBaseConfig.Password,
			config.QueueRabbitMqConfig.Host,
			config.QueueRabbitMqConfig.Port)
		conn, err := amqp.Dial(dsn)
		if err != nil {
			log.Fatal("连接rabbitmq失败：", err)
		}

		ch, err := conn.Channel()
		failOnError(err, "Failded to open a channel")

		singletonWechatPushServer = &WechatPushServer{
			conn: conn,
			ch: ch,
		}
	})

	return singletonWechatPushServer
}

func (w *WechatPushServer)Close() {
	w.ch.Close()
	w.conn.Close()
}

func WechatPushPut(body string) string {
	srv := NewWechatPushServer()
	_, err := srv.ch.QueueDeclare(
		config.QueueRabbitMqConfig.Queue, // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = srv.ch.Publish(
		config.QueueRabbitMqConfig.Exchange,           // exchange
		config.QueueRabbitMqConfig.RoutingKey,       // routing key
		false,        // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
	}
	log.Printf(" [x] Sent %s", body)

	return "ok"
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}