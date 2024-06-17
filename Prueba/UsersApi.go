package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User representa un usuario
type User struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

// Users es un slice que almacena todos los usuarios
var Users []User

// ObtenerUsuarios maneja la solicitud GET para obtener todos los usuarios
func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

func init() {
	// Inicializar el slice con algunos datos
	Users = append(Users, User{ID: "1", Nombre: "Juan Pérez", Email: "juan@example.com"})
	Users = append(Users, User{ID: "2", Nombre: "Ana Gómez", Email: "ana@example.com"})

	// Crear un nuevo router
	router := mux.NewRouter()

	// Definir las rutas
	router.HandleFunc("/ ", ObtenerUsuarios).Methods("GET")

	// Iniciar el servidor
	log.Println("Servidor iniciado en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
