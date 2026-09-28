package capaaccesoadatos

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type RepositorioStreaming struct{}

var (
	instancia *RepositorioStreaming
	once      sync.Once
)

// aplica singleton
func GetRepositorioStreaming() *RepositorioStreaming {
	once.Do(func() {
		instancia = &RepositorioStreaming{}
	})
	return instancia
}

// AbrirArchivo abre el archivo físico .mp3 desde la carpeta 'audios' y lo retorna como *os.File
func (r *RepositorioStreaming) AbrirArchivo(nombreArchivo string) (*os.File, error) {
	// Se busca en la carpeta audios del proyecto
	ruta := filepath.Join("audios", nombreArchivo)

	// Si no existe localmente en 'audios/', se busca en la carpeta del ServidorDeAudios por compatibilidad
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		ruta = filepath.Join("..", "ServidorDeAudios", "audios", nombreArchivo)
	}

	file, err := os.Open(ruta)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el archivo de audio %s: %v", nombreArchivo, err)
	}

	return file, nil
}
