package routes

import (
	"github.com/MeganViga/GoWithMYSQL/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router){
	r.HandleFunc("/books",controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books",controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{bookid}",controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/books/{bookid}",controllers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/{bookid}",controllers.UpdateBook).Methods("PUT")
}