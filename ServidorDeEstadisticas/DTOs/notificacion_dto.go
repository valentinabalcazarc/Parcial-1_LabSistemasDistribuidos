package dtos

type NotificacionReproduccion struct {
	Titulo    string `json:"titulo"`
	TipoAudio string `json:"tipo_audio,omitempty"`
	FechaHora string `json:"fecha_hora,omitempty"`
	Mensaje   string `json:"mensaje"`
}
