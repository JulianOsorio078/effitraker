package middleware

import (
	"EffiTrack/internal/services"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// TODO: Crear la clave secreta
var jwtKey = []byte("your_secret_key") // Cambia esto a una clave secreta segura

// Claims estructura para JWT
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// AuthMiddleware verifica el token JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No Authorization header provided", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Opcionalmente, verifica si el usuario existe en la base de datos
		_, err = services.GetUserByUsername(claims.Username)
		if err != nil {
			http.Error(w, "Invalid user", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
