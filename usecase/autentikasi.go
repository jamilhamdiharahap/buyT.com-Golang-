package usecase

import "kereta/models"

func (r *Uc) Register(data models.User) (models.User, error) {
	err := r.query.Create(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
