package usecase

import (
	"kereta/models"
)

func (r *Uc) Stasiun() ([]models.Stasiun, error) {
	data, err := r.query.FindAll()
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *Uc) GetFirstStasiun(id string) (models.Stasiun, error) {
	var data models.Stasiun
	err := r.query.FindOneId(&data, id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *Uc) GetAllStasiunByKota(id string) ([]models.Stasiun, error) {
	data, err := r.query.FindAllKotaId(id)
	if err != nil {
		return data, err
	}
	return data, nil
}
