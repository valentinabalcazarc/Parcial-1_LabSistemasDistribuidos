package clientestreaming

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	pb "servidorStreaming/serviciosAudio"
)

func DecodificarReproducir(reader io.Reader, stopChan chan struct{}, doneChan chan struct{}) {
	defer close(doneChan)

	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		log.Printf("\n⚠️ Error decodificando el stream MP3: %v", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	playingDone := make(chan struct{})
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(playingDone)
	})))

	select {
	case <-playingDone:
	case <-stopChan:
		speaker.Clear()
	}
}

func RecibirAudio(
	stream pb.AudioService_AudioStreamClient,
	writer *io.PipeWriter,
	stopChan chan struct{}) {

	defer writer.Close()
	noFragmento := 0
	for {
		select {
		case <-stopChan:
			return
		default:
			fragmento, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("\n📥 Audio recibido completo desde el servidor.")
				return
			}
			if err != nil {
				select {
				case <-stopChan:
					return
				default:
					log.Printf("\n⚠️ Stream interrumpido: %v", err)
					return
				}
			}

			noFragmento++
			fmt.Printf("\r🎵 Fragmento #%d recibido (%d bytes) reproduciendo...", noFragmento, len(fragmento.Data))

			if _, err := writer.Write(fragmento.Data); err != nil {
				return
			}
		}
	}
}
