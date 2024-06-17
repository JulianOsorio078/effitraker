package routes

import (
	"EffiTrack/internal/controllers"
	"EffiTrack/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigRoute() {
	router := mux.NewRouter()

	// Rutas públicas
	router.HandleFunc("/login", controllers.LoginHandler).Methods("POST")

	// Rutas protegidas
	protected := router.PathPrefix("/v1").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
	protected.HandleFunc("/users", controllers.GetAllUserHandler).Methods("GET")

	http.Handle("/", router)
}
