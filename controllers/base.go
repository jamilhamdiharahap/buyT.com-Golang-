package controllers

import (
	"kereta/usecase"
	"net/http"
)

type BaseController struct {
	us usecase.UsecaseInterface
}

type Controller interface {
	GetDataStasiun(w http.ResponseWriter, r *http.Request)
	GetStasiunById(w http.ResponseWriter, r *http.Request)
	GetKotaById(w http.ResponseWriter, r *http.Request)
	GetServisPemesanan(w http.ResponseWriter, r *http.Request)
	Pemesanan(w http.ResponseWriter, r *http.Request)
	DetailKereta(w http.ResponseWriter, r *http.Request)
	GetPemesanan(w http.ResponseWriter, r *http.Request)
}

func NewController(us usecase.UsecaseInterface) Controller {
	return &BaseController{us: us}
}
