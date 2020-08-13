package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pizza struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

var pizzas []Pizza

func getPizzas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pizzas)
}

func getPizza(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range pizzas {
		for item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/pizzas", getPizzas).Methods("GET")
	router.HandleFunc("/pizzas/{id}", getPizza).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
