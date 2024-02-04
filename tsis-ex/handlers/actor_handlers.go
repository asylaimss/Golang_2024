package handlers

import (
	"TSIS-EX/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var Actors []models.Actor

func GetActors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Actors)
}

func GetActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Actors {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Actor{})
}
