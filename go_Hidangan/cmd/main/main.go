package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Suandika12/go_Hidangan/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterHidangansRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
