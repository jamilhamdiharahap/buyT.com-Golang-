package repository

import (
	"kereta/models"

	"gorm.io/gorm"
)

type RepoSecond struct {
	app *gorm.DB
}

type RepositoryInterfaceSecond interface {
	InsertData(i interface{}) error
	FindAll(i interface{}) error
	UpdateData(i interface{}, where map[string]interface{}, data map[string]interface{}) error
	DeleteData(i interface{}) error
	ObjectByID(id int) (models models.Kategori, err error)
	ObjectByIdDetail(id int) (models models.DetailKereta, err error)
}

func NewRepositorySecond(app *gorm.DB) RepositoryInterfaceSecond {
	return &RepoSecond{app}
}
