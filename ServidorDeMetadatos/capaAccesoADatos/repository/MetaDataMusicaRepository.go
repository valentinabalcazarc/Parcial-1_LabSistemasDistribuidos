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

// CargarMetadataMusica inicializa el vector con canciones de ejemplo.
func (this *MetadataMusicaRepository) CargarMetadataMusica() {
	var objMusica1, objMusica2, objMusica3, objMusica4 entity.MetaDataMusica

	// Canción 1
	objMusica1.SetId(1)
	objMusica1.SetTipo("Música")
	objMusica1.SetArtistaPrincipal("Bohemian Rhapsody")
	objMusica1.SetArtistaPrincipal("Queen")
	objMusica1.SetAlbum("A Night at the Opera")
	objMusica1.SetGenero("Rock")
	objMusica1.SetTituloCancion("Bohemian Rhapsody")
	objMusica1.SetSelloDiscografico("EMI Records")
	objMusica1.SetAnioLanzamiento("1975")

	// Canción 2
	objMusica2.SetId(2)
	objMusica2.SetTipo("Música")
	objMusica2.SetArtistaPrincipal("Michael Jackson")
	objMusica2.SetAlbum("Thriller")
	objMusica2.SetGenero("Pop")
	objMusica2.SetTituloCancion("Billie Jean")
	objMusica2.SetSelloDiscografico("Epic Records")
	objMusica2.SetAnioLanzamiento("1982")

	// Canción 3
	objMusica3.SetId(3)
	objMusica3.SetTipo("Música")
	objMusica3.SetArtistaPrincipal("Daft Punk")
	objMusica3.SetAlbum("Random Access Memories")
	objMusica3.SetGenero("Pop")
	objMusica3.SetTituloCancion("Get Lucky")
	objMusica3.SetSelloDiscografico("Columbia Records")
	objMusica3.SetAnioLanzamiento("2013")

	// Canción 4
	objMusica4.SetId(4)
	objMusica4.SetTipo("Música")
	objMusica4.SetArtistaPrincipal("Coldplay")
	objMusica4.SetAlbum("A Rush of Blood to the Head")
	objMusica4.SetGenero("Jazz")
	objMusica4.SetTituloCancion("The Scientist")
	objMusica4.SetSelloDiscografico("Parlophone")
	objMusica4.SetAnioLanzamiento("2002")

	// Asignación al slice del repositorio
	this.vectorMetadataMusica = []entity.MetaDataMusica{
		objMusica1,
		objMusica2,
		objMusica3,
		objMusica4,
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
