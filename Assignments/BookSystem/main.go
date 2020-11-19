package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Books struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var book []Books

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range book {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Books{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bk Books
	_ = json.NewDecoder(r.Body).Decode(&bk)
	book = append(book, bk)
	json.NewEncoder(w).Encode(bk)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range book {
		if item.Name == params["name"] {
			book = append(book[:idx], book[idx+1:]...)
			var bk Books
			_ = json.NewDecoder(r.Body).Decode(&bk)
			bk.Name = params["name"]
			book = append(book, bk)
			json.NewEncoder(w).Encode(bk)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range book {
		if item.Name == params["name"] {
			book = append(book[:idx], book[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(book)
}

func main() {
	r := mux.NewRouter()

	book = append(book, Books{Name: "Book1", Title: "FirstOne", Author: "Ritwik Mondal"})
	book = append(book, Books{Name: "Book2", Title: "SecondOne", Author: "Samrat"})
	book = append(book, Books{Name: "Book3", Title: "ThirdOne", Author: "Ishan"})

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{name}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{name}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{name}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
