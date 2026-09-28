package clientestreaming

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "servidorStreaming/serviciosAudio"
)

const urlStreaming = "localhost:8083"

func ReproducirAudio(filename string) {
	fmt.Printf("\n📡 Conectando al Servidor de Streaming (%s)...\n", urlStreaming)

	conn, err := grpc.Dial(urlStreaming, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("❌ Error al conectar con servidor gRPC: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewAudioServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	stream, err := client.AudioStream(ctx, &pb.AudioRequest{Filename: filename})
	if err != nil {
		fmt.Printf("❌ Error al iniciar el stream: %v\n", err)
		return
	}

	reader, writer := io.Pipe()
	stopChan := make(chan struct{})
	doneChan := make(chan struct{})

	// Goroutine que escucha la tecla ENTER para detener la reproducción en cualquier momento
	go func() {
		fmt.Printf("▶️  Reproduciendo %s en tiempo real vía gRPC...\n", filename)
		fmt.Println("👉 Presione ENTER en cualquier momento para detener la reproducción y volver al menú.")
		bufio.NewReader(os.Stdin).ReadString('\n')
		close(stopChan)
	}()

	// Decodificar y reproducir sonido en el speaker
	go DecodificarReproducir(reader, stopChan, doneChan)

	// Recibir fragmentos de gRPC y escribirlos al reproductor
	RecibirAudio(stream, writer, stopChan)
}
