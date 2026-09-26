package capacontrollers

import (
	service "servidorMetadatos/capaServices/services"

	"github.com/gin-gonic/gin"
)

type TipoAudioController struct {
	service *service.TipoAudioService
}

func NewTipoAudioController(service *service.TipoAudioService) *TipoAudioController {
	return &TipoAudioController{service: service}
}

// ListarTipos - GET /tipos-audio
func (c *TipoAudioController) ListarTipos(ctx *gin.Context) {
	respuesta := c.service.ListarTipos()
	ctx.JSON(respuesta.Codigo, respuesta)
}
