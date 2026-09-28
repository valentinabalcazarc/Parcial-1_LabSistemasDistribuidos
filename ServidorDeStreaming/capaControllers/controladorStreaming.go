/**
 * @file controladorStreaming.go
 * @brief Controlador del servidor gRPC para el servicio de streaming.
 */
package capacontrollers

import (
	capaService "servidorStreaming/capaServices"
	pb "servidorStreaming/serviciosAudio"
)

/**
 * @struct ControladorServidor
 * @brief Estructura que implementa la interfaz generada por gRPC para el servicio de audio.
 */
type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
}

/**
 * @brief Procedimiento remoto para realizar el streaming de audio.
 * @details Delega la lógica de lectura y envío de fragmentos (chunks) a la capa de servicios.
 * @param req Petición gRPC que contiene el nombre del archivo de audio.
 * @param stream Stream del servidor gRPC para enviar los fragmentos de audio.
 * @return error Error en caso de fallo durante el streaming.
 */
func (s *ControladorServidor) AudioStream(req *pb.AudioRequest, stream pb.AudioService_AudioStreamServer) error {

	// Delegar la lógica de lectura y envío de chunks (fragmentos) a la fachada.
	return capaService.EnviarFragmentosAudio(req.Filename, stream)

}
