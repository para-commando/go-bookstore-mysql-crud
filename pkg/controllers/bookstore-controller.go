package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore-mysql-crud/pkg/models"
	"go-bookstore-mysql-crud/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// var NewBook models.Book

// GetBooks godoc
// @Summary Get all books
// @Description Retrieve all books from the database
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} models.BookResponse "List of books"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /books [get]
func GetBooks(w http.ResponseWriter, r *http.Request) {
	db, books := models.GetAllBooks()
	if db.Error != nil {
		fmt.Printf("error while db operation: %v", db.Error.Error())
		http.Error(w, db.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

// GetBookById godoc
// @Summary Get a book by ID
// @Description Retrieve a specific book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.BookResponse "Book details"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /books/{id} [get]
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	// "auto-detect the base from the string, and parse as an int64 of any size."
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("error while parsing: %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bookDetails, db := models.GetBookById(ID)
	if db.Error != nil {
		fmt.Printf("error while db operation: %v", db.Error.Error())
		http.Error(w, db.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookDetails)
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the database
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.BookRequest true "Book object"
// @Success 200 {object} models.BookResponse "Created book"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid book data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /books [post]
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b, _, err := CreateBook.CreateBook()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Remove a book from the database by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.BookResponse "Deleted book"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /books/{id} [delete]
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, book, err := models.DeleteBook(ID)
	if err != nil {
		fmt.Printf("error while parsing: %v", err.Error())

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update an existing book's information
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.BookRequest true "Updated book object"
// @Success 200 {object} models.BookResponse "Updated book"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid ID or book data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /books/{id} [put]
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)

	_, book, err := models.UpdateBook(ID, updatedBook)

	if err != nil {
		fmt.Printf("error while db operation: %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
