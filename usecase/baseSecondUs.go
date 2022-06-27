package usecase

import (
	"kereta/models"
	"kereta/repository"
)

type UcSecond struct {
	query repository.RepositoryInterfaceSecond
}

type UsecaseInterfaceSecond interface {
	GetDataKategori() ([]models.Kategori, error)
	InsertDataKategori(models.Kategori) error
	UpdateDataKategori(models.Kategori) error
	DeleteDataKereta(id int) error
}

func NewUsecaseSecond(r repository.RepositoryInterfaceSecond) UsecaseInterfaceSecond {
	return &UcSecond{query: r}
}
