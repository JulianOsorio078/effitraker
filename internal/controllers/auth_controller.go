package controllers

import (
	"EffiTrack/internal/models"
	"EffiTrack/internal/services"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// TODO: Crear la clave secreta
var jwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// LoginHandler maneja las solicitudes de inicio de sesión y genera un token JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
		return
	}

	storedUser, err := services.GetUserByUsername(user.Username)
	if err != nil || storedUser.CheckPassword(string(user.Password)) != nil {
		http.Error(w, "Nombre de usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: storedUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
