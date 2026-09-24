package dto

// MetadataPodcastDTO es el objeto de transferencia de datos que se recibe al
// registrar un Podcast y que se devuelve dentro de la respuesta de consulta.
type MetadataPodcastDTO struct {
	ID                     int    `json:"id"`
	Tipo                   string `json:"tipo"`
	Nombre                 string `json:"nombre"`
	TituloEpisodio         string `json:"titulo_episodio"`
	NumeroTemporada        int    `json:"numero_temporada"`
	NotasShow              string `json:"notas_show"`
	ClasificacionContenido string `json:"clasificacion_contenido"`
}

// RespuestaMetadataPodcastDTO es el DTO de respuesta para la consulta de un
// Podcast: incluye el Podcast encontrado (si aplica), un código de resultado y
// un mensaje descriptivo.
type RespuestaMetadataPodcastDTO struct {
	ObjPodcast MetadataPodcastDTO `json:"objPodcast"`
	Codigo     int                `json:"codigo"`
	Mensaje    string             `json:"mensaje"`
}
