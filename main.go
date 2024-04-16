package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Docker using Go and deploy in AWS..")

	r := mux.NewRouter()
	r.HandleFunc("/tejas/home", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("This page is specially made for you Tejas!!")
	}).Methods("GET")

	r.HandleFunc("/tejas/exit", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Bye Tejas!!")
	}).Methods("GET")

	fmt.Println("Listening on port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
