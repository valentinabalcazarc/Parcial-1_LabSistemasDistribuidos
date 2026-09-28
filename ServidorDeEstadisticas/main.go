package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	listener "servidorEstadisticas/componenteListener"
	consumer "servidorEstadisticas/componenteRecibirMensajes"
)

func main() {
	fmt.Println("Iniciando Servidor de Estadísticas...")

	rabbitListener, err := listener.NewRabbitListener()
	if err != nil {
		log.Fatalf("Fallo crítico iniciando escuchador de RabbitMQ: %v", err)
	}
	defer rabbitListener.Cerrar()

	messageConsumer := consumer.NewMessageConsumer(rabbitListener)

	// Iniciar la escucha continua de mensajes
	err = messageConsumer.IniciarConsumo()
	if err != nil {
		log.Fatalf("Fallo iniciando consumo de mensajes: %v", err)
	}

	// Mantener el proceso Go vivo escuchando señales de interrupción (CTRL+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("\nApagando Servidor de Estadísticas...")
}
