package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Suandika12/go_Hidangan/pkg/config"
)

var db *gorm.DB

type Hidangan struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Sku string `json:"Sku"`
	Description string `json:"description"`
	Price  string `json:"price"`
	Quantity string `json:"quantity"`
	Image string `json:"image"`
	Category string `json:"category"`

}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Hidangan{})
}
func (b *Hidangan) CreateHidangan() *Hidangan {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllHidangans() []Hidangan {
	var Hidangans []Hidangan
	db.Find(&Hidangans)
	return Hidangans
}
func GetHidanganbyId(name int64) (*Hidangan, *gorm.DB) {
	var getHidangan Hidangan
	db := db.Where("Name=?", name).Find(&getHidangan)
	return &getHidangan, db
}
func DeleteHidangan(name int64) Hidangan {
	var hidangan Hidangan
	db.Where("Name=?", name).Delete(hidangan)
	return hidangan
}
