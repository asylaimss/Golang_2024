package main

import (
	"net/http"

	"TSIS-EX/handlers"
	"TSIS-EX/models"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	handlers.Actors = append(handlers.Actors, models.Actor{Name: "Damon Salvatore", Age: 30, Role: "Vampire"})
	handlers.Actors = append(handlers.Actors, models.Actor{Name: "Elena Gilbert", Age: 25, Role: "Human"})

	router.HandleFunc("/actors", handlers.GetActors).Methods("GET")
	router.HandleFunc("/actors/{name}", handlers.GetActor).Methods("GET")

	http.ListenAndServe(":8080", router)
}
