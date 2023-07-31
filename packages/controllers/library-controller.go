package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/ssuloglu/simple-library-management-with-go/packages/models"
	"github.com/ssuloglu/simple-library-management-with-go/packages/utils"
)

func WriteHeader(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ExtractBookId(r *http.Request) int64 {
	bookId := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing bookid")
	}
	return id
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	res, _ := json.Marshal(books)

	WriteHeader(w, res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	//obtain book id
	id := ExtractBookId(r)

	// get the book from db
	thebook, _ := models.GetBookById(id)
	res, _ := json.Marshal(thebook)

	WriteHeader(w, res)
}

func GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	//obtain author name
	author := mux.Vars(r)["author"]

	books := models.GetBooksByAuthor(author)
	res, _ := json.Marshal(books)

	WriteHeader(w, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	b := &models.Book{}
	utils.ParseBody(r, b)

	newBook := b.CreateBook()
	res, _ := json.Marshal(newBook)

	WriteHeader(w, res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//obtain book id
	id := ExtractBookId(r)

	// delete the book
	thebook := models.DeleteBook(id)
	res, _ := json.Marshal(thebook)

	WriteHeader(w, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookDetails := &models.Book{}
	utils.ParseBody(r, bookDetails)

	//obtain book id
	id := ExtractBookId(r)

	var db *gorm.DB
	thebook, db := models.GetBookById(id)
	if bookDetails.Name != "" {
		thebook.Name = bookDetails.Name
	}
	if bookDetails.Author != "" {
		thebook.Author = bookDetails.Author
	}
	if bookDetails.Publisher != "" {
		thebook.Publisher = bookDetails.Publisher
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)

	WriteHeader(w, res)
}
