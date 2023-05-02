package controllers

import (
	"encoding/json"
	// "fmt"
	// "io/ioutil"
	"net/http"
	"strconv"

	"github.com/MeganViga/GoWithMYSQL/models"
	"github.com/MeganViga/GoWithMYSQL/utils"
	"github.com/gorilla/mux"
)

type NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request){
	newbooks := models.GetAllBooks()
	res,_ := json.Marshal(newbooks)
	w.Header().Set("Content-Type","applications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id ,_:= strconv.ParseInt(params["bookid"],0,0)
	newbook ,_ := models.GetBookByID(id)
	res, _ := json.Marshal(newbook)
	w.Header().Set("Content-Type","applications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r,CreateBook)
	// body, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(body,CreateBook)
	// fmt.Println("Createbook:",CreateBook,body)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type","applications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id ,_:= strconv.ParseInt(params["bookid"],0,0)
	deletedbook, _ := models.DeleteBook(id)
	res, _ := json.Marshal(deletedbook)
	w.Header().Set("Content-Type","applications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)
	params := mux.Vars(r)
	id ,_:= strconv.ParseInt(params["bookid"],0,0)
	bookDetails, db := models.GetBookByID(id)
	if UpdateBook.Name != ""{
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author!= ""{
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != ""{
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","applications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}