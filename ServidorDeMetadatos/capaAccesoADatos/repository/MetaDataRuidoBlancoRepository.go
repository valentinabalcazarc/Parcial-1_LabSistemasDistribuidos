package repository

import "servidorMetadatos/capaAccesoADatos/entity"

// MetadataRuidoBlancoRepository es un repositorio que mantiene en memoria
// el slice de ruido blanco y expone las operaciones de búsqueda y carga.
type MetadataRuidoBlancoRepository struct {
	vectorMetadataRuidoBlanco []entity.MetaDataRuidoBlanco
}

// NewMetadataRuidoBlancoRepository crea el repositorio y lo precarga con la
// metadata de ruidos blancos de ejemplo.
func NewMetadataRuidoBlancoRepository() *MetadataRuidoBlancoRepository {
	this := &MetadataRuidoBlancoRepository{}
	this.CargarMetadataRuidoBlanco()
	return this
}

// CargarMetadataRuidoBlanco inicializa el vector con 2 ruidos blancos de ejemplo.
func (this *MetadataRuidoBlancoRepository) CargarMetadataRuidoBlanco() {
	var objRuido1, objRuido2 entity.MetaDataRuidoBlanco

	tipoRuidoBlanco := entity.NewTipoAudio(2, "Ruido Blanco")

	// Ruido 1: Ruido Blanco / Lluvia
	objRuido1.SetId(1)
	objRuido1.SetTipo(tipoRuidoBlanco)
	objRuido1.SetTipoSonido("Ruido Blanco")
	objRuido1.SetFuenteAudio("Lluvia")
	objRuido1.SetUsoSugerido("Dormir")
	objRuido1.SetProveedor("SleepSounds Studio")
	objRuido1.SetDuracion(60) // Duración del bucle en segundos (1 min)
	objRuido1.SetFrecuenciaDominante("Agudos")

	// Ruido 2: Ruido Marrón / Ventilador
	objRuido2.SetId(2)
	objRuido2.SetTipo(tipoRuidoBlanco)
	objRuido2.SetTipoSonido("Ruido Marrón")
	objRuido2.SetFuenteAudio("Ventilador")
	objRuido2.SetUsoSugerido("Concentración")
	objRuido2.SetProveedor("FocusLoop Channel")
	objRuido2.SetDuracion(120) // Duración del bucle en segundos (2 min)
	objRuido2.SetFrecuenciaDominante("Graves")

	// Asignación al slice del repositorio
	this.vectorMetadataRuidoBlanco = []entity.MetaDataRuidoBlanco{
		objRuido1,
		objRuido2,
	}
}

// BuscarRuidoBlanco recorre el vector buscando un registro de ruido blanco por su ID.
func (this *MetadataRuidoBlancoRepository) BuscarRuidoBlanco(id int) (entity.MetaDataRuidoBlanco, bool) {
	for _, ruido := range this.vectorMetadataRuidoBlanco {
		if ruido.GetId() == id {
			return ruido, true
		}
	}

	return entity.MetaDataRuidoBlanco{}, false
}

func (this *MetadataRuidoBlancoRepository) ListarRuidoBlanco() []entity.MetaDataRuidoBlanco {
	return this.vectorMetadataRuidoBlanco
}
