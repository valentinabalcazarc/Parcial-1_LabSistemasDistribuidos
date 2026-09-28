package componenteconexioncola

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

type NotificacionReproduccion struct {
	Titulo    string `json:"titulo"`
	TipoAudio string `json:"tipo_audio,omitempty"`
	FechaHora string `json:"fecha_hora,omitempty"`
	Mensaje   string `json:"mensaje"`
}

func NewRabbitPublisher() (*RabbitPublisher, error) {
	conn, err := amqp.Dial("amqp://admin:1234@192.168.80.25:5672/")
	if err != nil {
		return nil, fmt.Errorf("Error conectando a rabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Error abriendo el canal: %v", err)
	}

	q, err := ch.QueueDeclare(
		"cola_estadisticas", //nombre de la cola
		true,                //durable
		false,               //autodelete
		false,               //exclusive
		false,               //no-wait
		nil,                 //args
	)

	if err != nil {
		return nil, fmt.Errorf("Error declarando la cola: %v", err)
	}

	return &RabbitPublisher{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (p *RabbitPublisher) PublicarNotificacion(msg NotificacionReproduccion) error {
	body, err := json.Marshal(msg)

	if err != nil {
		return fmt.Errorf("Error convirtiendo a JSON: %v", err)
	}

	err = p.channel.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return fmt.Errorf("Error publicando el mensaje: %v", err)
	}

	fmt.Println("Notificación enviada a RabbitMQ: ", string(body))
	return nil
}

func (p *RabbitPublisher) Cerrar() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}
