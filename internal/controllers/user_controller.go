package controllers

import (
	"EffiTrack/internal/models"
	"EffiTrack/internal/services"
	"encoding/json"
	"net/http"
)

// GetAllUserHandler maneja las solicitudes HTTP para los todos usuarios
func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User // Suponiendo que tienes una estructura User en services
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para crear el usuario
	createdUser, err := services.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Codificar la respuesta JSON con el usuario creado y enviarla
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}
