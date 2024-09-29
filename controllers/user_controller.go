package controllers

import (
	"FirstProyectWebEngineering/models"
	"FirstProyectWebEngineering/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
    Service services.UserService
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    result, err := c.Service.CreateUser(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(result)
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := c.Service.GetAllUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    user, err := c.Service.GetUserByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    result, err := c.Service.UpdateUser(id, &user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(result)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    result, err := c.Service.DeleteUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(result)
}


func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
    var user models.User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    log.Printf("User login attempt: %s", user.Email) // Add this for logging

    result, err := c.Service.Login(user.Email, user.Password) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(result)
}
