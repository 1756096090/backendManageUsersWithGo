package routes

import (
	"FirstProyectWebEngineering/controllers"
	"github.com/gorilla/mux"
)

// TaskRoutes sets up the routes for tasks
func TaskRoutes(router *mux.Router, taskController *controllers.TaskController, employeeController *controllers.EmployeeController, proyectController *controllers.ProyectController) {
	router.HandleFunc("/tasks", taskController.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskController.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskController.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskController.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/dates", taskController.GetTasksBetweenDates).Methods("POST")
	
	
}
