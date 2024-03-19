package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name"`
	City string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))	
}

func greet(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request)  {
	customers := []Customer {
		{"Ash","Atlanta", "32303"},
		{"Tom","Athens", "21303"}, 
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}