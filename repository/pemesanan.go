package repository

import (
	"kereta/config"
	"kereta/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = config.Connect()
}

func (r *Repo) FindAllPemesananByAll(i interface{}, args ...interface{}) error {
	result := r.app.Joins("Kereta").Joins("Kategori").First(i, args...)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repo) FindPemesananByID(id int) (data models.Pemesanan, err error) {
	err = r.app.Joins("Kereta").Joins("Kategori").First(&data, id).Error
	return
}

func (r *Repo) FindDetailByID(id int) (data models.DetailKereta, err error) {
	err = r.app.First(&data, id).Error
	return
}

func (r *Repo) InsertPemesanan(pemesanan models.Pemesanan) (data models.Pemesanan, err error) {
	err = r.app.Create(&pemesanan).Error
	return pemesanan, nil
}

func (r *Repo) InsertDetail(kereta models.DetailKereta) (data models.DetailKereta, err error) {
	err = r.app.Create(&kereta).Error
	return kereta, nil
}
