package services

import (
	"fmt"
	capaaccesoadatos "servidorAudios/capaAccesoADatos"
	dtos "servidorAudios/capaServices/DTOs"
)

type AudioService struct {
	repo *capaaccesoadatos.RepositorioAudios
}

func NewAudioService() *AudioService {
	fmt.Println("Inicializando service de almacenamiento")
	return &AudioService{
		repo: capaaccesoadatos.GetRepositorioAudios(),
	}
}

func (thisS *AudioService) GuardarAudio(objAudio dtos.AudioDTO, data []byte) error {
	return thisS.repo.GuardarAudio(objAudio.NombreArchivo, data)
}

// ObtenerAudio obtiene los bytes del archivo mp3 por su nombre
func (s *AudioService) ObtenerAudio(nombreArchivo string) ([]byte, error) {
	return s.repo.ObtenerAudio(nombreArchivo)
}
