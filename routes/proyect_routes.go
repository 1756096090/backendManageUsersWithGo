package routes

import (
    "FirstProyectWebEngineering/controllers"
    "github.com/gorilla/mux"
)

func ProyectRoutes(router *mux.Router, controller controllers.ProyectController) {
    router.HandleFunc("/proyects", controller.GetAllProyects).Methods("GET")
    router.HandleFunc("/proyects/{id}", controller.GetProyectByID).Methods("GET")
    router.HandleFunc("/proyects", controller.CreateProyect).Methods("POST")
    router.HandleFunc("/proyects/{id}", controller.UpdateProyect).Methods("PUT")
    router.HandleFunc("/proyects/{id}", controller.DeleteProyect).Methods("DELETE")
}
