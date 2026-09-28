package clientestreaming

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	pb "servidorStreaming/serviciosAudio"
)

func DecodificarReproducir(reader io.Reader, stopChan chan struct{}, doneChan chan struct{}) {
	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		log.Printf("\n⚠️ Error decodificando el stream MP3: %v", err)
		close(doneChan)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(doneChan)
	})))

	select {
	case <-doneChan:
		fmt.Println("\n✅ Audio finalizado.")
	case <-stopChan:
		speaker.Clear()
		fmt.Println("\n⏹️ Reproducción detenida por el usuario.")
	}
}

func RecibirAudio(
	stream pb.AudioService_AudioStreamClient,
	writer *io.PipeWriter,
	stopChan chan struct{}) {

	noFragmento := 0
	for {
		select {
		case <-stopChan:
			writer.Close()
			return
		default:
			fragmento, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("\n📥 Audio recibido completo desde el servidor.")
				writer.Close()
				return
			}
			if err != nil {
				// Si fue cancelado o cortado por el usuario no imprimimos error como falla
				select {
				case <-stopChan:
					writer.Close()
					return
				default:
					log.Printf("\n⚠️ Stream interrumpido: %v", err)
					writer.Close()
					return
				}
			}

			noFragmento++
			fmt.Printf("\r🎵 Fragmento #%d recibido (%d bytes) reproduciendo...", noFragmento, len(fragmento.Data))

			if _, err := writer.Write(fragmento.Data); err != nil {
				writer.Close()
				return
			}
		}
	}
}
