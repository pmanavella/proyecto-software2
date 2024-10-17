package queues

import (
	"context"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

// Estructura queue para manejar la cola de RabbitMQ
type queue struct {
	Name    string
	channel *amqp.Channel
}

// Interfaz para la cola que expone los métodos necesarios
type queueInterface interface {
	// Establece la conexión con RabbitMQ, abre el canal, y declara la cola para que esté lista para recibir mensajes.
	InitQueue(queueName string, conn *amqp.Connection) error 
	Publish(body []byte) error
}

// Variable global Queue para acceder a la implementación
var Queue queueInterface

// Inicializa la variable Queue al iniciar el paquete
func init() {
	Queue = &queue{}
}

// InitQueue inicializa la cola, crea el canal y declara la cola en RabbitMQ
func (q *queue) InitQueue(queueName string, conn *amqp.Connection) error {
	// Crear canal
	ch, err := conn.Channel()
	if err != nil {
		log.Error("Error al abrir el canal", err)
		return err
	}
	q.channel = ch
	q.Name = queueName

	// Declarar la cola en RabbitMQ
	_, err = ch.QueueDeclare(
		queueName, // Nombre de la cola
		true,      // Durable (la cola persiste)
		false,     // Auto-delete cuando no se usa
		false,     // Exclusiva
		false,     // No-wait
		nil,       // Argumentos adicionales
	)
	if err != nil {
		log.Error("Error al declarar la cola", err)
		return err
	}

	log.Infof("Cola '%s' declarada correctamente", queueName)
	return nil
}

// Publish envía un mensaje a la cola de RabbitMQ
func (q *queue) Publish(body []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publicar mensaje en la cola de RabbitMQ
	err := q.channel.PublishWithContext(
		ctx,
		"",       // Exchange (vacío para usar cola por defecto)
		q.Name,   // Routing key (nombre de la cola)
		false,    // Mandatory
		false,    // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		log.Error("Error al publicar el mensaje en la cola", err)
		return err
	}

	log.Infof("Mensaje publicado correctamente en la cola '%s'", q.Name)
	return nil
}
