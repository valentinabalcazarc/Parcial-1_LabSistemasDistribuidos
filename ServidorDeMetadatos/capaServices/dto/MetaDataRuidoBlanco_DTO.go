package dto

import "servidorMetadatos/capaAccesoADatos/entity"

// MetadataRuidoBlancoDTO es el objeto de transferencia de datos.
type MetadataRuidoBlancoDTO struct {
	ID                  int              `json:"id"`
	Tipo                entity.TipoAudio `json:"tipo"`
	TipoSonido          string           `json:"tipo_sonido"`
	FuenteAudio         string           `json:"fuente_audio"`
	UsoSugerido         string           `json:"uso_sugerido"`
	Proveedor           string           `json:"proveedor"`
	Duracion            int              `json:"duracion"`
	FrecuenciaDominante string           `json:"frecuencia_dominante"`
}

type RespuestaMetadataRuidoBlancoDTO struct {
	ObjRuidoBlanco MetadataRuidoBlancoDTO `json:"objRuidoBlanco"`
	Codigo         int                    `json:"codigo"`
	Mensaje        string                 `json:"mensaje"`
}

type RespuestaListaRuidoBlancoDTO struct {
	RuidosBlancos []MetadataRuidoBlancoDTO `json:"ruidosBlancos"`
	Codigo        int                      `json:"codigo"`
	Mensaje       string                   `json:"mensaje"`
}
