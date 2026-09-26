package services

import (
	"servidorMetadatos/capaAccesoADatos/repository"
	"servidorMetadatos/capaServices/dto"
)

type TipoAudioService struct {
	repository *repository.TipoAudioRepository
}

func NewTipoAudioService(repository *repository.TipoAudioRepository) *TipoAudioService {
	return &TipoAudioService{repository: repository}
}

func (s *TipoAudioService) ListarTipos() dto.RespuestaTiposAudioDTO {
	tipos := s.repository.ListarTipos()
	return dto.RespuestaTiposAudioDTO{
		Tipos:   tipos,
		Codigo:  200,
		Mensaje: "Tipos de audio obtenidos correctamente",
	}
}
