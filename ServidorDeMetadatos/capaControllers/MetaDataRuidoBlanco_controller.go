package capacontrollers

import (
	"net/http"
	"strconv"

	service "servidorMetadatos/capaServices"

	"github.com/gin-gonic/gin"
)

// MetadataRuidoBlancoController expone los servicios REST de RuidoBlanco usando Gin.
type MetadataRuidoBlancoController struct {
	service *service.MetadataRuidoBlancoService
}

func NewMetadataRuidoBlancoController(service *service.MetadataRuidoBlancoService) *MetadataRuidoBlancoController {
	return &MetadataRuidoBlancoController{service: service}
}

// ConsultarRuidoBlanco - GET /ruidoblanco/:id
func (this *MetadataRuidoBlancoController) ConsultarRuidoBlanco(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"codigo":  400,
			"mensaje": "El ID debe ser un número entero válido",
		})
		return
	}

	respuesta := this.service.ConsultarRuidoBlanco(id)
	ctx.JSON(respuesta.Codigo, respuesta)
}
