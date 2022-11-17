package main

import (
    "fmt"
    "log"
	"io/ioutil"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Product struct {
	Id string `json:"Id"`
	Name string  `json:"name"`
	Description string `json:"description"`
	Price string `json:"price"`
	ExpirationDate string `json:"expiration_date"`
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Products)
}

func returnSingleProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    key := vars["id"]
    for _, product := range Products {
        if product.Id == key {
            json.NewEncoder(w).Encode(product)
        }
    }
}

func createNewProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)
	Products = append(Products, product)
    json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]
    for index, product := range Products {
        if product.Id == id {
            Products = append(Products[:index], Products[index+1:]...)
        }
    }
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updatedEvent Product
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)
	for i, product := range Products {
		if product.Id == id {

			product.Name = updatedEvent.Name
			product.Description = updatedEvent.Description
			product.Price = updatedEvent.Price
			product.ExpirationDate = updatedEvent.ExpirationDate
			Products[i] = product
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleRequests() {
	
	myRouter := mux.NewRouter().StrictSlash(true)
	
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/products", allProducts).Methods("GET") // READ
	myRouter.HandleFunc("/product", createNewProduct).Methods("POST") // CREATE
	myRouter.HandleFunc("/product/{id}", returnSingleProduct).Methods("GET") // READ
	myRouter.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE") // DELETE
	myRouter.HandleFunc("/product/{id}", updateProduct).Methods("PUT") // UPDATE

    log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	Products = []Product{
        Product{Id: "1", Name: "Cola Cola", Description: "Bebida azucarada gaseosa", Price: "$1.000", ExpirationDate: "25-07-2024"},
        Product{Id: "2", Name: "Chocman", Description: " Bizcocho bañado relleno con manjar", Price: "$800", ExpirationDate: "21-03-2023"},
    }
    handleRequests()
}