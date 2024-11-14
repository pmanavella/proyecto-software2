package queues

import (
	"encoding/json"
	"fmt"
	"log"
	courses "search-api/dao"

	"github.com/streadway/amqp"
)

type RabbitConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	QueueName string
}

type Rabbit struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

// NewRabbit initializes a new RabbitMQ client
func NewRabbit(config RabbitConfig) (*Rabbit, error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	queue, err := channel.QueueDeclare(
		config.QueueName, // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %w", err)
	}

	return &Rabbit{
		conn:    conn,
		channel: channel,
		queue:   queue,
	}, nil
}

// StartConsumer starts consuming messages from the RabbitMQ queue
func (r *Rabbit) StartConsumer(handler func(courses.CourseNew)) error {

	msgs, err := r.channel.Consume(
		r.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	go func() {
		for d := range msgs {
			var courseNew courses.CourseNew
			if err := json.Unmarshal(d.Body, &courseNew); err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}
			handler(courseNew)
		}
	}()

	return nil
}

// Close closes the RabbitMQ connection and channel
func (r *Rabbit) Close() {
	if err := r.channel.Close(); err != nil {
		log.Printf("Error closing channel: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}
