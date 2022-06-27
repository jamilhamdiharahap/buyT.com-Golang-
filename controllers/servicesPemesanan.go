package controllers

import (
	"kereta/models"
	"kereta/tools"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (c *BaseController) GetServisPemesanan(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	method := "GET"
	baseUrl := "http://localhost:8002/api/k4/detail/" + id
	header := map[string][]string{
		"Content-type": {"application/json"},
	}
	var ResponseServiceKereta models.ResponseServiceKereta
	err := tools.ServicesKereta(method, baseUrl, nil, &ResponseServiceKereta, header)
	if err != nil {
		log.Println("Error Hit Service Notification=> ", err)
		Respon(w, 500, nil, "Internal Server Error")
		return
	}

	Respon(w, 200, ResponseServiceKereta.Kereta, "Get Data Success")

}
