package dto

// MetadataRuidoBlancoDTO es el objeto de transferencia de datos que se recibe
// al registrar un RuidoBlanco y que se devuelve dentro de la respuesta de consulta.
type MetadataRuidoBlancoDTO struct {
	ID                  int    `json:"id"`
	Tipo                string `json:"tipo"`
	TipoSonido          string `json:"tipo_sonido"`
	FuenteAudio         string `json:"fuente_audio"`
	UsoSugerido         string `json:"uso_sugerido"`
	Proveedor           string `json:"proveedor"`
	Duracion            int    `json:"duracion"`
	FrecuenciaDominante string `json:"frecuencia_dominante"`
}

// RespuestaMetadataRuidoBlancoDTO es el DTO de respuesta para la consulta de un
// RuidoBlanco: incluye el RuidoBlanco encontrado (si aplica), un código de resultado y
// un mensaje descriptivo.
type RespuestaMetadataRuidoBlancoDTO struct {
	ObjRuidoBlanco MetadataRuidoBlancoDTO `json:"objRuidoBlanco"`
	Codigo         int                    `json:"codigo"`
	Mensaje        string                 `json:"mensaje"`
}
