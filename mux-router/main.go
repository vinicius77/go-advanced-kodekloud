package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home page")
	fmt.Println("Endpoint hit: homepage")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: products")
	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["id"]

	for _, p := range Products {
		if p.Id == productId {
			json.NewEncoder(w).Encode(p)
		}
	}
}

// go get github.com/go-sql-driver/mysql
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/products", getProducts)
	router.HandleFunc("/product/{id}", getProduct)
	router.HandleFunc("/", homepage)

	fmt.Println("Server listening on http://localhost:8093")

	http.ListenAndServe("localhost:8093", router)
}

func main() {

	Products = []Product{
		{Id: "22", Name: "Product 1", Quantity: 55, Price: 22.54},
	}

	handleRequests()

}
