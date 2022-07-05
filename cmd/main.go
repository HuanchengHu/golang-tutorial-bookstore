package main

import (
	"log"
	"net/http"

	"github.com/HuanchengHu/golang-tutorial-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
