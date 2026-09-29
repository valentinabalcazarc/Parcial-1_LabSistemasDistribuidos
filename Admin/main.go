/**
 * @file main.go
 * @brief Punto de entrada principal de la aplicación Admin.
 *
 * Este archivo contiene la interfaz de consola para el rol de administrador,
 * que permite capturar el ID de un audio y su ruta local para subirlo
 * al servidor de audios.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * @brief Función principal del programa.
 *
 * Inicia un menú interactivo en consola, solicitando al administrador
 * un ID para el audio (añadiendo la extensión .mp3 si es necesario) 
 * y la ruta local del archivo. Posteriormente invoca a la función 
 * enviarAudioAlServidor para la transferencia.
 */
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(" ----- Rol: Admin -----")
	fmt.Println("Ingrese el id del audio a guardar:")
	scanner.Scan()
	nombre := strings.TrimSpace(scanner.Text())

	// Si el usuario no escribió .mp3 al final, se lo agregamos automáticamente
	if !strings.HasSuffix(nombre, ".mp3") {
		nombre = nombre + ".mp3"
	}

	fmt.Println("Ingrese la ruta del archivo (.mp3):")
	scanner.Scan()
	ruta := strings.TrimSpace(scanner.Text())

	// La URL fija del servidor de audios corriendo en el puerto 8082
	urlServidor := "http://localhost:8082/audio/upload"

	fmt.Println("\nEnviando archivo al servidor...")
	err := enviarAudioAlServidor(urlServidor, ruta, nombre)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
