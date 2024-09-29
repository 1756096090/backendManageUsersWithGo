package routes

import (
    "FirstProyectWebEngineering/controllers"
    "github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router, controller controllers.UserController) {
    router.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
    router.HandleFunc("/users/{id}", controller.GetUserByID).Methods("GET")
    router.HandleFunc("/users", controller.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")
    router.HandleFunc("/login", controller.Login).Methods("POST") 
}
