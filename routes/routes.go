package routes

import (
	"todo/v1/controllers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/api/home", controllers.GetAll)
	route.HandleFunc("/api/add", controllers.CreateOne).Methods("POST")
	route.HandleFunc("/api/delete/{id}", controllers.DeleteOne).Methods("DELETE")
	route.HandleFunc("/api/completed/{id}", controllers.CompleteOne).Methods("PUT")
	route.HandleFunc("/api/update/{id}", controllers.UpdateOne).Methods("PUT")
	return route
}
