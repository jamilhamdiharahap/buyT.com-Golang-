package models

type Stasiun struct {
	Id     int    `gorm:"primary_key;auto_increment" json:"id"`
	Nama   string `json:"nama"`
	KotaId int    `json:"kota_id"`
	Kota   Kota
}

type Kota struct {
	Id   int    `gorm:"primary_key;auto_increment" json:"id"`
	Nama string `json:"nama"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
