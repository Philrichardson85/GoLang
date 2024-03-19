package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
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

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
