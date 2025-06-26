package models

import (
	"go-bookstore-mysql-crud/pkg/config"
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

// gorm.Model: This is an embedded struct from the GORM library, a popular ORM (Object Relational Mapper) for Go. By embedding gorm.Model, your Book struct automatically gets common fields for database models: ID, CreatedAt, UpdatedAt, and DeletedAt.

// GORM Tags: The struct tags like gorm:"primaryKey" tell GORM how to map this struct to a database table. For example, gorm:"primaryKey" marks the ID field as the primary key in the database.
type Book struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  string `json:"price"`
}

func init() {
	config.ConnectDatabase()
	DB = config.GetDatabase()
	DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*gorm.DB, *Book) {
	// 	Before saving to the database:
	// b.ID == 0 means the object is new and hasn’t been saved yet.
	// After saving to the database:
	// GORM will set the ID to the value assigned by the database (usually a positive integer).
	if b.ID != 0 {
		// The record is not new, so you might want to handle this case
		// For example, return nil or an error, or just skip creation
		return nil, nil
	}
	db := DB.Create(&b) // This creates the record in the database
	if db.Error != nil {
		log.Println("Error creating book")
		return nil, nil
	}
	return db, b
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

func DeleteBook(ID int64) (*gorm.DB, Book) {
	var book Book
	db := DB.Where("ID=?", ID).Delete(book)
	return db, book
}

func UpdateBook(id int64, updatedData *Book) (*gorm.DB, *Book) {
	var book Book
	db := DB.First(&book, id)

	if db.Error == gorm.ErrRecordNotFound {
		log.Println("Book not found")
		return db, nil
	}
	if db.Error != nil {
		log.Println("Error updating book")
		return db, nil
	}

	db = DB.Model(&book).Updates(updatedData)
	return db, &book
}
