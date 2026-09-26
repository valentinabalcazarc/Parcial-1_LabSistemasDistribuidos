package dto

import "servidorMetadatos/capaAccesoADatos/entity"

// MetadataAudiolibroDTO es el objeto de transferencia de datos.
type MetadataAudiolibroDTO struct {
	ID          int              `json:"id"`
	Tipo        entity.TipoAudio `json:"tipo"`
	TituloLibro string           `json:"titulo_libro"`
	Autor       string           `json:"autor"`
	Narrador    string           `json:"narrador"`
	Editorial   string           `json:"editorial"`
	ISBN        int              `json:"isbn"`
	Capitulo    int              `json:"capitulo"`
}

type RespuestaMetadataAudiolibroDTO struct {
	ObjAudiolibro MetadataAudiolibroDTO `json:"objAudiolibro"`
	Codigo        int                   `json:"codigo"`
	Mensaje       string                `json:"mensaje"`
}

type RespuestaListaAudiolibroDTO struct {
	Audiolibros []MetadataAudiolibroDTO `json:"audiolibros"`
	Codigo      int                     `json:"codigo"`
	Mensaje     string                  `json:"mensaje"`
}
