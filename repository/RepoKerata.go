package repository

import (
	"kereta/models"
)

func (r *Repo) FindAll() ([]models.Stasiun, error) {
	var data []models.Stasiun
	result := r.app.Joins("Kota").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (r *Repo) FindOneId(i interface{}, args ...interface{}) error {
	result := r.app.Joins("Kota").First(i, args...)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repo) FindAllKotaId(id string) ([]models.Stasiun, error) {
	var data []models.Stasiun
	result := r.app.Joins("Kota").Where("kota_id = ?", id).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
