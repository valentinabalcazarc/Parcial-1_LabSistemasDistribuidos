package clientestreaming

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "servidorStreaming/serviciosAudio"
)

const urlStreaming = "localhost:8083"

func ReproducirAudio(scanner *bufio.Scanner, filename string) {
	fmt.Printf("\n📡 Conectando al Servidor de Streaming (%s)...\n", urlStreaming)

	conn, err := grpc.Dial(urlStreaming, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("❌ Error al conectar con servidor gRPC: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewAudioServiceClient(conn)

	ctx, cancelTimeout := context.WithTimeout(context.Background(), 5*time.Minute)
	ctx, cancelStream := context.WithCancel(ctx)
	defer cancelTimeout()
	defer cancelStream()

	stream, err := client.AudioStream(ctx, &pb.AudioRequest{Filename: filename})
	if err != nil {
		fmt.Printf("❌ Error al iniciar el stream: %v\n", err)
		return
	}

	reader, writer := io.Pipe()
	stopChan := make(chan struct{})
	doneChan := make(chan struct{})
	enterPressedChan := make(chan struct{})

	// Goroutine que escucha la tecla ENTER
	go func() {
		if scanner.Scan() {
			close(enterPressedChan)
		}
	}()

	fmt.Printf("▶️  Reproduciendo %s en tiempo real vía gRPC...\n", filename)
	fmt.Println("👉 Presione ENTER en cualquier momento para detener la reproducción y volver al menú.")

	// Decodificar y reproducir sonido en goroutine
	go DecodificarReproducir(reader, stopChan, doneChan)

	// Recibir fragmentos de gRPC en goroutine
	go RecibirAudio(stream, writer, stopChan)

	// Monitorear eventos: el usuario presiona ENTER o el audio termina
	select {
	case <-enterPressedChan:
		// El usuario presionó ENTER durante la reproducción
		close(stopChan)
		cancelStream()
		<-doneChan // Esperar a que la decodificación limpie el speaker
		fmt.Println("\n⏹️ Reproducción detenida por el usuario.")

	case <-doneChan:
		// El audio terminó por sí solo (o por error)
		close(stopChan)
		cancelStream()
		fmt.Println("\n✅ Audio finalizado. Presione ENTER para volver al menú...")
		<-enterPressedChan // Esperar a que el usuario presione ENTER para liberar el scanner
	}
}
