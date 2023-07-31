package models

import (
	"github.com/jinzhu/gorm"
	libdb "github.com/ssuloglu/simple-library-management-with-go/packages/db"
)

var conn *gorm.DB

type Book struct {
	gorm.Model
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	libdb.Connect()
	conn = libdb.GetConn()
	conn.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	conn.NewRecord(b)
	conn.Create(&b)
	return b
}

func GetBooks() []Book {
	var books []Book
	conn.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var thebook Book
	conn = conn.Where("ID=?", Id).Find(&thebook)
	return &thebook, conn
}

func GetBooksByAuthor(author string) []Book {
	var books []Book
	conn = conn.Where("author LIKE ?", "%"+author+"%").Find(&books)
	return books
}

func DeleteBook(Id int64) Book {
	var thebook Book
	conn.Where("ID=?", Id).Delete(thebook)
	return thebook
}
