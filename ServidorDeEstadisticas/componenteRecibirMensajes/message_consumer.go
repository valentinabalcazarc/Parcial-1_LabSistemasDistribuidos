/**
 * @file message_consumer.go
 * @brief Consumidor de mensajes para procesar notificaciones desde RabbitMQ.
 */
package componenterecibirmensajes

import (
	"encoding/json"
	"fmt"
	"log"

	dtos "servidorEstadisticas/DTOs"
	listener "servidorEstadisticas/componenteListener"
)

/**
 * @struct MessageConsumer
 * @brief Estructura que gestiona el consumo de mensajes de RabbitMQ.
 */
type MessageConsumer struct {
	rabbitListener *listener.RabbitListener /**< @brief Referencia al listener de RabbitMQ para acceder al canal y cola. */
}

/**
 * @brief Crea una nueva instancia de MessageConsumer.
 * 
 * @param rabbitListener Puntero al listener de RabbitMQ configurado.
 * @return Un puntero a la instancia de MessageConsumer creada.
 */
func NewMessageConsumer(rabbitListener *listener.RabbitListener) *MessageConsumer {
	return &MessageConsumer{
		rabbitListener: rabbitListener,
	}
}

/**
 * @brief Inicia el consumo de mensajes desde la cola de RabbitMQ.
 * 
 * Se suscribe a la cola "cola_estadisticas", procesa las notificaciones de reproducción
 * en formato JSON y las imprime en la consola con un formato específico.
 * 
 * @return Un error si ocurre un problema al registrar el consumidor, o nil en caso de éxito.
 */
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
