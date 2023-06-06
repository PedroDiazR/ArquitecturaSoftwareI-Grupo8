package controllers

// Funciones que hay que poner
// get y post

import (
	"dto/hotels_dto/hotel_dto.go"
	"fmt"
	"go-api/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	paramItemID = "itemID"
)

func GetHotels(c *gin.Context) {
	log.Println("get hotel en proceso")
	// Get id param from URL as string
	var HotelsDto hotels_dto.HotelsDto
	hotelsDto, err := service.HotelSer

	// Convert string ID to int ID
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		fmt.Println("Error parsing item ID", err)
		ctx.JSON(http.StatusBadRequest, err) // Then we will create a custom error struct
		return
	}

	// Call the service with int ID
	services.ItemClient = services.MlClient{}
	item, err := services.GetItem(id)
	if err != nil {
		fmt.Println("Error getting item", err)
		ctx.JSON(http.StatusInternalServerError, err) // Then we will create a custom error struct
		return
	}

	// Successful case
	ctx.JSON(http.StatusOK, item)
}
