package main

import (
    "log"
    "FirstProyectWebEngineering/config"
    "FirstProyectWebEngineering/controllers"
    "FirstProyectWebEngineering/repositories"
    "FirstProyectWebEngineering/routes"
    "FirstProyectWebEngineering/services"
    "net/http"
	"github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

func main() {
    config.ConnectDB()

    userRepository := repositories.UserRepository{}
    userService := services.UserService{Repository: userRepository}
    userController := controllers.UserController{Service: userService}
    
    employeeRepository := repositories.EmployeeRepository{}
    employeeService := services.EmployeeService{Repository: employeeRepository}
    employeeController := controllers.EmployeeController{Service: employeeService}

    proyectRepository := repositories.ProyectRepository{}
    proyectService := services.ProyectService{Repository: proyectRepository}
    proyectController := controllers.ProyectController{Service: proyectService}

    taskRepository := repositories.TaskRepository{}
    taskService := services.TaskService{Repository: taskRepository}
    taskController := controllers.TaskController{Service: taskService}
    

    router := mux.NewRouter()
    routes.UserRoutes(router, userController)
    routes.EmployeeRoutes(router, employeeController)
    routes.ProyectRoutes(router, proyectController)
    routes.TaskRoutes(router, &taskController, &employeeController, &proyectController )


    corsOptions := handlers.AllowedOrigins([]string{"*"}) 
    corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}) 
    corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}) 


    log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsOptions, corsMethods, corsHeaders)(router)))
}
