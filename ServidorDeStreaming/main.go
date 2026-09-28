package main

import (
	"fmt"
	"net"

	capaControladores "servidorStreaming/capacontrollers"
	pb "servidorStreaming/serviciosAudio"

	"google.golang.org/grpc"
)

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
