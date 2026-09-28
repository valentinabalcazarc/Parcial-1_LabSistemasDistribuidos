package componenterecibirmensajes

import (
	"encoding/json"
	"fmt"
	"log"

	dtos "servidorEstadisticas/DTOs"
	listener "servidorEstadisticas/componenteListener"
)

type MessageConsumer struct {
	rabbitListener *listener.RabbitListener
}

func NewMessageConsumer(rabbitListener *listener.RabbitListener) *MessageConsumer {
	return &MessageConsumer{
		rabbitListener: rabbitListener,
	}
}

// IniciarConsumo se suscribe a la cola y procesa las notificaciones de reproducción
func (mc *MessageConsumer) IniciarConsumo() error {
	ch := mc.rabbitListener.GetChannel()
	queueName := mc.rabbitListener.GetQueueName()

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer tag
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return fmt.Errorf("Error al registrar el consumidor: %v", err)
	}

	log.Println(" [*] Esperando mensajes de reproducción en 'cola_estadisticas'. Para salir presiona CTRL+C")

	go func() {
		for d := range msgs {
			var notificacion dtos.NotificacionReproduccion
			err := json.Unmarshal(d.Body, &notificacion)
			if err != nil {
				log.Printf("Error decodificando mensaje JSON: %v", err)
				continue
			}

			// Imprimir por pantalla según el requerimiento del parcial
			fmt.Println("\n==================================================")
			fmt.Println("   📊 [SERVIDOR DE ESTADÍSTICAS - NUEVO EVENTO]")
			fmt.Printf("   🎵 Audio:      %s\n", notificacion.Titulo)
			if notificacion.TipoAudio != "" {
				fmt.Printf("   🏷️  Tipo:       %s\n", notificacion.TipoAudio)
			}
			fmt.Printf("   🕒 Fecha/Hora: %s\n", notificacion.FechaHora)
			fmt.Printf("   💬 Mensaje:    %s\n", notificacion.Mensaje)
			fmt.Println("==================================================")
		}
	}()

	return nil
}
