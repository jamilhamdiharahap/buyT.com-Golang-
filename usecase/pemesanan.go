package usecase

import (
	"errors"
	"kereta/models"
	"log"
)

func (r *Uc) PemesananKereta(stasiunA int, stasiunT int, ktg int) ([]models.Stasiun, []models.Kategori, error) {
	var stasiun []models.Stasiun
	var kategori []models.Kategori
	err := r.query.FindOneId(&stasiun, &kategori, stasiunA, stasiunT, ktg)
	if err != nil {
		return stasiun, kategori, err
	}
	return stasiun, kategori, nil
}

func (r *Uc) FindAllPemesananById(id int) (models.Pemesanan, error) {
	var data models.Pemesanan
	err := r.query.FindAllPemesananByAll(&data, id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *Uc) InsertPemesanan(pemesanan models.Pemesanan) (data models.Pemesanan, err error) {
	insert, err := r.query.InsertPemesanan(pemesanan)
	if err != nil {
		return data, errors.New("gagal insert pemesanan")
	}

	log.Println(insert)

	data, err = r.query.FindPemesananByID(insert.Id)
	if err != nil {
		return data, errors.New("pemesanan tidak ditemukan")
	}

	return
}

func (r *Uc) InsertDetail(pemesanan models.DetailKereta) (data models.DetailKereta, err error) {
	insert, err := r.query.InsertDetail(pemesanan)
	if err != nil {
		return data, errors.New("gagal insert pemesanan")
	}

	log.Println(insert)

	data, err = r.query.FindDetailByID(int(insert.ID))
	if err != nil {
		return data, errors.New("pemesanan tidak ditemukan")
	}

	return
}
