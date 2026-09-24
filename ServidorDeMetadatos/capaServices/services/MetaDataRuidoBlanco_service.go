package services

import (
	"servidorMetadatos/capaAccesoADatos/repository"
	"servidorMetadatos/capaServices/dto"
)

// MetadataAudiolibroService es la fachada (facade) que expone al controlador
// las operaciones de negocio, ocultando el acceso al repositorio y la
// conversión entre Entity y DTO.
type MetadataRuidoBlancoService struct {
	repository *repository.MetadataRuidoBlancoRepository
}

func NewMetadataRuidoBlancoService(repository *repository.MetadataRuidoBlancoRepository) *MetadataRuidoBlancoService {
	return &MetadataRuidoBlancoService{repository: repository}
}

// ConsultarRuidoBlanco recibe un título, busca el Entity en el repositorio y lo
// convierte a RespuestaMetadataRuidoBlancoDTO con el código y mensaje según el
// resultado de la búsqueda.
func (this *MetadataRuidoBlancoService) ConsultarRuidoBlanco(id int) dto.RespuestaMetadataRuidoBlancoDTO {
	var respuesta dto.RespuestaMetadataRuidoBlancoDTO

	ruidoBlanco, encontrado := this.repository.BuscarRuidoBlanco(id)

	if encontrado {
		var ruidoBlancoDTO dto.MetadataRuidoBlancoDTO
		ruidoBlancoDTO.ID = ruidoBlanco.GetId()
		ruidoBlancoDTO.Tipo = ruidoBlanco.GetTipo()
		ruidoBlancoDTO.TipoSonido = ruidoBlanco.GetTipoSonido()
		ruidoBlancoDTO.FuenteAudio = ruidoBlanco.GetFuenteAudio()
		ruidoBlancoDTO.UsoSugerido = ruidoBlanco.GetUsoSugerido()
		ruidoBlancoDTO.Proveedor = ruidoBlanco.GetProveedor()
		ruidoBlancoDTO.Duracion = ruidoBlanco.GetDuracion()
		ruidoBlancoDTO.FrecuenciaDominante = ruidoBlanco.GetFrecuenciaDominante()

		respuesta.ObjRuidoBlanco = ruidoBlancoDTO
		respuesta.Codigo = 200
		respuesta.Mensaje = "Métadata del ruidoBlanco encontrada"
	} else {
		respuesta.Codigo = 400
		respuesta.Mensaje = "La métadata del ruidoBlanco no se encontró"
	}
	return respuesta
}
