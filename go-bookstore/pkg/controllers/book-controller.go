package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ForceGT/go-bookstore/pkg/models"
	"github.com/ForceGT/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	allBooks := models.GetAllBooks()
	res, _ := json.Marshal(allBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	book, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)

	vars := mux.Vars(r)
	bookID := vars["bookId"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	findBook, DB := models.GetBookByID(ID)
	if updatedBook.Name != "" {
		findBook.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		findBook.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		findBook.Publication = updatedBook.Publication
	}
	DB.Save(&findBook)
	res, _ := json.Marshal(findBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
