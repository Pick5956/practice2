package entity

import(
	"gorm.io/gorm"
)

type Room struct{
	gorm.Model
	Name string
	Type string
	Price float64
}