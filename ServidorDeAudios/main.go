/**
 * @file main.go
 * @brief Punto de entrada principal para el Servidor de Audios.
 */
package main

import (
	"fmt"
	"net/http"
	controller "servidorAudios/capaControllers"
	service "servidorAudios/capaServices/services"
)

/**
 * @brief Función principal que inicializa el servidor web y los controladores.
 * 
 * Inicia el servicio de almacenamiento de audios configurando las rutas y escuchando
 * en el puerto 8082.
 */
func main() {
	audioService := service.NewAudioService()
	audioController := controller.NewAudioController(audioService)

	// Registrar endpoints
	http.HandleFunc("/audio/upload", audioController.GuardarAudio)
	http.HandleFunc("/audio/obtener", audioController.ObtenerAudio)

	fmt.Println("Servicio de almacenamiento escuchando en el puerto 8082...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("Error iniciando el servidor: ", err)
	}
}
