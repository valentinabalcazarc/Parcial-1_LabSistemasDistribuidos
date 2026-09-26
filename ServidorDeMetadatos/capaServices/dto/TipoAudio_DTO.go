package dto

import "servidorMetadatos/capaAccesoADatos/entity"

type RespuestaTiposAudioDTO struct {
	Tipos   []entity.TipoAudio `json:"tipos"`
	Codigo  int                `json:"codigo"`
	Mensaje string             `json:"mensaje"`
}
