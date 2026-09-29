package clientemetadatos

import "fmt"

/**
 * @brief Estructura que representa un tipo de audio.
 */
type TipoAudio struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

/**
 * @brief Estructura que representa la respuesta del servidor para tipos de audio.
 */
type RespuestaTiposAudio struct {
	Tipos   []TipoAudio `json:"tipos"`
	Codigo  int         `json:"codigo"`
	Mensaje string      `json:"mensaje"`
}

/**
 * @brief Estructura que representa un ítem de audio de forma genérica.
 */
type ItemAudio struct {
	ID     int       `json:"id"`
	Tipo   TipoAudio `json:"tipo"`
	Titulo string    `json:"titulo_cancion"`
	Nombre string    `json:"nombre"`
	Libro  string    `json:"titulo_libro"`
	Sonido string    `json:"tipo_sonido"`
}

/**
 * @brief Obtiene el nombre o título del ítem de audio.
 * 
 * Verifica los diferentes campos donde podría estar almacenado
 * el nombre según el tipo real del audio.
 * 
 * @return Cadena con el nombre o título del audio.
 */
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
