package clientemetadatos

import "fmt"

type TipoAudio struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type RespuestaTiposAudio struct {
	Tipos   []TipoAudio `json:"tipos"`
	Codigo  int         `json:"codigo"`
	Mensaje string      `json:"mensaje"`
}

type ItemAudio struct {
	ID     int       `json:"id"`
	Tipo   TipoAudio `json:"tipo"`
	Titulo string    `json:"titulo_cancion"`
	Nombre string    `json:"nombre"`
	Libro  string    `json:"titulo_libro"`
	Sonido string    `json:"tipo_sonido"`
}

func (i ItemAudio) ObtenerNombre() string {
	if i.Titulo != "" {
		return i.Titulo
	}
	if i.Nombre != "" {
		return i.Nombre
	}
	if i.Libro != "" {
		return i.Libro
	}
	if i.Sonido != "" {
		return i.Sonido
	}
	return fmt.Sprintf("Audio #%d", i.ID)
}
