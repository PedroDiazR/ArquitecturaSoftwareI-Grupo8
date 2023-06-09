package db

import (
	"go-api/model"

	hotelClient "go-api/clients/hotel"
	imageClient "go-api/clients/image"
	reservationClient "go-api/clients/reservation"
	userClient "go-api/clients/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	//Parametros de coneccion
	DBNombre := "tpintegrador"
	DBUser := "adminTP"
	DBPass := "adminTP"
	DBHost := "localhost"

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBNombre)

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	hotelClient.Db = db
	imageClient.Db = db
	reservationClient.Db = db
	userClient.Db = db

}

func StartDbEngine() {
	db.AutoMigrate(&model.Hotel{})
	db.AutoMigrate(&model.Amenitie{})
	db.AutoMigrate(&model.Image{})
	db.AutoMigrate(&model.Reservation{})
	db.AutoMigrate(&model.User{})

	log.Info("Llegue hasta aca pana")
}
