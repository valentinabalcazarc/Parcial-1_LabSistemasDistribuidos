package services

import (
	"servidorMetadatos/capaAccesoADatos/repository"
	"servidorMetadatos/capaServices/dto"
)

// MetadataAudiolibroService es la fachada (facade) que expone al controlador
// las operaciones de negocio, ocultando el acceso al repositorio y la
// conversión entre Entity y DTO.
type MetadataMusicaService struct {
	repository *repository.MetadataMusicaRepository
}

func NewMetadataMusicaService(repository *repository.MetadataMusicaRepository) *MetadataMusicaService {
	return &MetadataMusicaService{repository: repository}
}

// ConsultarMusica recibe un título, busca el Entity en el repositorio y lo
// convierte a RespuestaMetadataMusicaDTO con el código y mensaje según el
// resultado de la búsqueda.
func (this *MetadataMusicaService) ConsultarMusica(id int) dto.RespuestaMetadataMusicaDTO {
	var respuesta dto.RespuestaMetadataMusicaDTO

	musica, encontrado := this.repository.BuscarMusica(id)

	if encontrado {
		var musicaDTO dto.MetadataMusicaDTO
		musicaDTO.ID = musica.GetId()
		musicaDTO.Tipo = musica.GetTipo()
		musicaDTO.ArtistaPrincipal = musica.GetArtistaPrincipal()
		musicaDTO.Album = musica.GetAlbum()
		musicaDTO.Genero = musica.GetGenero()
		musicaDTO.TituloCancion = musica.GetTituloCancion()
		musicaDTO.SelloDiscografico = musica.GetSelloDiscografico()
		musicaDTO.AnioLanzamiento = musica.GetAnioLanzamiento()

		respuesta.ObjMusica = musicaDTO
		respuesta.Codigo = 200
		respuesta.Mensaje = "Métadata de la musica encontrada"
	} else {
		respuesta.Codigo = 400
		respuesta.Mensaje = "La métadata de la musica no se encontró"
	}
	return respuesta
}

func (this *MetadataMusicaService) ListarMusica() dto.RespuestaListaMusicaDTO {
	lista := this.repository.ListarMusica()
	var dtos []dto.MetadataMusicaDTO

	for _, musica := range lista {
		var d dto.MetadataMusicaDTO
		d.ID = musica.GetId()
		d.Tipo = musica.GetTipo()
		d.ArtistaPrincipal = musica.GetArtistaPrincipal()
		d.Album = musica.GetAlbum()
		d.Genero = musica.GetGenero()
		d.TituloCancion = musica.GetTituloCancion()
		d.SelloDiscografico = musica.GetSelloDiscografico()
		d.AnioLanzamiento = musica.GetAnioLanzamiento()

		dtos = append(dtos, d)
	}

	return dto.RespuestaListaMusicaDTO{
		Musica:  dtos,
		Codigo:  200,
		Mensaje: "Lista de música obtenida correctamente",
	}
}
