package componentelistener

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitListener struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitListener() (*RabbitListener, error) {
	// Mismos datos de conexión usados en ServidorDeStreaming
	conn, err := amqp.Dial("amqp://admin:1234@192.168.80.25:5672/")
	if err != nil {
		return nil, fmt.Errorf("Error conectando a RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Error abriendo el canal: %v", err)
	}

	// Declarar exactamente la misma cola que el publicador
	q, err := ch.QueueDeclare(
		"cola_estadisticas", // nombre de la cola
		true,                // durable
		false,               // autodelete
		false,               // exclusive
		false,               // no-wait
		nil,                 // args
	)

	if err != nil {
		return nil, fmt.Errorf("Error declarando la cola: %v", err)
	}

	return &RabbitListener{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (l *RabbitListener) GetChannel() *amqp.Channel {
	return l.channel
}

func (l *RabbitListener) GetQueueName() string {
	return l.queue.Name
}

func (l *RabbitListener) Cerrar() {
	if l.channel != nil {
		l.channel.Close()
	}
	if l.conn != nil {
		l.conn.Close()
	}
}
