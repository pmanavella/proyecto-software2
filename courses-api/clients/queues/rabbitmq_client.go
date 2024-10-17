package queues

import (
	"github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)
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
