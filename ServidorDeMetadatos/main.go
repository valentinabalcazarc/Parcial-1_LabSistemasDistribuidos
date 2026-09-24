package main

import (
	repository "servidorMetadatos/ServidorDeMetadatos/capaAccesoADatos/repository"
	controller "servidorMetadatos/ServidorDeMetadatos/capaControllers"
	service "servidorMetadatos/ServidorDeMetadatos/capaServices"

	"github.com/gin-gonic/gin"
)

func main() {
	// AUDIOLIBRO
	audiolibroRepository := repository.NewMetadataAudiolibroRepository()
	audiolibroService := service.NewMetaDataAudiolibroService(audiolibroRepository)
	audiolibroController := controller.NewMetadataAudiolibroController(audiolibroService)

	// MUSICA
	musicaRepository := repository.NewMetadataMusicaRepository()
	musicaService := service.NewMetaDataMusicaService(musicaRepository)
	musicaController := controller.NewMetadataMusicaController(musicaService)

	// PODCAST
	podcastRepository := repository.NewMetadataPodcastRepository()
	podcastService := service.NewMetaDataPodcastService(podcastRepository)
	podcastController := controller.NewMetadataPodcastController(podcastService)

	// RUIDO BLANCO
	ruidoBlancoRepository := repository.NewMetadataRuidoBlancoRepository()
	ruidoBlancoService := service.NewMetaDataRuidoBlancoService(ruidoBlancoRepository)
	ruidoBlancoController := controller.NewMetadataRuidoBlancoController(ruidoBlancoService)

	// ROUTER GIN
	router := gin.Default()

	// Rutas de consulta por ID
	router.GET("/audiolibro/:id", audiolibroController.ConsultarAudiolibro)
	router.GET("/musica/:id", musicaController.ConsultarMusica)
	router.GET("/podcast/:id", podcastController.ConsultarPodcast)
	router.GET("/ruidoblanco/:id", ruidoBlancoController.ConsultarRuidoBlanco)

	// Iniciar servidor en el puerto 8081
	router.Run(":8081")
}
