package controllers

import (
	"context"
	"encoding/json"
	"kereta/models"
	"kereta/tools"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (c *BaseControllerSecond) GetData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	redisData, err := REDIS.Get(ctx, "List Kategoris").Result()

	if err != nil {
		log.Println("RGet Database !! ")
		listkategori, err := c.us.GetDataKategori()
		if err != nil {
			Respon(w, 500, nil, "Internal Server Error")
			return
		}
		data, err := json.Marshal(listkategori)
		if err != nil {
			Respon(w, 500, nil, "Internal Server Error")
			return
		}

		err = REDIS.Set(ctx, "kategoris", (data), 0).Err()
		if err != nil {
			log.Println("Redis Error", err)
			log.Println("Error Set Redis")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		w.WriteHeader(200)
		return

	}

	log.Println("Get Redis ")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(redisData))
	w.WriteHeader(200)
}

func PostData(w http.ResponseWriter, r *http.Request) {
	var user map[string]string
	method := "POST"
	baseUrl := "http://localhost:8002/api/k2/kategori"

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		Respon(w, 400, nil, "Bad Request")
		return
	}

	data := models.Kategori{
		Nama: user["nama"],
	}

	dataString, _ := json.Marshal(data)
	code, result, err := tools.HTTPResponse(method, baseUrl, string(dataString), nil)
	if err != nil {
		Respon(w, code, nil, err.Error())
		return
	}

	log.Println(result)

	if data.Nama == "" {
		Respon(w, 400, nil, "Bad Request")
		return
	}

	DB.Create(&data)
	Respon(w, 200, data, "Succes Menambahkan Kategori")
}

func (c *BaseControllerSecond) UpdateData(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Kategori
	err := decoder.Decode(&datarequest)
	if err != nil {
		Respon(w, 500, nil, "Internal Server Error")
		return
	}
	Id, err := strconv.Atoi(id)
	if err != nil {
		Respon(w, 500, nil, "Invalid Request")
		return
	}
	datarequest.Id = Id
	err = c.us.UpdateDataKategori(datarequest)
	if err != nil {
		Respon(w, 500, nil, "Invalid Request")
		return
	}
	ctx := context.Background()

	REDIS.Del(ctx, "kategoris")

	Respon(w, 200, nil, "Sukses Update Data")
	return
}
