/**
 * @file utilidades.go
 * @brief Funciones de utilidad para el cliente Admin.
 *
 * Contiene la lógica necesaria para leer archivos locales y enviarlos
 * a través de peticiones HTTP en formato multipart/form-data hacia el servidor.
 */
package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

/**
 * @brief Empaqueta un archivo de audio y su identificador en multipart/form-data y realiza una petición POST al servidor.
 *
 * Esta función lee el archivo local, crea los campos requeridos ("Archivo" y "nombre_archivo")
 * y los envía a la URL configurada del servidor.
 *
 * @param urlServidor   La URL completa del endpoint donde se subirá el archivo (ej. http://localhost:8082/audio/upload).
 * @param rutaArchivo   La ruta absoluta o relativa del archivo de audio en el sistema de archivos local.
 * @param nombreArchivo El ID o nombre que se le asignará al archivo en el servidor.
 *
 * @return error Retorna nil si la subida fue exitosa, o un error con detalles si algo falló.
 */
func enviarAudioAlServidor(urlServidor string, rutaArchivo string, nombreArchivo string) error {
	// 1. Abrir el archivo MP3 desde la ruta ingresada
	file, err := os.Open(rutaArchivo)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo en la ruta especificada: %v", err)
	}
	defer file.Close()

	// 2. Preparar el buffer en memoria y el escritor multipart
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 3. Crear el campo del archivo con la clave "Archivo" (tal como lo espera r.FormFile("Archivo"))
	part, err := writer.CreateFormFile("Archivo", filepath.Base(rutaArchivo))
	if err != nil {
		return fmt.Errorf("error al crear el campo 'Archivo': %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("error al copiar el archivo al buffer: %v", err)
	}

	// 4. Crear el campo de texto con la clave "nombre_archivo" (tal como lo espera r.FormValue("nombre_archivo"))
	err = writer.WriteField("nombre_archivo", nombreArchivo)
	if err != nil {
		return fmt.Errorf("error al escribir 'nombre_archivo': %v", err)
	}

	// Cerrar el writer para colocar los delimitadores finales de multipart
	writer.Close()

	// 5. Crear y enviar la petición HTTP POST
	req, err := http.NewRequest(http.MethodPost, urlServidor, body)
	if err != nil {
		return fmt.Errorf("error creando la petición HTTP: %v", err)
	}

	// Es indispensable indicar el Content-Type correcto generado por el writer
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("falló la conexión con el servidor (¿está corriendo en el puerto 8082?): %v", err)
	}
	defer resp.Body.Close()

	// 6. Leer e imprimir la respuesta
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("el servidor respondió con error (%d): %s", resp.StatusCode, string(respBody))
	}

	fmt.Printf("¡Éxito!: %s\n", string(respBody))
	return nil
}
