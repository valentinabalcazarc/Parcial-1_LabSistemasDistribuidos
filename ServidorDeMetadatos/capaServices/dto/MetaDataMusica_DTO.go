package dto

// MetadataMusicaDTO es el objeto de transferencia de datos que se recibe al
// registrar un audio y que se devuelve dentro de la respuesta de consulta.
type MetadataMusicaDTO struct {
	ID                int    `json:"id"`
	Tipo              string `json:"tipo"`
	ArtistaPrincipal  string `json:"artista_principal"`
	Album             string `json:"album"`
	Genero            string `json:"genero_"`
	TituloCancion     string `json:"titulo_cancion"`
	SelloDiscografico string `json:"sello_discografico"`
	AnioLanzamiento   string `json:"anio_lanzamiento"`
}

// RespuestaMetadataMusicaDTO es el DTO de respuesta para la consulta de un
// Musica: incluye el Musica encontrado (si aplica), un código de resultado y
// un mensaje descriptivo.
type RespuestaMetadataMusicaDTO struct {
	ObjMusica MetadataMusicaDTO `json:"objMusica"`
	Codigo    int               `json:"codigo"`
	Mensaje   string            `json:"mensaje"`
}
