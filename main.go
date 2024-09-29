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
    // Connect to the database
    config.ConnectDB()

    // Initialize repositories, services, and controllers
    userRepository := repositories.UserRepository{}
    userService := services.UserService{Repository: userRepository}
    userController := controllers.UserController{Service: userService}

    // Create a new router
    router := mux.NewRouter()
    routes.UserRoutes(router, userController)


    corsOptions := handlers.AllowedOrigins([]string{"*"}) 
    corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}) 
    corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}) 


    log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsOptions, corsMethods, corsHeaders)(router)))
}
