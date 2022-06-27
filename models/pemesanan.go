package models

type Pemesanan struct {
	Id         int `gorm:primary_key,autoincrement;json:"id"`
	KeretaId   int `json:"tujuan_id"`
	Kereta     Kereta
	AsalId     int `json:"asal_id"`
	KategoriId int `json:"ketegori_id"`
	Kategori   Kategori
}

type ResponPemesanan struct {
	Pemesanan
	Status string `json:"status"`
}

type DetailKereta struct {
	ID         uint `gorm:"primary_key"`
	KeretaId   int  `json:"kereta_id"`
	Kereta     Kereta
	KategoriId int `json:"kategori_id"`
	Kategori   Kategori
}

type Kereta struct {
	Id    int    `gorm:"primary_key;auto_increment" json:"id"`
	Nama  string `json:"nama"`
	Harga uint64 `json:"harga"`
}

type Kategori struct {
	Id   int    `gorm:"primary_key;auto_increment" json:"id"`
	Nama string `json:"nama"`
}
