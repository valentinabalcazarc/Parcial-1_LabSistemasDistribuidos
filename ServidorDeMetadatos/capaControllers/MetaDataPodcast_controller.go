package capacontrollers

import (
	"net/http"
	"strconv"

	service "servidorMetadatos/capaServices/services"

	"github.com/gin-gonic/gin"
)

// MetadataPodcastController expone los servicios REST de Podcast usando Gin.
type MetadataPodcastController struct {
	service *service.MetadataPodcastService
}

func NewMetadataPodcastController(service *service.MetadataPodcastService) *MetadataPodcastController {
	return &MetadataPodcastController{service: service}
}

// ConsultarPodcast - GET /podcast/:id
func (this *MetadataPodcastController) ConsultarPodcast(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"codigo":  400,
			"mensaje": "El ID debe ser un número entero válido",
		})
		return
	}

	respuesta := this.service.ConsultarPodcast(id)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// ListarPodcasts - GET /podcasts
func (this *MetadataPodcastController) ListarPodcasts(ctx *gin.Context) {
	respuesta := this.service.ListarPodcasts()
	ctx.JSON(respuesta.Codigo, respuesta)
}
