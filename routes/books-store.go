package routes

import (
	"github.com/gorilla/mux"
	"github.com/MeganViga/controllers"
)

var registeBookStoreRoutes = func(r *mux.Router){
	r.HandleFunc("/books",createBook).Methods("POST")
	r.HandleFunc("/books",getBooks).Methods("GET")
	r.HandleFunc("/books/{bookid}",getBook).Methods("GET")
	r.HandleFunc("/books/{bookid}",deleteBook).Methods("DELETE")
	r.HandleFunc("/books/{bookid}",updateBook).Methods("PUT")
}