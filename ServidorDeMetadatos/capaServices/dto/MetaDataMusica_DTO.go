package dto

import "servidorMetadatos/capaAccesoADatos/entity"

// MetadataMusicaDTO es el objeto de transferencia de datos.
type MetadataMusicaDTO struct {
	ID                int              `json:"id"`
	Tipo              entity.TipoAudio `json:"tipo"`
	ArtistaPrincipal  string           `json:"artista_principal"`
	Album             string           `json:"album"`
	Genero            string           `json:"genero"`
	TituloCancion     string           `json:"titulo_cancion"`
	SelloDiscografico string           `json:"sello_discografico"`
	AnioLanzamiento   string           `json:"anio_lanzamiento"`
}

type RespuestaMetadataMusicaDTO struct {
	ObjMusica MetadataMusicaDTO `json:"objMusica"`
	Codigo    int               `json:"codigo"`
	Mensaje   string            `json:"mensaje"`
}

type RespuestaListaMusicaDTO struct {
	Musica  []MetadataMusicaDTO `json:"musica"`
	Codigo  int                 `json:"codigo"`
	Mensaje string              `json:"mensaje"`
}
