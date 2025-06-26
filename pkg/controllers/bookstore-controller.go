package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore-mysql-crud/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	db, books := models.GetAllBooks()
	if db.Error != nil {
		http.Error(w, db.Error.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	// "auto-detect the base from the string, and parse as an int64 of any size."
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if db.Error != nil {
		http.Error(w, db.Error.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(bookDetails)
}
