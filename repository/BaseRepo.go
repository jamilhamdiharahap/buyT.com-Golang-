package repository

import (
	"kereta/models"

	"gorm.io/gorm"
)

type Repo struct {
	app *gorm.DB
}

type RepositoryInterface interface {
	FindAll() ([]models.Stasiun, error)
	FindOneId(i interface{}, args ...interface{}) error
	FindAllPemesananByAll(i interface{}, args ...interface{}) error
	FindAllKotaId(i string) ([]models.Stasiun, error)
	Create(i interface{}) error
	InsertPemesanan(pemesanan models.Pemesanan) (data models.Pemesanan, err error)
	InsertDetail(detail models.DetailKereta) (data models.DetailKereta, err error)
	FindPemesananByID(id int) (data models.Pemesanan, err error)
	FindDetailByID(id int) (data models.DetailKereta, err error)
}

func NewRepository(app *gorm.DB) RepositoryInterface {
	return &Repo{app}
}
