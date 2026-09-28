package dto

// AudioDTO representa los metadatos básicos opcionales recibidos al subir o consultar un audio
type AudioDTO struct {
	NombreArchivo string `json:"nombre_archivo"`
}
