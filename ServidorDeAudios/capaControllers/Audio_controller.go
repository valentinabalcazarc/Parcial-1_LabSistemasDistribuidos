package capacontrollers

import (
	"fmt"
	"io"
	"net/http"
	dtos "servidorAudios/capaServices/DTOs"
	service "servidorAudios/capaServices/services"
)

type AudioController struct {
	service *service.AudioService
}

func NewAudioController(service *service.AudioService) *AudioController {
	return &AudioController{service: service}
}

func (thisC *AudioController) GuardarAudio(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(50 << 20)
	file, _, err := r.FormFile("Archivo")
	if err != nil {
		http.Error(w, "Error leyendo el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	dto := dtos.AudioDTO{
		NombreArchivo: r.FormValue("nombre_archivo"),
	}

	err = thisC.service.GuardarAudio(dto, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al guardar: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Audio guardado exitosamente"))
}

func (thisC *AudioController) ObtenerAudio(w http.ResponseWriter, r *http.Request) {
	nombreArchivo := r.URL.Query().Get("nombre")
	if nombreArchivo == "" {
		http.Error(w, "Parámetro 'nombre' requerido", http.StatusBadRequest)
		return
	}

	bytesAudio, err := thisC.service.ObtenerAudio(nombreArchivo)
	if err != nil {
		http.Error(w, "Archivo no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Write(bytesAudio)
}
