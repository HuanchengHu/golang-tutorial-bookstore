package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
