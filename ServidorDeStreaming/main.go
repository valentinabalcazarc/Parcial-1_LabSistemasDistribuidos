/**
 * @file main.go
 * @brief Archivo principal que inicia el servidor de streaming gRPC.
 */
package main

import (
	"fmt"
	"net"

	capaControladores "servidorStreaming/capacontrollers"
	pb "servidorStreaming/serviciosAudio"

	"google.golang.org/grpc"
)

/**
 * @brief Función principal que inicializa y levanta el servidor gRPC para streaming de audio.
 * @details Escucha en el puerto TCP 8083 y registra el servicio de audio.
 */
func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAudioServiceServer(grpcServer, &capaControladores.ControladorServidor{})

	fmt.Println("Servidor gRPC escuchando en :8083...")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
