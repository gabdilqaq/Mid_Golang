package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

var (
	products []Product
)

func init() {
	products = []Product{{Id: "1", Key: "Key1", Value: "Value1"},
		{Id: "2", Key: "Key2", Value: "Value2"}}
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	for _, product := range products {
		if product.Key == key {
			json.NewEncoder(w).Encode(product)

		}
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updatedEvent Product
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)
	for i, product := range products {
		if product.Id == id {
			product.Key = updatedEvent.Key
			product.Value = updatedEvent.Value
			products[i] = product
			json.NewEncoder(w).Encode(product)
		}
	}
}
