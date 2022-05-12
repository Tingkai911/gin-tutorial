package main

import (
	"errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = openDB()

// Database connection
func openDB() *gorm.DB {
	dsn := "root:password@tcp(localhost:3306)/gin_tutorial?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connect to db")
	}

	return db
}

func getBooksFromDB() ([]Book, error) {
	books := []Book{}
	result := db.Find(&books) // SELECT * FROM book

	if result.Error != nil {
		log.Println("fail to get result from db")
		return nil, errors.New("fail to get books from db")
	}

	return books, nil
}

// return *Book so that we can modify the attributes of book in a different function
func getBookByIdFromDB(id string) (*Book, error) {
	book := Book{}
	result := db.First(&book, id) // SELECT * FROM book WHERE id = :id;

	if result.Error != nil {
		log.Println("fail to get result from db")
		return nil, errors.New("fail to get result from db")
	}

	return &book, nil
}

func createBookInDB(book Book) (*Book, error) {
	result := db.Create(&book)

	if result.Error != nil {
		log.Println("fail to create new book in db")
		return nil, errors.New("fail to create new book in db")
	}

	return &book, nil
}

func editBookInDB(book Book) (*Book, error) {
	result := db.Save(&book) // Update all columns

	if result.Error != nil {
		log.Println("fail to update book in db")
		return nil, errors.New("fail to update book in db")
	}

	return &book, nil
}
