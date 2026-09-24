package dto

// MetadataAudioDTO es el objeto de transferencia de datos que se recibe al
// registrar un audio y que se devuelve dentro de la respuesta de consulta.
type MetadataAudiolibroDTO struct {
	ID          int    `json:"id"`
	Tipo        string `json:"tipo"`
	TituloLibro string `json:"titulo_libro"`
	Autor       string `json:"autor"`
	Narrador    string `json:"narrador"`
	Editorial   string `json:"editorial"`
	ISBN        int    `json:"isbn"`
	Capitulo    int    `json:"capitulo"`
}

// RespuestaMetadataAudiolibroDTO es el DTO de respuesta para la consulta de un
// Audiolibro: incluye el Audiolibro encontrado (si aplica), un código de resultado y
// un mensaje descriptivo.
type RespuestaMetadataAudiolibroDTO struct {
	ObjAudiolibro MetadataAudiolibroDTO `json:"objAudiolibro"`
	Codigo        int                   `json:"codigo"`
	Mensaje       string                `json:"mensaje"`
}
