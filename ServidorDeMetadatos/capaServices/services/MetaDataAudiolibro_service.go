package services

import (
	"servidorMetadatos/capaAccesoADatos/repository"
	"servidorMetadatos/capaServices/dto"
)

// MetadataAudiolibroService es la fachada (facade) que expone al controlador
// las operaciones de negocio, ocultando el acceso al repositorio y la
// conversión entre Entity y DTO.
type MetadataAudiolibroService struct {
	repository *repository.MetadataAudiolibroRepository
}

func NewMetadataAudiolibroService(repository *repository.MetadataAudiolibroRepository) *MetadataAudiolibroService {
	return &MetadataAudiolibroService{repository: repository}
}

// ConsultarAudiolibro recibe un título, busca el Entity en el repositorio y lo
// convierte a RespuestaMetadataAudiolibroDTO con el código y mensaje según el
// resultado de la búsqueda.
func (this *MetadataAudiolibroService) ConsultarAudiolibro(id int) dto.RespuestaMetadataAudiolibroDTO {
	var respuesta dto.RespuestaMetadataAudiolibroDTO

	audiolibro, encontrado := this.repository.BuscarAudiolibro(id)

	if encontrado {
		var audiolibroDTO dto.MetadataAudiolibroDTO
		audiolibroDTO.ID = audiolibro.GetId()
		audiolibroDTO.Tipo = audiolibro.GetTipo()
		audiolibroDTO.TituloLibro = audiolibro.GetTituloLibro()
		audiolibroDTO.Autor = audiolibro.GetAutor()
		audiolibroDTO.Narrador = audiolibro.GetNarrador()
		audiolibroDTO.Editorial = audiolibro.GetEditorial()
		audiolibroDTO.ISBN = audiolibro.GetISBN()
		audiolibroDTO.Capitulo = audiolibro.GetCapitulo()

		respuesta.ObjAudiolibro = audiolibroDTO
		respuesta.Codigo = 200
		respuesta.Mensaje = "Métadata del audiolibro encontrada"
	} else {
		respuesta.Codigo = 400
		respuesta.Mensaje = "La métadata del audiolibro no se encontró"
	}
	return respuesta
}

func (this *MetadataAudiolibroService) ListarAudiolibros() dto.RespuestaListaAudiolibroDTO {
	lista := this.repository.ListarAudiolibros()
	var dtos []dto.MetadataAudiolibroDTO

	for _, audiolibro := range lista {
		var d dto.MetadataAudiolibroDTO
		d.ID = audiolibro.GetId()
		d.Tipo = audiolibro.GetTipo()
		d.TituloLibro = audiolibro.GetTituloLibro()
		d.Autor = audiolibro.GetAutor()
		d.Narrador = audiolibro.GetNarrador()
		d.Editorial = audiolibro.GetEditorial()
		d.ISBN = audiolibro.GetISBN()
		d.Capitulo = audiolibro.GetCapitulo()

		dtos = append(dtos, d)
	}

	return dto.RespuestaListaAudiolibroDTO{
		Audiolibros: dtos,
		Codigo:      200,
		Mensaje:     "Lista de audiolibros obtenida correctamente",
	}
}
