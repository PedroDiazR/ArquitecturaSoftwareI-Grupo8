package model

type Amenitie struct {
	Id     int      `gorm:"primaryKey"`
	Name   string   `gorm:"varchar(100);not null"`
	Hotels []*Hotel `gorm:"many2many:hotel_amenities"`
}

type Amenities []Amenitie
