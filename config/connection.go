package config

import (
	"kereta/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Load Config Files !! ")
	}
}

func Connect() *gorm.DB {
	url := os.Getenv("DATABASE_URL")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	connect := user + ":" + password + "@tcp(" + url + ":" + port + ")/" + dbname
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(models.Kategori{}, models.Kereta{}, models.DetailKereta{}, models.Stasiun{}, models.Kota{}, models.User{}, models.Pemesanan{})
	return db
}
