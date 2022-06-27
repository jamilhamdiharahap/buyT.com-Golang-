package models

type RequestNotify struct {
	IdPemesanan int         `json:"pemesanan_id"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}
