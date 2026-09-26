package repository

import "servidorMetadatos/capaAccesoADatos/entity"

type TipoAudioRepository struct {
	vectorTipos []entity.TipoAudio
}

func NewTipoAudioRepository() *TipoAudioRepository {
	repo := &TipoAudioRepository{}
	repo.CargarTipos()
	return repo
}

func (r *TipoAudioRepository) CargarTipos() {
	r.vectorTipos = []entity.TipoAudio{
		entity.NewTipoAudio(1, "Música"),
		entity.NewTipoAudio(2, "Ruido Blanco"),
		entity.NewTipoAudio(3, "Podcast"),
		entity.NewTipoAudio(4, "Audiolibro"),
	}
}

func (r *TipoAudioRepository) ListarTipos() []entity.TipoAudio {
	return r.vectorTipos
}
