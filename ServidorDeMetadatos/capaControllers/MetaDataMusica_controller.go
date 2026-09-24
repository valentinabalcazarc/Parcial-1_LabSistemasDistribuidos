package capacontrollers

import (
	"net/http"
	"strconv"

	service "servidorMetadatos/capaServices"

	"github.com/gin-gonic/gin"
)

// MetadataMusicaController expone los servicios REST de Musica usando Gin.
type MetadataMusicaController struct {
	service *service.MetadataMusicaService
}

func NewMetadataMusicaController(service *service.MetadataMusicaService) *MetadataMusicaController {
	return &MetadataMusicaController{service: service}
}

// ConsultarMusica - GET /musica/:id
func (this *MetadataMusicaController) ConsultarMusica(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"codigo":  400,
			"mensaje": "El ID debe ser un número entero válido",
		})
		return
	}

	respuesta := this.service.ConsultarMusica(id)
	ctx.JSON(respuesta.Codigo, respuesta)
}
