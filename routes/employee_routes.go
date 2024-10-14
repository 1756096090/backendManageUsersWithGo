package routes

import (
    "FirstProyectWebEngineering/controllers"
    "github.com/gorilla/mux"
)

func EmployeeRoutes(router *mux.Router, controller controllers.EmployeeController) {
    router.HandleFunc("/employees", controller.GetAllEmployees).Methods("GET")
    router.HandleFunc("/employees/{id}", controller.GetEmployeeByID).Methods("GET")
    router.HandleFunc("/employees", controller.CreateEmployee).Methods("POST")
    router.HandleFunc("/employees/{id}", controller.UpdateEmployee).Methods("PUT")
    router.HandleFunc("/employees/{id}", controller.DeleteEmployee).Methods("DELETE")
}
