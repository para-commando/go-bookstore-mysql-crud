package models

import (
	"fmt"
	"go-bookstore-mysql-crud/pkg/config"
	"log"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

// Book represents a book in the bookstore
// @Description Book model for the bookstore API
type Book struct {
	gorm.Model
	// @Description Unique identifier for the book
	// @Example 1
	ID uint `json:"id" gorm:"primaryKey" example:"1"`

	// @Description Title of the book
	// @Example "The Great Gatsby"
	Title string `json:"title" example:"The Great Gatsby"`

	// @Description Author of the book
	// @Example "F. Scott Fitzgerald"
	Author string `json:"author" example:"F. Scott Fitzgerald"`

	// @Description Price of the book
	// @Example "$15.99"
	Price string `json:"price" example:"$15.99"`
}

// BookResponse represents the book response structure for API documentation
// @Description Book response model for API documentation
type BookResponse struct {
	// @Description Unique identifier for the book
	// @Example 1
	ID uint `json:"id" example:"1"`

	// @Description Title of the book
	// @Example "The Great Gatsby"
	Title string `json:"title" example:"The Great Gatsby"`

	// @Description Author of the book
	// @Example "F. Scott Fitzgerald"
	Author string `json:"author" example:"F. Scott Fitzgerald"`

	// @Description Price of the book
	// @Example "$15.99"
	Price string `json:"price" example:"$15.99"`

	// @Description When the book was created
	// @Example "2023-01-01T00:00:00Z"
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`

	// @Description When the book was last updated
	// @Example "2023-01-01T00:00:00Z"
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// BookRequest represents the book request structure for API documentation
// @Description Book request model for API documentation
type BookRequest struct {
	// @Description Title of the book
	// @Example "The Great Gatsby"
	Title string `json:"title" example:"The Great Gatsby" binding:"required"`

	// @Description Author of the book
	// @Example "F. Scott Fitzgerald"
	Author string `json:"author" example:"F. Scott Fitzgerald" binding:"required"`

	// @Description Price of the book
	// @Example "$15.99"
	Price string `json:"price" example:"$15.99" binding:"required"`
}

func init() {
	config.ConnectDatabase()
	DB = config.GetDatabase()
	DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, *gorm.DB, error) {
	// 	Before saving to the database:
	// b.ID == 0 means the object is new and hasn't been saved yet.
	// After saving to the database:
	// GORM will set the ID to the value assigned by the database (usually a positive integer).
	if b.ID != 0 {
		// The record is not new, so you might want to handle this case
		// For example, return nil or an error, or just skip creation
		return nil, nil, fmt.Errorf("book already exists with ID %d", b.ID)
	}
	db := DB.Create(&b) // This creates the record in the database
	if db.Error != nil {
		log.Println("Error creating book")
		return nil, db, db.Error
	}
	return b, db, nil
}

func GetAllBooks() (*gorm.DB, []Book) {
	var Books []Book
	// If you pass a slice of a type that is not mapped to a table in the DB (i.e., not a GORM model),
	// GORM will not know how to map it to a table and will return an error.
	// For example, if you define:
	// type Foo struct { Bar string }
	// and call DB.Find(&[]Foo{}), GORM will look for a table named "foos" and fail if it doesn't exist.
	db := DB.Find(&Books) // This retrieves all records from the database

	if db.Error == gorm.ErrRecordNotFound {
		log.Println("No books found")
		return nil, nil
	}

	return db, Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := DB.Where("ID=?", Id).Find(&getBook)

	if db.Error == gorm.ErrRecordNotFound {
		log.Println("Book not found")
		return nil, db
	}

	return &getBook, db
}

func DeleteBook(ID int64) (*gorm.DB, Book, error) {
	var book Book
	db := DB.Where("ID=?", ID).Delete(book)
	if db.Error == gorm.ErrRecordNotFound {
		log.Println("Book not found")
		return nil, book, db.Error
	}
	if db.Error != nil {
		log.Println("Error deleting book")
		return nil, book, db.Error
	}
	return db, book, nil
}

func UpdateBook(id int64, updatedData *Book) (*gorm.DB, *Book, error) {
	var book Book
	db := DB.First(&book, id)

	if db.Error == gorm.ErrRecordNotFound {
		log.Println("Book not found")
		return db, nil, db.Error
	}
	if db.Error != nil {
		log.Println("Error updating book")
		return db, nil, db.Error
	}

	db = DB.Model(&book).Updates(updatedData)
	return db, &book, nil
}
