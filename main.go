package main

import (
	"log"
	"net/http"

	"github.com/MeganViga/GoWithMYSQL/routes"
	"github.com/gorilla/mux"
)


func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	r.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9090",r))
}