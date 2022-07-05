package routes

import (
	"github.com/HuanchengHu/golang-tutorial-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.Handlefunc("/book/", controllers.CreateBook).Methods("POST")
	router.Handlefunc("/book/", controllers.GetBook).Methods("GET")
	router.Handlefunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
