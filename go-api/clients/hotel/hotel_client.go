package clients

import (
	"go-api/model"
)

func GetHotel() {
	var hotel model.Hotel

	hotel.Id = 1
	hotel.Description = "Prueba"
	hotel.Name = "eia"
	log.debug("hotel", hotel)
	return hotel

}
