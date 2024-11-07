package queues

import (
    "encoding/json"
    "fmt"
    "github.com/streadway/amqp"
    "log"
    "courses-api/dto/courses"
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

func (r *Rabbit) Notify(message courses.CourseResponse_Full) error {
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