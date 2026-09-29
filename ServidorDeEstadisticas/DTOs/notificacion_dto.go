/**
 * @file notificacion_dto.go
 * @brief Definición del DTO para las notificaciones de reproducción.
 */
package dtos

/**
 * @struct NotificacionReproduccion
 * @brief Estructura de transferencia de datos (DTO) que representa una notificación de reproducción.
 * 
 * Esta estructura contiene la información necesaria para notificar al servidor
 * de estadísticas sobre la reproducción de un determinado contenido de audio.
 */
type NotificacionReproduccion struct {
	Titulo    string `json:"titulo"`               /**< @brief Título del audio reproducido. */
	TipoAudio string `json:"tipo_audio,omitempty"` /**< @brief Tipo de audio (ej. Canción, Podcast), opcional. */
	FechaHora string `json:"fecha_hora,omitempty"` /**< @brief Fecha y hora de la reproducción, opcional. */
	Mensaje   string `json:"mensaje"`              /**< @brief Mensaje descriptivo de la notificación. */
}
