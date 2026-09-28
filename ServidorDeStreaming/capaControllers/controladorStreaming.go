package capacontrollers

import (
	capaService "servidorStreaming/capaServices"
	pb "servidorStreaming/serviciosAudio"
)

type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
}

// Implementación del procedimiento remoto
func (s *ControladorServidor) AudioStream(req *pb.AudioRequest, stream pb.AudioService_AudioStreamServer) error {

	// Delegar la lógica de lectura y envío de chunks (fragmentos) a la fachada.
	return capaService.EnviarFragmentosAudio(req.Filename, stream)

}
