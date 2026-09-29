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

/**
 * @brief Decodifica un flujo MP3 y lo reproduce a través de la salida de audio.
 * 
 * @param reader Lector de donde provienen los datos MP3
 * @param stopChan Canal para recibir la señal de detención manual de reproducción.
 * @param doneChan Canal que se cierra al finalizar la decodificación de todo el stream.
 */
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

/**
 * @brief Recibe los fragmentos de audio desde el stream gRPC y los escribe en un PipeWriter.
 * 
 * @param stream Cliente del stream de gRPC desde donde llegan los fragmentos de audio.
 * @param writer Escritor del Pipe que envía los datos al decodificador.
 * @param stopChan Canal para interrumpir la recepción en caso de cancelación por el usuario.
 */
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
