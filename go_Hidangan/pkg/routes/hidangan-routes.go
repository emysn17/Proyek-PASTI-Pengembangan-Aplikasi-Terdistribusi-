package routes

import (
	"github.com/gorilla/mux"
	"github.com/Suandika12/go_Hidangan/pkg/controllers"
)

var RegisterHidangansRoutes = func(router *mux.Router) {
	router.HandleFunc("/hidangan/",
		controllers.CreateHidangan).Methods("POST")
	router.HandleFunc("/hidangan/", controllers.GetHidangan).Methods("GET")
	router.HandleFunc("/hidangan/{hidanganId}",
		controllers.GetHidanganById).Methods("GET")
	router.HandleFunc("/hidangan/{hidanganId}",
		controllers.UpdateHidangan).Methods("PUT")
	router.HandleFunc("/hidangan/{hidanganId}",
		controllers.DeleteHidangan).Methods("DELETE")
}
