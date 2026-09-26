package repository

import "servidorMetadatos/capaAccesoADatos/entity"

// MetadataPodcastRepository es un repositorio que mantiene en memoria
// el slice de podcasts y expone las operaciones de búsqueda y carga.
type MetadataPodcastRepository struct {
	vectorMetadataPodcasts []entity.MetaDataPodcast
}

// NewMetadataPodcastRepository crea el repositorio y lo precarga con la
// metadata de episodios de podcast de ejemplo.
func NewMetadataPodcastRepository() *MetadataPodcastRepository {
	this := &MetadataPodcastRepository{}
	this.CargarMetadataPodcasts()
	return this
}

// CargarMetadataPodcasts inicializa el vector con episodios de ejemplo.
func (this *MetadataPodcastRepository) CargarMetadataPodcasts() {
	var objPod1, objPod2, objPod3, objPod4 entity.MetaDataPodcast

	tipoPodcast := entity.NewTipoAudio(3, "Podcast")

	// Podcast 1
	objPod1.SetId(1)
	objPod1.SetTipo(tipoPodcast)
	objPod1.SetNombre("Radio Ambulante")
	objPod1.SetTituloEpisodio("El Polizón")
	objPod1.SetNumeroTemporada(12)
	objPod1.SetNotasShow("Historia de un viaje inesperado hacia el norte.")
	objPod1.SetClasificacionContenido("Explicito")

	// Podcast 2
	objPod2.SetId(2)
	objPod2.SetTipo(tipoPodcast)
	objPod2.SetNombre("The Daily")
	objPod2.SetTituloEpisodio("Global Economy Trends")
	objPod2.SetNumeroTemporada(8)
	objPod2.SetNotasShow("Análisis semanal de los mercados internacionales.")
	objPod2.SetClasificacionContenido("Para toda la familia")

	// Podcast 3
	objPod3.SetId(3)
	objPod3.SetTipo(tipoPodcast)
	objPod3.SetNombre("Entiende Tu Mente")
	objPod3.SetTituloEpisodio("Gestionar la Ansiedad")
	objPod3.SetNumeroTemporada(5)
	objPod3.SetNotasShow("Consejos prácticos de psicología para el día a día.")
	objPod3.SetClasificacionContenido("Para toda la familia")

	// Podcast 4
	objPod4.SetId(4)
	objPod4.SetTipo(tipoPodcast)
	objPod4.SetNombre("La Pulla")
	objPod4.SetTituloEpisodio("¿Qué pasa con la educación?")
	objPod4.SetNumeroTemporada(6)
	objPod4.SetNotasShow("Un análisis crítico de la situación educativa actual.")
	objPod4.SetClasificacionContenido("Explicito")

	// Asignación al slice del repositorio
	this.vectorMetadataPodcasts = []entity.MetaDataPodcast{
		objPod1,
		objPod2,
		objPod3,
		objPod4,
	}
}

// BuscarPodcast recorre el vector buscando un podcast por su ID. Retorna el
// podcast encontrado y un booleano que indica si la búsqueda tuvo éxito.
func (this *MetadataPodcastRepository) BuscarPodcast(id int) (entity.MetaDataPodcast, bool) {
	for _, podcast := range this.vectorMetadataPodcasts {
		if podcast.GetId() == id {
			return podcast, true
		}
	}

	return entity.MetaDataPodcast{}, false
}

func (this *MetadataPodcastRepository) ListarPodcasts() []entity.MetaDataPodcast {
	return this.vectorMetadataPodcasts
}
