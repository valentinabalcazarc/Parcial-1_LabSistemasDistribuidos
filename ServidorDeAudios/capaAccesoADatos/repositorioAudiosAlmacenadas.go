package capaaccesoadatos

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type RepositorioAudios struct {
	mu sync.Mutex
}

var (
	instancia *RepositorioAudios
	once      sync.Once
)

// aplica singleton
func GetRepositorioAudios() *RepositorioAudios {
	once.Do(func() {
		instancia = &RepositorioAudios{}
	})
	return instancia
}

// GuardarAudio guarda el archivo físico .mp3 en la carpeta 'audios'
func (r *RepositorioAudios) GuardarAudio(nombreArchivo string, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Crear carpeta si no existe
	if err := os.MkdirAll("audios", os.ModePerm); err != nil {
		return fmt.Errorf("error al crear directorio audios: %v", err)
	}

	filePath := filepath.Join("audios", nombreArchivo)

	// Guardar archivo fisico
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error al guardar el archivo: %v", err)
	}

	return nil
}

// ObtenerAudio lee y devuelve los bytes del archivo físico .mp3
func (r *RepositorioAudios) ObtenerAudio(nombreArchivo string) ([]byte, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	filePath := filepath.Join("audios", nombreArchivo)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo %s: %v", nombreArchivo, err)
	}

	return data, nil
}
