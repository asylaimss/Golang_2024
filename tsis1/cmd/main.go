package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Actor struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
}

var actors []Actor

func GetActors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actors)
}

func GetActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range actors {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Actor{})
}

func main() {
	router := mux.NewRouter()

	actors = append(actors, Actor{Name: "Damon Salvatore", Age: 30, Role: "Vampire"})
	actors = append(actors, Actor{Name: "Elena Gilbert", Age: 25, Role: "Human"})
	actors = append(actors, Actor{Name: "Stefan Salvatore", Age: 28, Role: "Vampire"})
	actors = append(actors, Actor{Name: "Caroline Forbes", Age: 27, Role: "Vampire"})
	actors = append(actors, Actor{Name: "Bonnie Bennett", Age: 26, Role: "Witch"})
	actors = append(actors, Actor{Name: "Jeremy Gilbert", Age: 28, Role: "Human"})
	actors = append(actors, Actor{Name: "Alaric Saltzman", Age: 40, Role: "Human"})
	actors = append(actors, Actor{Name: "Klaus Mikaelson", Age: 1000, Role: "Original Vampire"})
	actors = append(actors, Actor{Name: "Rebekah Mikaelson", Age: 1000, Role: "Original Vampire"})
	actors = append(actors, Actor{Name: "Matt Donovan", Age: 29, Role: "Human"})
	actors = append(actors, Actor{Name: "Tyler Lockwood", Age: 30, Role: "Werewolf"})

	router.HandleFunc("/actors", GetActors).Methods("GET")
	router.HandleFunc("/actors/{name}", GetActor).Methods("GET")

	http.ListenAndServe(":8080", router)
}