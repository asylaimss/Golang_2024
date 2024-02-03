package main

import (
	"net/http"

	"golang/handlers"
	"golang/models"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Добавление данных для примера
	handlers.Actors = append(handlers.Actors, models.Actor{Name: "Damon Salvatore", Age: 30, Role: "Vampire"})
	handlers.Actors = append(handlers.Actors, models.Actor{Name: "Elena Gilbert", Age: 25, Role: "Human"})

	// Установка маршрутов
	router.HandleFunc("/actors", handlers.GetActors).Methods("GET")
	router.HandleFunc("/actors/{name}", handlers.GetActor).Methods("GET")

	// Запуск сервера на порту 8080
	http.ListenAndServe(":8080", router)
}
