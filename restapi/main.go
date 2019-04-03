package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	//"fmt"
	"log"
	"math/rand"
	"strconv"
)

// Books Struct (Model)

type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Titulo string `json:"titulo"`
	Autor  *Autor `json:"autor"`
}

// Autor struct

type Autor struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

// Init libros como slice
var books []Book

// Buscar libros
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Buscar libro
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get parametro
	// Realiza ciclo en books y busca con id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// crear libro
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Actualizar libro
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get parametro
	// Realiza ciclo en books y busca con id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] // dejar el mismo id
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Borrar libro
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get parametro
	// Realiza ciclo en books y busca con id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) // borra el id
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "234235", Titulo: "Libro 1", Autor: &Autor{Nombre: "John", Apellido: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "235235", Titulo: "Libro 2", Autor: &Autor{Nombre: "Juan", Apellido: "Guerra"}})

	// Router Handlers / Endpoint

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
