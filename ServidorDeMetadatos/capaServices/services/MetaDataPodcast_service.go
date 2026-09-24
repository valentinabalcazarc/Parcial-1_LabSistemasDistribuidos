package services

import (
	"servidorMetadatos/capaAccesoADatos/repository"
	"servidorMetadatos/capaServices/dto"
)

// MetadataAudiolibroService es la fachada (facade) que expone al controlador
// las operaciones de negocio, ocultando el acceso al repositorio y la
// conversión entre Entity y DTO.
type MetadataPodcastService struct {
	repository *repository.MetadataPodcastRepository
}

func NewMetadataPodcastService(repository *repository.MetadataPodcastRepository) *MetadataPodcastService {
	return &MetadataPodcastService{repository: repository}
}

// ConsultarPodcast recibe un título, busca el Entity en el repositorio y lo
// convierte a RespuestaMetadataPodcastDTO con el código y mensaje según el
// resultado de la búsqueda.
func (this *MetadataPodcastService) ConsultarPodcast(id int) dto.RespuestaMetadataPodcastDTO {
	var respuesta dto.RespuestaMetadataPodcastDTO

	podcast, encontrado := this.repository.BuscarPodcast(id)

	if encontrado {
		var podcastDTO dto.MetadataPodcastDTO
		podcastDTO.ID = podcast.GetId()
		podcastDTO.Tipo = podcast.GetTipo()
		podcastDTO.Nombre = podcast.GetNombre()
		podcastDTO.TituloEpisodio = podcast.GetTituloEpisodio()
		podcastDTO.NumeroTemporada = podcast.GetNumeroTemporada()
		podcastDTO.NotasShow = podcast.GetNotasShow()
		podcastDTO.ClasificacionContenido = podcast.GetClasificacionContenido()

		respuesta.ObjPodcast = podcastDTO
		respuesta.Codigo = 200
		respuesta.Mensaje = "Métadata del podcast encontrada"
	} else {
		respuesta.Codigo = 400
		respuesta.Mensaje = "La métadata del podcast no se encontró"
	}
	return respuesta
}
