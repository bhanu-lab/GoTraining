package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id       string  `json:"id"`
	Isbn     string  `json:"isbn"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Author   *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var books []Book

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

//create a book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//update a book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//Init router using mux
	r := mux.NewRouter()

	books = append(books, Book{Id: "1", Isbn: "4353534", Title: "Elon Musk", Category: "AutoBiography", Author: &Author{FirstName: "ASHLEE", LastName: "VANCE"}})
	books = append(books, Book{Id: "2", Isbn: "4353534", Title: "Psycho-Cybernetics", Category: "SelgHelp", Author: &Author{FirstName: "MAXWELL", LastName: "MALTZ"}})

	//end points
	r.HandleFunc("/br/books", getBooks).Methods("GET")
	r.HandleFunc("/br/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/br/books", createBook).Methods("POST")
	r.HandleFunc("/br/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/br/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8900", r))

}
