package capacontrollers

import (
	"net/http"
	"strconv"

	service "servidorMetadatos/capaServices"

	"github.com/gin-gonic/gin"
)

// MetadataAudiolibroController expone los servicios REST de audiolibro usando Gin.
type MetadataAudiolibroController struct {
	service *service.MetadataAudiolibroService
}

func NewMetadataAudiolibroController(service *service.MetadataAudiolibroService) *MetadataAudiolibroController {
	return &MetadataAudiolibroController{service: service}
}

// ConsultarAudiolibro - GET /audiolibro/:id
func (this *MetadataAudiolibroController) ConsultarAudiolibro(ctx *gin.Context) {
	idParam := ctx.Param("id")

	// Convertir el id de string a int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"codigo":  400,
			"mensaje": "El ID debe ser un número entero válido",
		})
		return
	}

	respuesta := this.service.ConsultarAudiolibro(id)
	ctx.JSON(respuesta.Codigo, respuesta)
}
