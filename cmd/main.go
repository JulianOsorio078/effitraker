package main

import (
	"EffiTrack/internal/db"
	"EffiTrack/internal/routes"
	"log"
	"net/http"
)

func main() {
	// Inicializar la conexión a la base de datos
	db.ConfigureConection()
	//Configuración de las Apis del proyecto
	routes.ConfigRoute()

	// Configurar las rutas y controladores
	// http.HandleFunc("/users", controllers.UserHandler)

	// Iniciar el servidor HTTP
	log.Println("Servidor iniciado en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
