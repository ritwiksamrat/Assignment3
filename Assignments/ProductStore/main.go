package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Products struct {
	Name           string `json:"name"`
	ProductDesc    string `json:"desc"`
	ProductCompany string `json:"company"`
}

var product []Products

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range product {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Products{})
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pk Products
	_ = json.NewDecoder(r.Body).Decode(&pk)
	product = append(product, pk)
	json.NewEncoder(w).Encode(pk)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range product {
		if item.Name == params["name"] {
			product = append(product[:idx], product[idx+1:]...)
			var pk Products
			_ = json.NewDecoder(r.Body).Decode(&pk)
			pk.Name = params["name"]
			product = append(product, pk)
			json.NewEncoder(w).Encode(pk)
			return
		}
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range product {
		if item.Name == params["name"] {
			product = append(product[:idx], product[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(product)
}

func main() {
	r := mux.NewRouter()

	product = append(product, Products{Name: "Bat", ProductDesc: "CricketBat", ProductCompany: "Mongoose"})
	product = append(product, Products{Name: "Football", ProductDesc: "Football", ProductCompany: "Adidaas"})
	product = append(product, Products{Name: "CricketBall", ProductDesc: "Ball", ProductCompany: "SG"})

	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products/{name}", getProduct).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{name}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{name}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
