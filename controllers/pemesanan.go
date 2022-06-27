package controllers

import (
	"encoding/json"
	"kereta/config"
	"kereta/models"
	"kereta/tools"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = config.Connect()
}

func (c *BaseController) Pemesanan(w http.ResponseWriter, r *http.Request) {
	var pemesanan models.Pemesanan
	method := "POST"
	baseUrl := "https://buytnotifikasimessagebroker.herokuapp.com/publish"
	err := json.NewDecoder(r.Body).Decode(&pemesanan)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	insert, err := c.us.InsertPemesanan(pemesanan)
	if err != nil {
		Respon(w, 500, nil, err.Error())
		return
	}

	data := models.RequestNotify{
		IdPemesanan: insert.Id,
		Message:     "Selamat Pemesanan anda di proses",
		Data:        insert,
	}

	dataString, _ := json.Marshal(data)
	code, result, err := tools.HTTPNotif(method, baseUrl, string(dataString), nil)
	if err != nil {
		Respon(w, code, nil, err.Error())
		return
	}
	log.Println(result)

	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	Respon(w, 200, insert, "pemesanan Success")
}

func (c *BaseController) DetailKereta(w http.ResponseWriter, r *http.Request) {
	var detail models.DetailKereta
	method := "POST"
	baseUrl := "http://localhost:8002/api/k4/detailkereta"
	err := json.NewDecoder(r.Body).Decode(&detail)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	insert, err := c.us.InsertDetail(detail)
	if err != nil {
		Respon(w, 500, nil, err.Error())
		return
	}

	data := models.DetailKereta{
		KeretaId:   insert.KeretaId,
		KategoriId: insert.KategoriId,
	}

	dataString, _ := json.Marshal(data)
	code, result, err := tools.HTTPResponse(method, baseUrl, string(dataString), nil)
	if err != nil {
		Respon(w, code, nil, err.Error())
		return
	}
	log.Println(result)

	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	Respon(w, 200, insert, "Berhasil Menambahkan Kereta")
}

func (c *BaseController) GetPemesanan(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idPemesanan, err := strconv.Atoi(id)

	if err != nil {
		Respon(w, 500, nil, "Internal Server Error")
		return
	}

	Pemesanan, err := c.us.FindAllPemesananById(idPemesanan)
	if err != nil {
		Respon(w, 500, nil, "Internal Server Error")
		return
	}

	data, err := json.Marshal(Pemesanan)
	if err != nil {
		Respon(w, 500, nil, "Error marshall")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(200)
	return
}
