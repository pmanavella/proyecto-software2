package queues

import (
    "encoding/json"
    "fmt"
    "github.com/streadway/amqp"
    "log"
    "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type RabbitConfig struct {
    Host      string
    Port      string
    Username  string
    Password  string
    QueueName string
}

type Rabbit struct {
    connection *amqp.Connection
    channel    *amqp.Channel
    queue      amqp.Queue
}

func NewRabbit(config RabbitConfig) Rabbit {
    connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port))
    if err != nil {
        log.Fatal(err)
    }

    channel, err := connection.Channel()
    if err != nil {
        log.Fatal(err)
    }

    queue, err := channel.QueueDeclare(
        config.QueueName,
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }

    return Rabbit{
        connection: connection,
        channel:    channel,
        queue:      queue,
    }
}

func (queue Rabbit) Publish(courseNew hotels.CourseNew) error {
	bytes, err := json.Marshal(courseNew)
	if err != nil {
		return fmt.Errorf("error marshaling Rabbit courseNew: %w", err)
	}
	if err := queue.channel.Publish(
		"",
		queue.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		}); err != nil {
		return fmt.Errorf("error publishing to Rabbit: %w", err)
	}
	return nil
}

// Close cleans up the RabbitMQ resources
func (queue Rabbit) Close() {
	if err := queue.channel.Close(); err != nil {
		log.Printf("error closing Rabbit channel: %v", err)
	}
	if err := queue.connection.Close(); err != nil {
		log.Printf("error closing Rabbit connection: %v", err)
	}
}

//////

func (r *Rabbit) Notify(message interface{}) error {
    body, err := json.Marshal(message)
    if err != nil {
        return err
    }

    err = r.channel.Publish(
        "",
        r.queue.Name,
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
    return err
}

// creo este nuevo archivo que indica que esta parte del sistema
// interactúa con clientes externos (en este caso, RabbitMQ).

var rabbitConn *amqp.Connection

// InitRabbitMQ crea una conexión global con RabbitMQ
func InitRabbitMQ() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Errorf("Error al conectarse a RabbitMQ: %v", err)
		return err
	}
	rabbitConn = conn

	log.Info("Conexión a RabbitMQ establecida")
	return nil
}

// devuelve la conexión activa a RabbitMQ
func GetRabbitConnection() *amqp.Connection {
	return rabbitConn
}

// CloseRabbitMQ cierra la conexión con RabbitMQ al finalizar el uso
func CloseRabbitMQ() {
	if rabbitConn != nil {
		rabbitConn.Close()
		log.Info("Conexión a RabbitMQ cerrada")
	}
}
