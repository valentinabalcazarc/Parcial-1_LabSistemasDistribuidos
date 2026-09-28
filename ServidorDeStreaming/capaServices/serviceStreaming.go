package capaservices

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	capaAccesoDatos "servidorStreaming/capaAccesoADatos"
	componenteconexioncola "servidorStreaming/componenteConexionCola"
	pb "servidorStreaming/serviciosAudio"
)

type StreamingService struct {
	repo         *capaAccesoDatos.RepositorioStreaming
	conexionCola *componenteconexioncola.RabbitPublisher
}

func NewStreamingService() *StreamingService {
	fmt.Println("Inicializando service de streaming")
	repo := capaAccesoDatos.GetRepositorioStreaming()
	conexionCola, err := componenteconexioncola.NewRabbitPublisher()

	if err != nil {
		fmt.Println("Error al conectar con RabbitMQ: ", err)
		conexionCola = nil
	}

	return &StreamingService{
		repo:         repo,
		conexionCola: conexionCola,
	}
}

// abrirArchivo actua como fachada para obtener el archivo de la cancion usando la capa de acceso a datos.
func abrirArchivo(titulo string) (*os.File, error) {
	log.Printf("GetAudioFile llamado con titulo=%s", titulo)
	return capaAccesoDatos.GetRepositorioStreaming().AbrirArchivo(titulo)
}

// EnviarFragmentosAudio lee el archivo de la canción en chunks y los envía al stream gRPC.
func EnviarFragmentosAudio(titulo string, stream pb.AudioService_AudioStreamServer) error {
	log.Printf("Enviando fragmentos de audio para titulo=%s", titulo)
	file, err := abrirArchivo(titulo)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	// Notificar inicio de reproducción a RabbitMQ si la conexión existe
	pub, err := componenteconexioncola.NewRabbitPublisher()
	if err == nil {
		defer pub.Cerrar()
		msg := componenteconexioncola.NotificacionReproduccion{
			Titulo:    titulo,
			FechaHora: time.Now().Format("2006-01-02 15:04:05"),
			Mensaje:   fmt.Sprintf("Reproduciendo audio %s por streaming", titulo),
		}
		pub.PublicarNotificacion(msg)
	} else {
		log.Printf("Advertencia: No se pudo conectar a RabbitMQ para enviar notificación: %v", err)
	}

	buf := make([]byte, 32*1024) // 32KB por chunk (fragmento)
	chunkNum := 0

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			log.Println("Audio enviado completo")
			break
		}
		if err != nil {
			return fmt.Errorf("error leyendo archivo: %w", err)
		}

		chunkNum++

		if n > 0 {
			objChunk := &pb.AudioChunk{Data: buf[:n]}
			if err := stream.Send(objChunk); err != nil {
				return fmt.Errorf("Error enviando chunk #%d: %w", chunkNum, err)
			}
			log.Printf("Chunk #%d enviado (%d bytes)\n", chunkNum, n)
		}
	}

	return nil
}
