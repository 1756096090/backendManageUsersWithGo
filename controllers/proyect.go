package controllers

import (
	"FirstProyectWebEngineering/models"
	"FirstProyectWebEngineering/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProyectController struct {
	Service services.ProyectService
}

func (c *ProyectController) CreateProyect(w http.ResponseWriter, r *http.Request) {
	var proyect models.Proyect
	if err := json.NewDecoder(r.Body).Decode(&proyect); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := c.Service.CreateProyect(&proyect)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (c *ProyectController) GetAllProyects(w http.ResponseWriter, r *http.Request) {
	proyects, err := c.Service.GetAllProyects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(proyects)
}

func (c *ProyectController) GetProyectByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid proyect ID", http.StatusBadRequest)
		return
	}

	proyect, err := c.Service.GetProyectByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(proyect)
}

func (c *ProyectController) UpdateProyect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid proyect ID", http.StatusBadRequest)
		return
	}

	var proyect models.Proyect
	if err := json.NewDecoder(r.Body).Decode(&proyect); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := c.Service.UpdateProyect(id, &proyect)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (c *ProyectController) DeleteProyect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid proyect ID", http.StatusBadRequest)
		return
	}

	result, err := c.Service.DeleteProyect(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
