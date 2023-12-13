package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Suandika12/go_Hidangan/pkg/models"
	"github.com/Suandika12/go_Hidangan/pkg/utils"
)

var NewHidangan models.Hidangan

func GetHidangan(w http.ResponseWriter, r *http.Request) {
	newHidangans := models.GetAllHidangans()
	res, _ := json.Marshal(newHidangans)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetHidanganById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hidanganId := vars["hidanganId"]
	Name, err := strconv.ParseInt(hidanganId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	hidanganDetails, _ := models.GetHidanganbyId(Name)
	res, _ := json.Marshal(hidanganDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateHidangan(w http.ResponseWriter, r *http.Request) {
	CreateHidangan := &models.Hidangan{}
	utils.ParseBody(r, CreateHidangan)
	b := CreateHidangan.CreateHidangan()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteHidangan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hidanganId := vars["hidangannId"]
	Name, err := strconv.ParseInt(hidanganId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	hidangan := models.DeleteHidangan(Name)
	res, _ := json.Marshal(hidangan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateHidangan(w http.ResponseWriter, r *http.Request) {
	var updateHidangan = &models.Hidangan{}
	utils.ParseBody(r, updateHidangan)
	vars := mux.Vars(r)
	hidanganId := vars["hidanganId"]
	Name, err := strconv.ParseInt(hidanganId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	hidanganDetails, db := models.GetHidanganbyId(Name)
	if updateHidangan.Name != "" {
		hidanganDetails.Name = updateHidangan.Name
	}
	if updateHidangan.Sku != "" {
		hidanganDetails.Sku = updateHidangan.Sku
	}
	if updateHidangan.Description != "" {
		hidanganDetails.Description = updateHidangan.Description
	}
	if updateHidangan.Price != "" {
		hidanganDetails.Price = updateHidangan.Price
	}
	if updateHidangan.Quantity != "" {
		hidanganDetails.Quantity = updateHidangan.Quantity
	}
	if updateHidangan.Image!= "" {
		hidanganDetails.Image = updateHidangan.Image
	}
	if updateHidangan.Category!= "" {
		hidanganDetails.Category = updateHidangan.Category
	}
	db.Save(&hidanganDetails)
	res, _ := json.Marshal(hidanganDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
