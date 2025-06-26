package models

import (
	"go-bookstore-mysql-crud/pkg/config"

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

func (b *Book) CreateBook() *Book {
	if !DB.NewRecord(b) {
		// The record is not new, so you might want to handle this case
		// For example, return nil or an error, or just skip creation
		return nil
	}
	DB.Create(&b) // This creates the record in the database
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	// If you pass a slice of a type that is not mapped to a table in the DB (i.e., not a GORM model),
	// GORM will not know how to map it to a table and will return an error.
	// For example, if you define:
	// type Foo struct { Bar string }
	// and call DB.Find(&[]Foo{}), GORM will look for a table named "foos" and fail if it doesn't exist.
	DB.Find(&Books) // This retrieves all records from the database
	return Books
}
