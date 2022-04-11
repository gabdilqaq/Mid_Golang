package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	router.HandleFunc("/store/{key}", getProduct).Methods("GET")
	router.HandleFunc("/store/{id}", UpdateProduct).Methods("PUT")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
