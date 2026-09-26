package dto

import "servidorMetadatos/capaAccesoADatos/entity"

// MetadataPodcastDTO es el objeto de transferencia de datos.
type MetadataPodcastDTO struct {
	ID                     int              `json:"id"`
	Tipo                   entity.TipoAudio `json:"tipo"`
	Nombre                 string           `json:"nombre"`
	TituloEpisodio         string           `json:"titulo_episodio"`
	NumeroTemporada        int              `json:"numero_temporada"`
	NotasShow              string           `json:"notas_show"`
	ClasificacionContenido string           `json:"clasificacion_contenido"`
}

type RespuestaMetadataPodcastDTO struct {
	ObjPodcast MetadataPodcastDTO `json:"objPodcast"`
	Codigo     int                `json:"codigo"`
	Mensaje    string             `json:"mensaje"`
}

type RespuestaListaPodcastDTO struct {
	Podcasts []MetadataPodcastDTO `json:"podcasts"`
	Codigo   int                  `json:"codigo"`
	Mensaje  string               `json:"mensaje"`
}
