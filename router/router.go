package router

import (
	"ToDOList/controller"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {

	route := mux.NewRouter()

	route.HandleFunc("/", controller.Show).Methods("GET")
	route.HandleFunc("/add", controller.Add).Methods("POST")
	route.HandleFunc("/delete/{id}", controller.Delete)
	route.HandleFunc("/complete/{id}", controller.Complete)
	route.HandleFunc("/updateTask/{id}", controller.Update)

	return route
}
