package repository

import "servidorMetadatos/capaAccesoADatos/entity"

// MetadataMusicaRepository es un repositorio que mantiene en memoria
// el slice de música y expone las operaciones de búsqueda y carga.
type MetadataMusicaRepository struct {
	vectorMetadataMusica []entity.MetaDataMusica
}

// NewMetadataMusicaRepository crea el repositorio y lo precarga con la
// metadata de canciones de ejemplo.
func NewMetadataMusicaRepository() *MetadataMusicaRepository {
	this := &MetadataMusicaRepository{}
	this.CargarMetadataMusica()
	return this
}

// CargarMetadataMusica inicializa el vector con 2 canciones de ejemplo.
func (this *MetadataMusicaRepository) CargarMetadataMusica() {
	var objMusica1, objMusica2 entity.MetaDataMusica

	tipoMusica := entity.NewTipoAudio(1, "Música")

	// Canción 1
	objMusica1.SetId(1)
	objMusica1.SetTipo(tipoMusica)
	objMusica1.SetArtistaPrincipal("Queen")
	objMusica1.SetAlbum("A Night at the Opera")
	objMusica1.SetGenero("Rock")
	objMusica1.SetTituloCancion("Bohemian Rhapsody")
	objMusica1.SetSelloDiscografico("EMI Records")
	objMusica1.SetAnioLanzamiento("1975")

	// Canción 2
	objMusica2.SetId(2)
	objMusica2.SetTipo(tipoMusica)
	objMusica2.SetArtistaPrincipal("Michael Jackson")
	objMusica2.SetAlbum("Thriller")
	objMusica2.SetGenero("Pop")
	objMusica2.SetTituloCancion("Billie Jean")
	objMusica2.SetSelloDiscografico("Epic Records")
	objMusica2.SetAnioLanzamiento("1982")

	// Asignación al slice del repositorio
	this.vectorMetadataMusica = []entity.MetaDataMusica{
		objMusica1,
		objMusica2,
	}
}

// BuscarMusica recorre el vector buscando una canción por su ID. Retorna la
// canción encontrada y un booleano que indica si la búsqueda tuvo éxito.
func (this *MetadataMusicaRepository) BuscarMusica(id int) (entity.MetaDataMusica, bool) {
	for _, musica := range this.vectorMetadataMusica {
		if musica.GetId() == id {
			return musica, true
		}
	}

	return entity.MetaDataMusica{}, false
}

func (this *MetadataMusicaRepository) ListarMusica() []entity.MetaDataMusica {
	return this.vectorMetadataMusica
}
