package main

import (
    "fmt"
    "log"
	"strconv"
	"io/ioutil"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Product struct {
	Id int `json:"Id"`
	Name string  `json:"name"`
	Description string `json:"description"`
	Price int `json:"price"`
	ExpirationDate string `json:"expiration_date"`
}

var Products []Product

func check(e error) {
	/* 
	Genera un panic en caso de haber un error.
	Based on https://gobyexample.com/writing-files
	*/
    if e != nil {
        panic(e)
    }
}

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
	intKey, err := strconv.Atoi(key)
	check(err)
    for _, product := range Products {
        if product.Id == intKey {
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
	intId, err := strconv.Atoi(id)
	check(err)
    for index, product := range Products {
        if product.Id == intId {
            Products = append(Products[:index], Products[index+1:]...)
        }
    }
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := strconv.Atoi(id)
	check(err)
	var updatedEvent Product
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)
	for i, product := range Products {
		if product.Id == intId {
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
        Product{Id: 1, Name: "Cola Cola", Description: "Bebida azucarada gaseosa", Price: 1000, ExpirationDate: "25-07-2024"},
        Product{Id: 2, Name: "Chocman", Description: " Bizcocho bañado relleno con manjar", Price: 800, ExpirationDate: "21-03-2023"},
    }
    handleRequests()
}