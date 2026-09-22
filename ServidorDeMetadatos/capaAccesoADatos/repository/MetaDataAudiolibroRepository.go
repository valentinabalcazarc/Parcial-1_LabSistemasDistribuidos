package repository

import "servidorMetadatos/capaAccesoADatos/entity"

// MetadataAudioRepository es un repositorio que mantiene en memoria
// el slice de audios y expone las operaciones de búsqueda y registro.
type MetadataAudiolibroRepository struct {
	vectorMetadataAudiolibros []entity.MetaDataAudiolibro
}

// NewMetadataAudioRepository crea el repositorio y lo precarga con la
// metadata de audios de ejemplo.
func NewMetadataAudiolibroRepository() *MetadataAudiolibroRepository {
	this := &MetadataAudiolibroRepository{}
	this.CargarMetadataAudiolibros()
	return this
}

// CargarMetadataAudios inicializa el vector con 5 audios de ejemplo.
func (this *MetadataAudiolibroRepository) CargarMetadataAudiolibros() {
	var objAudio1, objAudio2, objAudio3, objAudio4 entity.MetaDataAudiolibro

	// Audiolibro 1
	objAudio1.SetId(1)
	objAudio1.SetTipo("Audiolibro")
	objAudio1.SetTituloLibro("Cien Años de Soledad")
	objAudio1.SetAutor("Gabriel García Márquez")
	objAudio1.SetNarrador("Gustavo Bonfigli")
	objAudio1.SetEditorial("Penguin Random House")
	objAudio1.SetISBN(978030747)
	objAudio1.SetCapitulo(1)

	// Audiolibro 2
	objAudio2.SetId(2)
	objAudio2.SetTipo("Audiolibro")
	objAudio2.SetTituloLibro("El Principito")
	objAudio2.SetAutor("Antoine de Saint-Exupéry")
	objAudio2.SetNarrador("José María Carrascal")
	objAudio2.SetEditorial("Salamandra")
	objAudio2.SetISBN(978847888)
	objAudio2.SetCapitulo(3)

	// Audiolibro 3
	objAudio3.SetId(3)
	objAudio3.SetTipo("Audiolibro")
	objAudio3.SetTituloLibro("1984")
	objAudio3.SetAutor("George Orwell")
	objAudio3.SetNarrador("Raúl Llorens")
	objAudio3.SetEditorial("Debolsillo")
	objAudio3.SetISBN(978849989)
	objAudio3.SetCapitulo(5)

	// Audiolibro 4
	objAudio4.SetId(4)
	objAudio4.SetTipo("Audiolibro")
	objAudio4.SetTituloLibro("Don Quijote de la Mancha")
	objAudio4.SetAutor("Miguel de Cervantes")
	objAudio4.SetNarrador("Juan Echanove")
	objAudio4.SetEditorial("Planeta")
	objAudio4.SetISBN(978840806)
	objAudio4.SetCapitulo(12)

	// Asignación al slice del repositorio
	this.vectorMetadataAudiolibros = []entity.MetaDataAudiolibro{
		objAudio1,
		objAudio2,
		objAudio3,
		objAudio4,
	}
}

// BuscarAudio recorre el vector buscando un audio por su título. Retorna el
// audio encontrado y un booleano que indica si la búsqueda tuvo éxito.
func (this *MetadataAudiolibroRepository) BuscarAudio(id int) (entity.MetaDataAudiolibro, bool) {
	for _, audio := range this.vectorMetadataAudiolibros {
		if audio.GetId() == id {
			return audio, true
		}
	}

	return entity.MetaDataAudiolibro{}, false
}
