package models

import "golang.org/x/crypto/bcrypt"

// User representa un usuario en la base de datos
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Active   string `json:"active"`
}

// // CheckPassword verifica si la contraseña proporcionada coincide con la contraseña encriptada del usuario
// func (u User) CheckPassword(password string) error {
// 	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
// }

// CheckPassword verifica si la contraseña proporcionada coincide con la contraseña encriptada del usuario
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
