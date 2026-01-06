package entity

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name   string  `valid:"required~Name is required,length(5|1000)~Name must be at least 5 char"` //ห้ามว่าง ห้ามสั้นกว่า 5 ตัว
	Price  float64 `valid:"required~Price must be greater than 0,range(0|99999)~Price must be greater than 0"` //ห้ามติดลบ ห้ามเป็น 0
	Price2 float64 `valid:"range(0|99999)~Price2 cannot be negative"` //ยังเป็น 0 ได้ แค่ห้ามติดลบ
	Type   string  `valid:"alpha~Type must be alphabet,required~Type is required"` //ห้ามว่าง ห้ามมีตัวเลขในข้อความ
	
}
