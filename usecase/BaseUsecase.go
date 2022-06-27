package usecase

import (
	"kereta/models"
	"kereta/repository"
)

type Uc struct {
	query repository.RepositoryInterface
}

type UsecaseInterface interface {
	Stasiun() ([]models.Stasiun, error)
	GetFirstStasiun(id string) (models.Stasiun, error)
	GetAllStasiunByKota(id string) ([]models.Stasiun, error)
	Register(data models.User) (models.User, error)
	FindAllPemesananById(Id int) (models.Pemesanan, error)
	InsertPemesanan(pemesanan models.Pemesanan) (data models.Pemesanan, err error)
	InsertDetail(pemesanan models.DetailKereta) (data models.DetailKereta, err error)
}

func NewUsecase(r repository.RepositoryInterface) UsecaseInterface {
	return &Uc{query: r}
}
