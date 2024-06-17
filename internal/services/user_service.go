package services

import (
	"EffiTrack/internal/db"
	"EffiTrack/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers obtiene todos los usuarios de la base de datos
func GetAllUsers() ([]models.User, error) {
	rows, err := db.DB.Query("SELECT id, name, email, username, active FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Active); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(newUser models.User) (models.User, error) {
	hashedPassword, errorCript := HashPassword(string(newUser.Password))
	if errorCript != nil {
		return models.User{}, errorCript
	}
	newUser.Password = hashedPassword

	// Por ejemplo, puedes insertar el usuario en la base de datos usando db.DB.Exec() o similar
	_, err := db.DB.Exec("INSERT INTO users (name, email, password, username, active) VALUES (?, ?, ?, ?, 1)", newUser.Name, newUser.Email, newUser.Password, newUser.Username)
	if err != nil {
		return models.User{}, err
	}

	// Si la inserción es exitosa, devuelve el usuario creado
	return newUser, nil
}

// GetUserByUsername obtiene un usuario por su nombre de usuario
func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := db.DB.QueryRow("SELECT id, name, email, username, password, active FROM users WHERE username = ?", username).Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Password, &user.Active)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
