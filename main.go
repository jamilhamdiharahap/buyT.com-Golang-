package main

import (
	"kereta/config"
	"kereta/controllers"
	"kereta/controllers/auth"
	"kereta/repository"
	"kereta/usecase"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	connection := config.Connect()
	repositoryf := repository.NewRepository(connection)
	usecasef := usecase.NewUsecase(repositoryf)
	controller := controllers.NewController(usecasef)

	repositorySecond := repository.NewRepositorySecond(connection)
	usecaseSecond := usecase.NewUsecaseSecond(repositorySecond)
	controllerSecond := controllers.NewControllerSecond(usecaseSecond)

	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(auth.GetTokenJwt)
		// daftar stasiun
		r.Get("/api/s1/stasiuns", controller.GetDataStasiun)      //jalan
		r.Get("/api/s1/stasiuns/{id}", controller.GetStasiunById) //jalan
		// daftar kota yang memiliki stasiun
		r.Get("/api/k1/kota/{id}", controller.GetKotaById) //jalan
		// pemesanan
		r.Get("/api/k2/kereta/pemesanan/{id}", controller.GetServisPemesanan)    //jalan
		r.Get("/api/p2/pemesanan/detailpemesanan/{id}", controller.GetPemesanan) //jalan
		// microservices
		r.Post("/api/p1/kereta/pemesanan", controller.Pemesanan)     //jalan
		r.Post("/api/k1/post/ketegoris", controllers.PostData)       //jalan
		r.Post("/api/p1/post/detailkereta", controller.DetailKereta) //jalan
		// kategoris
		r.Get("/api/k1/get/kategoris", controllerSecond.GetData)         //jalan
		r.Put("/api/k1/put/kategoris/{id}", controllerSecond.UpdateData) //jalan
		r.Delete("/api/k1/delete/detailkereta/{id}", controllerSecond.DeleteDataDetailKereta)
	})

	router.Group(func(r chi.Router) {
		r.Use(auth.ApiKey)
		r.Post("/api/au1/register", auth.Register)
		r.Post("/api/au2/login", auth.Login)
	})

	if err := http.ListenAndServe(":"+os.Getenv("HOST")+"", router); err != nil {
		log.Fatal(err)
	}
	log.Println("Server Running on port 8001")
}
